//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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
	"github.com/hetznercloud/hcloud-go/hcloud"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudLoadBalancerSpec) DeepCopyInto(out *HCloudLoadBalancerSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]LoadBalancerServiceSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudLoadBalancerSpec.
func (in *HCloudLoadBalancerSpec) DeepCopy() *HCloudLoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(HCloudLoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudLoadBalancerStatus) DeepCopyInto(out *HCloudLoadBalancerStatus) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudLoadBalancerStatus.
func (in *HCloudLoadBalancerStatus) DeepCopy() *HCloudLoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(HCloudLoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachine) DeepCopyInto(out *HCloudMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachine.
func (in *HCloudMachine) DeepCopy() *HCloudMachine {
	if in == nil {
		return nil
	}
	out := new(HCloudMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HCloudMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineList) DeepCopyInto(out *HCloudMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HCloudMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineList.
func (in *HCloudMachineList) DeepCopy() *HCloudMachineList {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HCloudMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineSpec) DeepCopyInto(out *HCloudMachineSpec) {
	*out = *in
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.SSHKey != nil {
		in, out := &in.SSHKey, &out.SSHKey
		*out = make([]HCloudSSHKeySpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PlacementGroupName != nil {
		in, out := &in.PlacementGroupName, &out.PlacementGroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineSpec.
func (in *HCloudMachineSpec) DeepCopy() *HCloudMachineSpec {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineStatus) DeepCopyInto(out *HCloudMachineStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		*out = new(hcloud.ServerStatus)
		**out = **in
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineStatus.
func (in *HCloudMachineStatus) DeepCopy() *HCloudMachineStatus {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineTemplate) DeepCopyInto(out *HCloudMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineTemplate.
func (in *HCloudMachineTemplate) DeepCopy() *HCloudMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HCloudMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineTemplateList) DeepCopyInto(out *HCloudMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HCloudMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineTemplateList.
func (in *HCloudMachineTemplateList) DeepCopy() *HCloudMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HCloudMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineTemplateResource) DeepCopyInto(out *HCloudMachineTemplateResource) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineTemplateResource.
func (in *HCloudMachineTemplateResource) DeepCopy() *HCloudMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudMachineTemplateSpec) DeepCopyInto(out *HCloudMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudMachineTemplateSpec.
func (in *HCloudMachineTemplateSpec) DeepCopy() *HCloudMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HCloudMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudPlacementGroupSpec) DeepCopyInto(out *HCloudPlacementGroupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudPlacementGroupSpec.
func (in *HCloudPlacementGroupSpec) DeepCopy() *HCloudPlacementGroupSpec {
	if in == nil {
		return nil
	}
	out := new(HCloudPlacementGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudPlacementGroupStatus) DeepCopyInto(out *HCloudPlacementGroupStatus) {
	*out = *in
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudPlacementGroupStatus.
func (in *HCloudPlacementGroupStatus) DeepCopy() *HCloudPlacementGroupStatus {
	if in == nil {
		return nil
	}
	out := new(HCloudPlacementGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HCloudSSHKeySpec) DeepCopyInto(out *HCloudSSHKeySpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HCloudSSHKeySpec.
func (in *HCloudSSHKeySpec) DeepCopy() *HCloudSSHKeySpec {
	if in == nil {
		return nil
	}
	out := new(HCloudSSHKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerCluster) DeepCopyInto(out *HetznerCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerCluster.
func (in *HetznerCluster) DeepCopy() *HetznerCluster {
	if in == nil {
		return nil
	}
	out := new(HetznerCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterList) DeepCopyInto(out *HetznerClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HetznerCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterList.
func (in *HetznerClusterList) DeepCopy() *HetznerClusterList {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterSpec) DeepCopyInto(out *HetznerClusterSpec) {
	*out = *in
	in.NetworkSpec.DeepCopyInto(&out.NetworkSpec)
	if in.ControlPlaneRegion != nil {
		in, out := &in.ControlPlaneRegion, &out.ControlPlaneRegion
		*out = make([]HCloudRegion, len(*in))
		copy(*out, *in)
	}
	if in.SSHKey != nil {
		in, out := &in.SSHKey, &out.SSHKey
		*out = make([]HCloudSSHKeySpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ControlPlaneEndpoint != nil {
		in, out := &in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint
		*out = new(apiv1beta1.APIEndpoint)
		**out = **in
	}
	in.ControlPlaneLoadBalancer.DeepCopyInto(&out.ControlPlaneLoadBalancer)
	if in.PlacementGroupSpec != nil {
		in, out := &in.PlacementGroupSpec, &out.PlacementGroupSpec
		*out = make([]HCloudPlacementGroupSpec, len(*in))
		copy(*out, *in)
	}
	if in.HCloudTokenRef != nil {
		in, out := &in.HCloudTokenRef, &out.HCloudTokenRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterSpec.
func (in *HetznerClusterSpec) DeepCopy() *HetznerClusterSpec {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterStatus) DeepCopyInto(out *HetznerClusterStatus) {
	*out = *in
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(NetworkStatus)
		(*in).DeepCopyInto(*out)
	}
	in.ControlPlaneLoadBalancer.DeepCopyInto(&out.ControlPlaneLoadBalancer)
	if in.PlacementGroup != nil {
		in, out := &in.PlacementGroup, &out.PlacementGroup
		*out = make([]HCloudPlacementGroupStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(apiv1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterStatus.
func (in *HetznerClusterStatus) DeepCopy() *HetznerClusterStatus {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterTemplate) DeepCopyInto(out *HetznerClusterTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterTemplate.
func (in *HetznerClusterTemplate) DeepCopy() *HetznerClusterTemplate {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerClusterTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterTemplateList) DeepCopyInto(out *HetznerClusterTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HetznerClusterTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterTemplateList.
func (in *HetznerClusterTemplateList) DeepCopy() *HetznerClusterTemplateList {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HetznerClusterTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterTemplateResource) DeepCopyInto(out *HetznerClusterTemplateResource) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterTemplateResource.
func (in *HetznerClusterTemplateResource) DeepCopy() *HetznerClusterTemplateResource {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HetznerClusterTemplateSpec) DeepCopyInto(out *HetznerClusterTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HetznerClusterTemplateSpec.
func (in *HetznerClusterTemplateSpec) DeepCopy() *HetznerClusterTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(HetznerClusterTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerServiceSpec) DeepCopyInto(out *LoadBalancerServiceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerServiceSpec.
func (in *LoadBalancerServiceSpec) DeepCopy() *LoadBalancerServiceSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(Subnets, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkStatus) DeepCopyInto(out *NetworkStatus) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(Subnets, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AttachedServer != nil {
		in, out := &in.AttachedServer, &out.AttachedServer
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkStatus.
func (in *NetworkStatus) DeepCopy() *NetworkStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Subnets) DeepCopyInto(out *Subnets) {
	{
		in := &in
		*out = make(Subnets, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnets.
func (in Subnets) DeepCopy() Subnets {
	if in == nil {
		return nil
	}
	out := new(Subnets)
	in.DeepCopyInto(out)
	return *out
}