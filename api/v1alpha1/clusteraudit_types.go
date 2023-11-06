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

type ResourceTarget struct {
	Kind       string `json:"kind,omitempty"`
	Name       string `json:"name,omitempty"`
	ApiVersion string `json:"apiVersion,omitempty"`
}

// ClusterAuditSpec defines the desired state of ClusterAudit
type ClusterAuditSpec struct {
	Type string // ClusterUpgrade | CVEChecks | ClusterAudit | ClusterBackup | ClusterRestore | ClusterUpgradeRollback | ClusterUpgradeRollback
	// +optional
	// Foo is an example field of ClusterAudit. Edit clusteraudit_types.go to remove/update
	Resources map[string]*ResourceTarget `json:"resources,omitempty"`
}

// ClusterAuditStatus defines the observed state of ClusterAudit
type ClusterAuditStatus struct {
	Conditions []condition.Condition `json:"conditions,omitempty"`

	Output string `json:"output,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ClusterAudit is the Schema for the clusteraudits API
type ClusterAudit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterAuditSpec   `json:"spec,omitempty"`
	Status ClusterAuditStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterAuditList contains a list of ClusterAudit
type ClusterAuditList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterAudit `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterAudit{}, &ClusterAuditList{})
}
