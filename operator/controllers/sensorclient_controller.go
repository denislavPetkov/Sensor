/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	sensorCliV1 "github.com/denislavPetkov/sensor/operator/api/v1"
)

// SensorClientReconciler reconciles a SensorClient object
type SensorClientReconciler struct {
	client.Client
	Log            logr.Logger
	Scheme         *runtime.Scheme
	timeOutChannel <-chan time.Time
	deltaChannel   <-chan time.Time
}

//+kubebuilder:rbac:groups=sensor.cli,resources=sensorclients,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=sensor.cli,resources=sensorclients/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=sensor.cli,resources=sensorclients/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SensorClient object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile

var logger logr.Logger

const retryPeriod int32 = 5

func (r *SensorClientReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger = r.Log.WithValues("sensorClient", req.NamespacedName)

	crInstance := &sensorCliV1.SensorClient{}

	err := r.Get(ctx, req.NamespacedName, crInstance)
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Info("Ignoring since CustomRersource is deleted")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Failed to get SensorClient")
		return ctrl.Result{}, err
	}

	deployment := &appsv1.Deployment{}
	err = r.Get(ctx, types.NamespacedName{Name: crInstance.Name, Namespace: crInstance.Namespace}, deployment)

	if err != nil {

		if !errors.IsNotFound(err) {
			logger.Error(err, "Failed to get existing deployment", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
			r.setCrStatusCondition(crInstance, metav1.ConditionFalse, sensorCliV1.FailedDeployment, err.Error())
			crStatusUpdateError := r.updateCrStatus(ctx, crInstance)
			if crStatusUpdateError != nil {
				return reconcile.Result{}, crStatusUpdateError
			}
			return ctrl.Result{}, err
		}

		err := r.installDeployment(ctx, crInstance, deployment)

		if err != nil {
			r.setCrStatusCondition(crInstance, metav1.ConditionFalse, sensorCliV1.FailedDeployment, err.Error())
			crStatusUpdateError := r.updateCrStatus(ctx, crInstance)
			if crStatusUpdateError != nil {
				return reconcile.Result{}, crStatusUpdateError
			}
			return reconcile.Result{}, err
		}

		return reconcile.Result{}, nil
	}

	err = r.updateDeployment(ctx, crInstance, deployment)
	if err != nil {
		r.setCrStatusCondition(crInstance, metav1.ConditionFalse, sensorCliV1.FailedDeployment, err.Error())
		crStatusUpdateError := r.updateCrStatus(ctx, crInstance)
		if crStatusUpdateError != nil {
			return reconcile.Result{}, crStatusUpdateError
		}
		return reconcile.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *SensorClientReconciler) installDeployment(ctx context.Context, crInstance *sensorCliV1.SensorClient, deployment *appsv1.Deployment) error {
	err := r.createCrDeployment(ctx, crInstance)
	if err != nil {
		return err
	}

	r.setCrStatusCondition(crInstance, metav1.ConditionUnknown, sensorCliV1.PendingDeployment, sensorCliV1.PendingInstallMsg)
	crStatusUpdateError := r.updateCrStatus(ctx, crInstance)
	if crStatusUpdateError != nil {
		return crStatusUpdateError
	}

	err = r.setCrStatusConditionBasedOnDeploymentReadiness(ctx, crInstance, deployment, sensorCliV1.SuccessfulInstallMsg)
	if err != nil {
		return err
	}

	crStatusUpdateError = r.updateCrStatus(ctx, crInstance)
	if crStatusUpdateError != nil {
		return crStatusUpdateError
	}

	return nil
}

func (r *SensorClientReconciler) setCrStatusConditionBasedOnDeploymentReadiness(ctx context.Context, crInstance *sensorCliV1.SensorClient, deployment *appsv1.Deployment, statusUpdateMsg string) error {
	err := r.checkDeploymentReadiness(ctx, crInstance, deployment)
	if err != nil {
		return err
	}
	r.setCrStatusCondition(crInstance, metav1.ConditionTrue, sensorCliV1.SuccessfullDeployment, statusUpdateMsg)
	return nil
}

func (r *SensorClientReconciler) checkDeploymentReadiness(ctx context.Context, crInstance *sensorCliV1.SensorClient, deployment *appsv1.Deployment) error {

	r.timeOutChannel = time.After(time.Duration(crInstance.Spec.DeploymentWaitingTimeout) * time.Second)
	r.deltaChannel = time.Tick(time.Duration(retryPeriod) * time.Millisecond)

	for {
		select {
		case <-r.deltaChannel:
			err := r.Get(ctx, types.NamespacedName{Name: crInstance.Name, Namespace: crInstance.Namespace}, deployment)
			if err != nil {
				logger.Error(err, "Failed to get existing deployment", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
				return err
			}
			isDepoymentReady := r.getDeploymentReadiness(deployment, crInstance)
			if isDepoymentReady {
				return nil
			}
		case <-r.timeOutChannel:
			logger.Info("Deployment not ready in timeout period")
			return fmt.Errorf("Timeout occurred during waiting Deployment readiness")
		}
	}
}

func (r *SensorClientReconciler) getDeploymentReadiness(deployment *appsv1.Deployment, crInstance *sensorCliV1.SensorClient) bool {

	deploymentReadyReplicas := deployment.Status.ReadyReplicas
	crInstanceSpecifiedReplicas := crInstance.Spec.ReplicasCount
	return deploymentReadyReplicas == crInstanceSpecifiedReplicas

}

func (r *SensorClientReconciler) updateDeployment(ctx context.Context, crInstance *sensorCliV1.SensorClient, deployment *appsv1.Deployment) error {
	areReplicasUpdated, err := r.updateCrDeploymentReplicas(ctx, crInstance, deployment)
	if err != nil {
		return err
	}

	areArgsUpdated, err := r.updateCrDeploymentArguments(ctx, crInstance, deployment)
	if err != nil {
		return err
	}

	isImageUpdated, err := r.updateCrDeploymentImage(ctx, crInstance, deployment)
	if err != nil {
		return err
	}

	if areArgsUpdated || areReplicasUpdated || isImageUpdated {
		r.setCrStatusCondition(crInstance, metav1.ConditionUnknown, sensorCliV1.PendingDeployment, sensorCliV1.PendingUpdateMsg)
		crStatusUpdateError := r.updateCrStatus(ctx, crInstance)
		if crStatusUpdateError != nil {
			return crStatusUpdateError
		}

		setStatusError := r.setCrStatusConditionBasedOnDeploymentReadiness(ctx, crInstance, deployment, sensorCliV1.SuccessfulUpdateMsg)
		if setStatusError != nil {
			return setStatusError
		}

		crStatusUpdateError = r.updateCrStatus(ctx, crInstance)
		if crStatusUpdateError != nil {
			return crStatusUpdateError
		}

		return nil

	}

	return nil

}

func (r *SensorClientReconciler) updateCrDeploymentImage(ctx context.Context, crInstance *sensorCliV1.SensorClient, deployment *appsv1.Deployment) (bool, error) {
	deploymentImage := deployment.Spec.Template.Spec.Containers[0].Image
	crInstanceSpecifiedImage := crInstance.Spec.SensorClientImage

	if deploymentImage == crInstanceSpecifiedImage {
		return false, nil
	}

	logger.Info("Updating deployment image", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
	deployment.Spec.Template.Spec.Containers[0].Image = crInstanceSpecifiedImage
	err := r.Update(ctx, deployment)
	if err == nil {
		logger.Info("Successfully updated deployment image", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
		return true, nil
	}

	logger.Error(err, "Failed to update deployment image", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
	return false, err
}

func (r *SensorClientReconciler) updateCrDeploymentReplicas(ctx context.Context, crInstance *sensorCliV1.SensorClient, deployment *appsv1.Deployment) (bool, error) {

	crInstanceSpecifiedReplicas := crInstance.Spec.ReplicasCount
	deploymentReplicas := *deployment.Spec.Replicas

	if deploymentReplicas == crInstanceSpecifiedReplicas {
		return false, nil
	}

	logger.Info("Updating deployment replicas", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
	deployment.Spec.Replicas = &crInstanceSpecifiedReplicas
	err := r.Update(ctx, deployment)
	if err == nil {
		logger.Info("Successfully updated deployment replicas count", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
		return true, nil
	}

	logger.Error(err, "Failed to update deployment replicas", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
	return false, err

}

func (r *SensorClientReconciler) updateCrDeploymentArguments(ctx context.Context, crInstance *sensorCliV1.SensorClient, deployment *appsv1.Deployment) (bool, error) {

	newSensorClientArgs := crInstance.Spec.Args.GetArgs()
	currentSensorClientArgs := deployment.Spec.Template.Spec.Containers[0].Args

	if reflect.DeepEqual(currentSensorClientArgs, newSensorClientArgs) {
		return false, nil
	}

	logger.Info("Updating deployment arguments", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
	deployment = r.getDeployment(crInstance)
	err := r.Update(ctx, deployment)
	if err == nil {
		logger.Info("Successfully updated deployment arguments", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
		return true, nil
	}
	logger.Error(err, "Failed to update deployment arguments", "Deployment.Namespace", deployment.Namespace, "Deployment.Name", deployment.Name)
	return false, err

}

func (r *SensorClientReconciler) createCrDeployment(ctx context.Context, crInstance *sensorCliV1.SensorClient) error {

	dep := r.getDeployment(crInstance)
	logger.Info("Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
	err := r.Create(ctx, dep)
	if err == nil {
		logger.Info("Deployment successfully created", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
		return nil
	}
	logger.Error(err, "Failed to create deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
	return err

}

func (r *SensorClientReconciler) updateCrStatus(ctx context.Context, crInstance *sensorCliV1.SensorClient) error {

	err := r.Status().Update(ctx, crInstance)
	if err == nil {
		logger.Info("Successfully updated CR status", "Reason", crInstance.Status.Conditions.Reason, "Message", crInstance.Status.Conditions.Message)
		return nil
	}
	logger.Error(err, "Failed to update CR status", "CR.Namespace", crInstance.Namespace, "CR.Name", crInstance.Name)
	return err

}

func (r *SensorClientReconciler) setCrStatusCondition(crInstance *sensorCliV1.SensorClient, newStatus metav1.ConditionStatus, reason string, message string) bool {

	oldCondition := &crInstance.Status.Conditions

	newCondition := sensorCliV1.SensorClientCondition{
		Status:  newStatus,
		Reason:  reason,
		Message: message,
	}

	return sensorCliV1.SetStatusCondition(oldCondition, newCondition)

}

func (r *SensorClientReconciler) getDeployment(crInstance *sensorCliV1.SensorClient) *appsv1.Deployment {

	lbls := labelsForApp(crInstance.Name)
	replicas := crInstance.Spec.ReplicasCount
	args := crInstance.Spec.Args.GetArgs()
	secrets := crInstance.Spec.ImagePullSecrets

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      crInstance.Name,
			Namespace: crInstance.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: lbls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: lbls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:           crInstance.Spec.SensorClientImage,
						ImagePullPolicy: corev1.PullAlways,
						Name:            "sensor-client",
						Args:            args,
					}},
					ImagePullSecrets: secrets,
				},
			},
		},
	}

	controllerutil.SetControllerReference(crInstance, dep, r.Scheme)
	return dep
}

func labelsForApp(name string) map[string]string {
	return map[string]string{"app": name}
}

// SetupWithManager sets up the controller with the Manager.
func (r *SensorClientReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sensorCliV1.SensorClient{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
