/*
Copyright The Kubernetes Authors.

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


package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RealtimeJobDispatcherSpec defines the desired state of RealtimeJobDispatcher
type RealtimeJobDispatcherSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RealtimeJobDispatcher. Edit realtimejobdispatcher_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// RealtimeJobDispatcherStatus defines the observed state of RealtimeJobDispatcher
type RealtimeJobDispatcherStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// +genclient
// +genclient:noStatus
// +k8s:openapi-gen=true

// RealtimeJobDispatcher is the Schema for the realtimejobdispatchers API
type RealtimeJobDispatcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RealtimeJobDispatcherSpec   `json:"spec,omitempty"`
	Status RealtimeJobDispatcherStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RealtimeJobDispatcherList contains a list of RealtimeJobDispatcher
type RealtimeJobDispatcherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RealtimeJobDispatcher `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RealtimeJobDispatcher{}, &RealtimeJobDispatcherList{})
}
