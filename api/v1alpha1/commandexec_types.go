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

// CommandExecSpec defines the desired state of CommandExec
type CommandExecSpec struct {
	// Input is the input of the command
	Input string `json:"input,omitempty"`
}

// CommandExecStatus defines the observed state of CommandExec
type CommandExecStatus struct {
	// Conditions represent the latest available observations of an object's state
	// +optional
	Conditions []condition.Condition `json:"conditions,omitempty"`

	// Output is the output of the command
	Output string `json:"output,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CommandExec is the Schema for the commandexecs API
type CommandExec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CommandExecSpec   `json:"spec,omitempty"`
	Status CommandExecStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CommandExecList contains a list of CommandExec
type CommandExecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CommandExec `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CommandExec{}, &CommandExecList{})
}
