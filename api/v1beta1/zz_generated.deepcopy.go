//go:build !ignore_autogenerated

/*
Copyright 2024.

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
func (in *ProjectDevelopmentStream) DeepCopyInto(out *ProjectDevelopmentStream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectDevelopmentStream.
func (in *ProjectDevelopmentStream) DeepCopy() *ProjectDevelopmentStream {
	if in == nil {
		return nil
	}
	out := new(ProjectDevelopmentStream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectDevelopmentStream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectDevelopmentStreamList) DeepCopyInto(out *ProjectDevelopmentStreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProjectDevelopmentStream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectDevelopmentStreamList.
func (in *ProjectDevelopmentStreamList) DeepCopy() *ProjectDevelopmentStreamList {
	if in == nil {
		return nil
	}
	out := new(ProjectDevelopmentStreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectDevelopmentStreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectDevelopmentStreamSpec) DeepCopyInto(out *ProjectDevelopmentStreamSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]UnstructuredObj, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectDevelopmentStreamSpec.
func (in *ProjectDevelopmentStreamSpec) DeepCopy() *ProjectDevelopmentStreamSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectDevelopmentStreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectDevelopmentStreamStatus) DeepCopyInto(out *ProjectDevelopmentStreamStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectDevelopmentStreamStatus.
func (in *ProjectDevelopmentStreamStatus) DeepCopy() *ProjectDevelopmentStreamStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectDevelopmentStreamStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnstructuredObj) DeepCopyInto(out *UnstructuredObj) {
	*out = *in
	in.Unstructured.DeepCopyInto(&out.Unstructured)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnstructuredObj.
func (in *UnstructuredObj) DeepCopy() *UnstructuredObj {
	if in == nil {
		return nil
	}
	out := new(UnstructuredObj)
	in.DeepCopyInto(out)
	return out
}
