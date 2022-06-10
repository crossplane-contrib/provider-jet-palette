//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ippool) DeepCopyInto(out *Ippool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ippool.
func (in *Ippool) DeepCopy() *Ippool {
	if in == nil {
		return nil
	}
	out := new(Ippool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ippool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IppoolList) DeepCopyInto(out *IppoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ippool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IppoolList.
func (in *IppoolList) DeepCopy() *IppoolList {
	if in == nil {
		return nil
	}
	out := new(IppoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IppoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IppoolObservation) DeepCopyInto(out *IppoolObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IppoolObservation.
func (in *IppoolObservation) DeepCopy() *IppoolObservation {
	if in == nil {
		return nil
	}
	out := new(IppoolObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IppoolParameters) DeepCopyInto(out *IppoolParameters) {
	*out = *in
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.IPEndRange != nil {
		in, out := &in.IPEndRange, &out.IPEndRange
		*out = new(string)
		**out = **in
	}
	if in.IPStartRange != nil {
		in, out := &in.IPStartRange, &out.IPStartRange
		*out = new(string)
		**out = **in
	}
	if in.NameserverAddresses != nil {
		in, out := &in.NameserverAddresses, &out.NameserverAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NameserverSearchSuffix != nil {
		in, out := &in.NameserverSearchSuffix, &out.NameserverSearchSuffix
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NetworkType != nil {
		in, out := &in.NetworkType, &out.NetworkType
		*out = new(string)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(float64)
		**out = **in
	}
	if in.PrivateCloudGatewayID != nil {
		in, out := &in.PrivateCloudGatewayID, &out.PrivateCloudGatewayID
		*out = new(string)
		**out = **in
	}
	if in.RestrictToSingleCluster != nil {
		in, out := &in.RestrictToSingleCluster, &out.RestrictToSingleCluster
		*out = new(bool)
		**out = **in
	}
	if in.SubnetCidr != nil {
		in, out := &in.SubnetCidr, &out.SubnetCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IppoolParameters.
func (in *IppoolParameters) DeepCopy() *IppoolParameters {
	if in == nil {
		return nil
	}
	out := new(IppoolParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IppoolSpec) DeepCopyInto(out *IppoolSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IppoolSpec.
func (in *IppoolSpec) DeepCopy() *IppoolSpec {
	if in == nil {
		return nil
	}
	out := new(IppoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IppoolStatus) DeepCopyInto(out *IppoolStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IppoolStatus.
func (in *IppoolStatus) DeepCopy() *IppoolStatus {
	if in == nil {
		return nil
	}
	out := new(IppoolStatus)
	in.DeepCopyInto(out)
	return out
}
