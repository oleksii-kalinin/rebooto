/*
Copyright 2025.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RebootSpec defines the desired state of Reboot.
type RebootSpec struct {
	// Labels selector for the resources to restart during reconcile
	MatchLabels map[string]string `json:"matchLabels"`
	// The name of the target resource to watch (configmap/secret)
	TargetResources []string `json:"targetResources"`
	// Target namespaces where to watch for the targets
	// +optional
	TargetNamespaces []string `json:"targetNamespaces,omitempty"`
	// Delay before restart (rollout) triggered
	GracePeriodSeconds *int64 `json:"gracePeriodSeconds"`
	// Delay between latest triggered restart
	// +optional
	DelayBetweenRestarts *metav1.Time `json:"delayBetweenRestarts,omitempty"`
}

// RebootStatus defines the observed state of Reboot.
type RebootStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Track last triggered restart time
	// +optional
	LastRestartTime *metav1.Timestamp `json:"lastRestartTime,omitempty"`
	// Restart count
	// +optional
	RestartCount *int32 `json:"restartCount,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Reboot is the Schema for the reboots API.
type Reboot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RebootSpec   `json:"spec,omitempty"`
	Status RebootStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RebootList contains a list of Reboot.
type RebootList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Reboot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Reboot{}, &RebootList{})
}
