//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealtimeJobDispatcher) DeepCopyInto(out *RealtimeJobDispatcher) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealtimeJobDispatcher.
func (in *RealtimeJobDispatcher) DeepCopy() *RealtimeJobDispatcher {
	if in == nil {
		return nil
	}
	out := new(RealtimeJobDispatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RealtimeJobDispatcher) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealtimeJobDispatcherList) DeepCopyInto(out *RealtimeJobDispatcherList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RealtimeJobDispatcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealtimeJobDispatcherList.
func (in *RealtimeJobDispatcherList) DeepCopy() *RealtimeJobDispatcherList {
	if in == nil {
		return nil
	}
	out := new(RealtimeJobDispatcherList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RealtimeJobDispatcherList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealtimeJobDispatcherSpec) DeepCopyInto(out *RealtimeJobDispatcherSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealtimeJobDispatcherSpec.
func (in *RealtimeJobDispatcherSpec) DeepCopy() *RealtimeJobDispatcherSpec {
	if in == nil {
		return nil
	}
	out := new(RealtimeJobDispatcherSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RealtimeJobDispatcherStatus) DeepCopyInto(out *RealtimeJobDispatcherStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RealtimeJobDispatcherStatus.
func (in *RealtimeJobDispatcherStatus) DeepCopy() *RealtimeJobDispatcherStatus {
	if in == nil {
		return nil
	}
	out := new(RealtimeJobDispatcherStatus)
	in.DeepCopyInto(out)
	return out
}
