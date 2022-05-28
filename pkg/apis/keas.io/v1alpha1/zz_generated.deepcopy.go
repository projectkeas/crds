//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Auto-generated for github.com/projectkeas/crds
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventType) DeepCopyInto(out *EventType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventType.
func (in *EventType) DeepCopy() *EventType {
	if in == nil {
		return nil
	}
	out := new(EventType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTypeList) DeepCopyInto(out *EventTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTypeList.
func (in *EventTypeList) DeepCopy() *EventTypeList {
	if in == nil {
		return nil
	}
	out := new(EventTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTypeSpecification) DeepCopyInto(out *EventTypeSpecification) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubTypes != nil {
		in, out := &in.SubTypes, &out.SubTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTypeSpecification.
func (in *EventTypeSpecification) DeepCopy() *EventTypeSpecification {
	if in == nil {
		return nil
	}
	out := new(EventTypeSpecification)
	in.DeepCopyInto(out)
	return out
}

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
func (in *IngestionPolicyDefaults) DeepCopyInto(out *IngestionPolicyDefaults) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngestionPolicyDefaults.
func (in *IngestionPolicyDefaults) DeepCopy() *IngestionPolicyDefaults {
	if in == nil {
		return nil
	}
	out := new(IngestionPolicyDefaults)
	in.DeepCopyInto(out)
	return out
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
	out.Defaults = in.Defaults
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
