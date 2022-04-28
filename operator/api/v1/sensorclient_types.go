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

package v1

import (
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SensorClientSpec defines the desired state of SensorClient
type SensorClientSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Arguments are the commands to run the sensor client with
	Args Args `json:"arguments"`

	//+kubebuilder:validation:Minimum=1
	// replicas_count is the number of replicas in the sensor client deployment
	ReplicasCount int32 `json:"replicas_count"`

	//+kubebuilder:validation:Minimum=10
	// deployment_waiting_timeout is the time waited in seconds for the deployment to become ready
	DeploymentWaitingTimeout int32 `json:"deployment_waiting_timeout"`

	// Sensor_client_image is the image for the sensor client
	SensorClientImage string `json:"sensor_client_image"`

	// Image_pull_secrets are the secrets needed to pull the image
	ImagePullSecrets []corev1.LocalObjectReference `json:"image_pull_secrets,omitempty"`
}

type Args struct {
	// Delta_duration is the time in seconds between each measurement by the sensor
	DeltaDuration int `json:"delta_duration,omitempty"`
	// Total_duration is the time in seconds after which the sensor stops
	TotalDuration int `json:"total_duration,omitempty"`
	// Format is the format of the output
	Format string `json:"format,omitempty"`
	// Sensor_groups are the sensor groups to run the sensor with
	SensorGroups []string `json:"sensor_groups"`
	// Web_hook_url is the url to send the measurements to
	WebHook string `json:"web_hook_url,omitempty"`
}

func (args Args) GetArgs() []string {
	var arguments []string

	for _, sGroup := range args.SensorGroups {
		arguments = append(arguments, fmt.Sprintf("-sensor_group=%v", sGroup))
	}
	arguments = append(arguments, fmt.Sprintf("-format=%v", args.Format))
	arguments = append(arguments, fmt.Sprintf("-delta_duration=%v", args.DeltaDuration))
	arguments = append(arguments, fmt.Sprintf("-total_duration=%v", args.TotalDuration))
	arguments = append(arguments, fmt.Sprintf("-web_hook_url=%v", args.WebHook))

	return arguments
}

// SetStatusCondition sets the corresponding condition to newCondition.
func SetStatusCondition(condition *SensorClientCondition, newCondition SensorClientCondition) bool {
	if condition.Status == newCondition.Status {
		return false
	}
	condition.Type = LastOperationSucceeded
	condition.LastTransitionTime = metav1.NewTime(time.Now())
	condition.Status = newCondition.Status
	condition.Reason = newCondition.Reason
	condition.Message = newCondition.Message
	return true
}

// IsStatusConditionTrue returns true when the SensorClientConditionType is present and set to `metav1.ConditionTrue`
func IsStatusConditionTrue(condition SensorClientCondition) bool {
	return IsStatusConditionPresentAndEqual(condition, metav1.ConditionTrue)
}

// IsStatusConditionFalse returns true when the SensorClientConditionType is present and set to `metav1.ConditionFalse`
func IsStatusConditionFalse(condition SensorClientCondition) bool {
	return IsStatusConditionPresentAndEqual(condition, metav1.ConditionFalse)
}

// IsStatusConditionPresentAndEqual returns true when SensorClientConditionType is present and equal to status.
func IsStatusConditionPresentAndEqual(condition SensorClientCondition, status metav1.ConditionStatus) bool {
	return condition.Status == status
}

type SensorClientConditionType string

// These are valid conditions of a sensor client.
const (
	// LastOperationSucceeded describes the last known state of the last deployment
	LastOperationSucceeded SensorClientConditionType = "LastOperationSucceeded"

	PendingDeployment     = "PendingDeployment"
	FailedDeployment      = "FailedDeployment"
	SuccessfullDeployment = "SuccessfulDeployment"
)

const (
	PendingInstallMsg    = "Install pending"
	SuccessfulInstallMsg = "Install successful"
	PendingUpdateMsg     = "Update pending"
	SuccessfulUpdateMsg  = "Update successful"
)

// SensorClientCondition describes the state of a sensor client at a certain point.
type SensorClientCondition struct {
	// Type of sensor client condition.
	// +required
	// +kubebuilder:validation:Required
	Type SensorClientConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=True;False;Unknown
	Status metav1.ConditionStatus `json:"status"`
	// Last time the condition transitioned from one status to another.
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// The reason for the condition's last transition.
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:validation:MinLength=1
	Reason string `json:"reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// SensorClientStatus defines the observed state of SensorClient
type SensorClientStatus struct {
	// Conditions represent the latest available observations of the sensor client's state
	Conditions SensorClientCondition `json:"conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

//+kubebuilder:printcolumn:name="LastOperationSucceeded",type=string,JSONPath=`.status.conditions.status`
//+kubebuilder:printcolumn:name="Replicas",type=integer,JSONPath=`.spec.replicas_count`
// SensorClient is the Schema for the sensorclients API
type SensorClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SensorClientSpec   `json:"spec,omitempty"`
	Status SensorClientStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SensorClientList contains a list of SensorClient
type SensorClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SensorClient `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SensorClient{}, &SensorClientList{})
}
