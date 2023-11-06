/*
Copyright 2023.

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

package v1alpha1

import (
	"github.com/yolo-operator/yolo-operator/pkg/condition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ChaosSimulationSpec defines the desired state of ChaosSimulation
type ChaosSimulationSpec struct {
	Level string `json:"foo,omitempty"`
}

// ChaosSimulationStatus defines the observed state of ChaosSimulation
type ChaosSimulationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Conditions represent the latest available observations of an object's state
	// +optional
	Conditions []condition.Condition `json:"conditions,omitempty"`

	// Output is the output of the command
	Output string `json:"output,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ChaosSimulation is the Schema for the chaossimulations API
type ChaosSimulation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChaosSimulationSpec   `json:"spec,omitempty"`
	Status ChaosSimulationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ChaosSimulationList contains a list of ChaosSimulation
type ChaosSimulationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChaosSimulation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChaosSimulation{}, &ChaosSimulationList{})
}
