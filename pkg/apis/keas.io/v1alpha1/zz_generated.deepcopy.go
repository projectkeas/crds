//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Auto-generated for github.com/projectkeas/crds
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngestionPolicy) DeepCopyInto(out *IngestionPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngestionPolicy.
func (in *IngestionPolicy) DeepCopy() *IngestionPolicy {
	if in == nil {
		return nil
	}
	out := new(IngestionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngestionPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngestionPolicyList) DeepCopyInto(out *IngestionPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngestionPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngestionPolicyList.
func (in *IngestionPolicyList) DeepCopy() *IngestionPolicyList {
	if in == nil {
		return nil
	}
	out := new(IngestionPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngestionPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngestionPolicySpecification) DeepCopyInto(out *IngestionPolicySpecification) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngestionPolicySpecification.
func (in *IngestionPolicySpecification) DeepCopy() *IngestionPolicySpecification {
	if in == nil {
		return nil
	}
	out := new(IngestionPolicySpecification)
	in.DeepCopyInto(out)
	return out
}
