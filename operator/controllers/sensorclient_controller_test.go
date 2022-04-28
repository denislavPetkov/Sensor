package controllers

import (
	"context"
	"time"

	sensorCliV1 "github.com/denislavPetkov/sensor/operator/api/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
)

var _ = Describe("SensorClientController", func() {

	const (
		apiVersion = "sensor.cli/v1"
		kind       = "SensorClient"

		SensorName = "test-sensor-client"

		// Modify according to your needs
		SensorNamespace = "sensor-app"
		// Modify according to your needs
		sensorImage = "docker.io/sensor/cli:0.0.1"
		// Modify according to your needs
		newSensorClientImage = "docker.io/sensor/cli:0.0.2"
		// Modify according to your needs
		secretName = "docker-registry-secret"

		deploymentSize int32 = 1

		timeout           = time.Second * 25
		interval          = time.Millisecond * 5
		deploymentTimeout = 20
	)

	var (
		sensorGroups = []string{"CPU_TEMP", "MEMORY_USAGE"}
		arguments    = sensorCliV1.Args{
			TotalDuration: 300,
			DeltaDuration: 10,
			SensorGroups:  sensorGroups,
		}

		sensorClientLookupKey types.NamespacedName
		sensorClientCr        *sensorCliV1.SensorClient
		ctx                   context.Context
	)

	sensorClientCr = &sensorCliV1.SensorClient{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiVersion,
			Kind:       kind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      SensorName,
			Namespace: SensorNamespace,
		},
		Spec: sensorCliV1.SensorClientSpec{
			Args:                     arguments,
			ReplicasCount:            deploymentSize,
			SensorClientImage:        sensorImage,
			DeploymentWaitingTimeout: deploymentTimeout,
			ImagePullSecrets: []corev1.LocalObjectReference{{
				Name: secretName,
			}},
		},
	}

	Context("[in cluster] When a SensorClient CR is deployed, then a SensorClient deployment is created", func() {

		AfterEach(func() {
			Expect(k8sClient.Delete(ctx, sensorClientCr)).Should(Succeed())
		})

		It("Should create a SensorClient deployment and update the SensorClient status", func() {
			ctx = context.Background()
			sensorClientLookupKey = types.NamespacedName{Name: SensorName, Namespace: SensorNamespace}

			By("Creating a new SensorClient")
			Expect(k8sClient.Create(ctx, sensorClientCr)).Should(Succeed())

			By("Checking CR status is set to 'Unknown' and status.reason to 'PendingDeployment'")
			Eventually(func() []string {
				err := k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)
				Expect(err).NotTo(HaveOccurred())
				return []string{string(sensorClientCr.Status.Conditions.Status), sensorClientCr.Status.Conditions.Reason}
			}, timeout, interval).Should(Equal([]string{string(metav1.ConditionUnknown), sensorCliV1.PendingDeployment}))

			By("Getting SensorClient CR from the cluster")
			sensorClientCr = &sensorCliV1.SensorClient{}
			Expect(k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)).Should(Succeed())

			deployment := &appsv1.Deployment{}

			By("Checking that the sensorClient deployment is created")
			Eventually(func() bool {
				err := k8sClient.Get(ctx, sensorClientLookupKey, deployment)
				return err == nil
			}, timeout, interval).Should(BeTrue())

			By("Checking CR status is set to 'True' and status.reason to 'SuccessfulDeployment'")
			Eventually(func() []string {
				err := k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)
				Expect(err).NotTo(HaveOccurred())
				return []string{string(sensorClientCr.Status.Conditions.Status), sensorClientCr.Status.Conditions.Reason}
			}, timeout, interval).Should(Equal([]string{string(metav1.ConditionTrue), sensorCliV1.SuccessfullDeployment}))

			By("Getting SensorClient CR from the cluster")
			Expect(k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)).Should(Succeed())

			By("Updating SensorClient arguments")
			newSensorGroups := []string{"CPU_USAGE"}
			sensorClientCr.Spec.Args.SensorGroups = newSensorGroups
			Expect(k8sClient.Update(ctx, sensorClientCr)).Should(Succeed())

			By("Checking CR status is set to 'Unknown' and status.reason to 'PendingDeployment'")
			Eventually(func() []string {
				err := k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)
				Expect(err).NotTo(HaveOccurred())
				return []string{string(sensorClientCr.Status.Conditions.Status), sensorClientCr.Status.Conditions.Reason}
			}, timeout, interval).Should(Equal([]string{string(metav1.ConditionUnknown), sensorCliV1.PendingDeployment}))

			By("Checking updated deployment arguments")
			Eventually(func() ([]string, error) {
				err := k8sClient.Get(ctx, sensorClientLookupKey, deployment)
				if err != nil {
					return nil, err
				}

				deploymentArgs := deployment.Spec.Template.Spec.Containers[0].Args
				return deploymentArgs, nil
			}, timeout, interval).Should(Equal(sensorClientCr.Spec.Args.GetArgs()))

			By("Checking CR status is set to 'True' and status.reason to 'SuccessfulDeployment'")
			Eventually(func() []string {
				err := k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)
				Expect(err).NotTo(HaveOccurred())
				return []string{string(sensorClientCr.Status.Conditions.Status), sensorClientCr.Status.Conditions.Reason}
			}, timeout, interval).Should(Equal([]string{string(metav1.ConditionTrue), sensorCliV1.SuccessfullDeployment}))

			By("Getting SensorClient CR from the cluster")
			Expect(k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)).Should(Succeed())

			By("Updating SensorClient deployment image")
			sensorClientCr.Spec.SensorClientImage = newSensorClientImage
			Expect(k8sClient.Update(ctx, sensorClientCr)).Should(Succeed())

			By("Checking CR status is set to 'Unknown' and status.reason to 'PendingDeployment'")
			Eventually(func() []string {
				err := k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)
				Expect(err).NotTo(HaveOccurred())
				return []string{string(sensorClientCr.Status.Conditions.Status), sensorClientCr.Status.Conditions.Reason}
			}, timeout, interval).Should(Equal([]string{string(metav1.ConditionUnknown), sensorCliV1.PendingDeployment}))

			By("Checking updated deployment image")
			Eventually(func() (string, error) {
				err := k8sClient.Get(ctx, sensorClientLookupKey, deployment)
				if err != nil {
					return "", err
				}
				deploymentImage := deployment.Spec.Template.Spec.Containers[0].Image
				return deploymentImage, nil
			}, timeout, interval).Should(Equal(newSensorClientImage))

			By("Checking CR status is set to 'True' and status.reason to 'SuccessfulDeployment'")
			Eventually(func() []string {
				err := k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)
				Expect(err).NotTo(HaveOccurred())
				return []string{string(sensorClientCr.Status.Conditions.Status), sensorClientCr.Status.Conditions.Reason}
			}, timeout, interval).Should(Equal([]string{string(metav1.ConditionTrue), sensorCliV1.SuccessfullDeployment}))

			By("Getting SensorClient CR from the cluster")
			Expect(k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)).Should(Succeed())

			By("Updating SensorClient deployment size")
			newDeploymentSize := 2
			sensorClientCr.Spec.ReplicasCount = int32(newDeploymentSize)
			Expect(k8sClient.Update(ctx, sensorClientCr)).Should(Succeed())

			By("Checking CR status is set to 'Unknown' and status.reason to 'PendingDeployment'")
			Eventually(func() []string {
				err := k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)
				Expect(err).NotTo(HaveOccurred())
				return []string{string(sensorClientCr.Status.Conditions.Status), sensorClientCr.Status.Conditions.Reason}
			}, timeout, interval).Should(Equal([]string{string(metav1.ConditionUnknown), sensorCliV1.PendingDeployment}))

			By("Checking the replicas count of the new deployment")
			Eventually(func() int {
				err := k8sClient.Get(ctx, sensorClientLookupKey, deployment)
				if err != nil {
					return 0
				}
				return int(*deployment.Spec.Replicas)
			}, timeout, interval).Should(Equal(newDeploymentSize))

			By("Checking CR status is set to 'True' and status.reason to 'SuccessfulDeployment'")
			Eventually(func() []string {
				err := k8sClient.Get(ctx, sensorClientLookupKey, sensorClientCr)
				Expect(err).NotTo(HaveOccurred())
				return []string{string(sensorClientCr.Status.Conditions.Status), sensorClientCr.Status.Conditions.Reason}
			}, timeout, interval).Should(Equal([]string{string(metav1.ConditionTrue), sensorCliV1.SuccessfullDeployment}))

		})
	})

	Describe("getDeployment()", func() {
		Context("When getDeployment(), then correct k8s Deployment resource is returned", func() {

			testSensorClientCr := &sensorCliV1.SensorClient{
				TypeMeta: metav1.TypeMeta{
					APIVersion: apiVersion,
					Kind:       kind,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      SensorName,
					Namespace: SensorNamespace,
				},
				Spec: sensorCliV1.SensorClientSpec{
					Args:                     arguments,
					ReplicasCount:            deploymentSize,
					SensorClientImage:        sensorImage,
					DeploymentWaitingTimeout: deploymentTimeout,
					ImagePullSecrets: []corev1.LocalObjectReference{
						{
							Name: secretName,
						},
					},
				},
			}

			var SensorClientReconciler = &SensorClientReconciler{Scheme: scheme.Scheme}
			var replicasCount int32 = deploymentSize
			expectedSensorClientDeployment := &appsv1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testSensorClientCr.Name,
					Namespace: testSensorClientCr.Namespace,
				},
				Spec: appsv1.DeploymentSpec{
					Replicas: &replicasCount,
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{"app": testSensorClientCr.Name},
					},
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{"app": testSensorClientCr.Name},
						},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{{
								Image:           testSensorClientCr.Spec.SensorClientImage,
								ImagePullPolicy: corev1.PullAlways,
								Name:            "sensor-client",
								Args:            testSensorClientCr.Spec.Args.GetArgs(),
							}},
							ImagePullSecrets: []corev1.LocalObjectReference{{
								Name: secretName,
							}},
						},
					},
				},
			}
			It("Should return valid Deployment resource", func() {
				actualSensorClientDeployment := SensorClientReconciler.getDeployment(testSensorClientCr)
				Expect(actualSensorClientDeployment.Spec).To(Equal(expectedSensorClientDeployment.Spec))
			})
		})
	})
})
