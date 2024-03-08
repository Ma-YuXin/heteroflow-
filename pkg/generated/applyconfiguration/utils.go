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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1beta1 "heterflow/api/v1beta1"
	apiv1beta1 "heterflow/pkg/generated/applyconfiguration/api/v1beta1"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=core, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("RealtimeJobDispatcher"):
		return &apiv1beta1.RealtimeJobDispatcherApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("RealtimeJobDispatcherSpec"):
		return &apiv1beta1.RealtimeJobDispatcherSpecApplyConfiguration{}

	}
	return nil
}
