// +build !ignore_autogenerated

/*
Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.

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

// autogenerated by controller-gen object, do not modify manually

package v1

import (
	common "github.com/aws/amazon-sagemaker-operator-for-k8s/api/v1/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostingDeployment) DeepCopyInto(out *HostingDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostingDeployment.
func (in *HostingDeployment) DeepCopy() *HostingDeployment {
	if in == nil {
		return nil
	}
	out := new(HostingDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostingDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostingDeploymentList) DeepCopyInto(out *HostingDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HostingDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostingDeploymentList.
func (in *HostingDeploymentList) DeepCopy() *HostingDeploymentList {
	if in == nil {
		return nil
	}
	out := new(HostingDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostingDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostingDeploymentSpec) DeepCopyInto(out *HostingDeploymentSpec) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SageMakerEndpoint != nil {
		in, out := &in.SageMakerEndpoint, &out.SageMakerEndpoint
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyId != nil {
		in, out := &in.KmsKeyId, &out.KmsKeyId
		*out = new(string)
		**out = **in
	}
	if in.ProductionVariants != nil {
		in, out := &in.ProductionVariants, &out.ProductionVariants
		*out = make([]common.ProductionVariant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Models != nil {
		in, out := &in.Models, &out.Models
		*out = make([]common.Model, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]common.Tag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostingDeploymentSpec.
func (in *HostingDeploymentSpec) DeepCopy() *HostingDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(HostingDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostingDeploymentStatus) DeepCopyInto(out *HostingDeploymentStatus) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = new(metav1.Time)
		(*in).DeepCopyInto(*out)
	}
	if in.LastCheckTime != nil {
		in, out := &in.LastCheckTime, &out.LastCheckTime
		*out = new(metav1.Time)
		(*in).DeepCopyInto(*out)
	}
	if in.LastModifiedTime != nil {
		in, out := &in.LastModifiedTime, &out.LastModifiedTime
		*out = new(metav1.Time)
		(*in).DeepCopyInto(*out)
	}
	if in.ProductionVariants != nil {
		in, out := &in.ProductionVariants, &out.ProductionVariants
		*out = make([]*common.ProductionVariantSummary, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(common.ProductionVariantSummary)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ModelNames != nil {
		in, out := &in.ModelNames, &out.ModelNames
		*out = make([]*common.KeyValuePair, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(common.KeyValuePair)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostingDeploymentStatus.
func (in *HostingDeploymentStatus) DeepCopy() *HostingDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(HostingDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}
