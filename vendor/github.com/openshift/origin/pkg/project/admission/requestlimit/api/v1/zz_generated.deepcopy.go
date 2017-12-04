// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ProjectLimitBySelector, InType: reflect.TypeOf(&ProjectLimitBySelector{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_ProjectRequestLimitConfig, InType: reflect.TypeOf(&ProjectRequestLimitConfig{})},
	)
}

// DeepCopy_v1_ProjectLimitBySelector is an autogenerated deepcopy function.
func DeepCopy_v1_ProjectLimitBySelector(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ProjectLimitBySelector)
		out := out.(*ProjectLimitBySelector)
		*out = *in
		if in.Selector != nil {
			in, out := &in.Selector, &out.Selector
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.MaxProjects != nil {
			in, out := &in.MaxProjects, &out.MaxProjects
			*out = new(int)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_v1_ProjectRequestLimitConfig is an autogenerated deepcopy function.
func DeepCopy_v1_ProjectRequestLimitConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ProjectRequestLimitConfig)
		out := out.(*ProjectRequestLimitConfig)
		*out = *in
		if in.Limits != nil {
			in, out := &in.Limits, &out.Limits
			*out = make([]ProjectLimitBySelector, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_ProjectLimitBySelector(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.MaxProjectsForSystemUsers != nil {
			in, out := &in.MaxProjectsForSystemUsers, &out.MaxProjectsForSystemUsers
			*out = new(int)
			**out = **in
		}
		if in.MaxProjectsForServiceAccounts != nil {
			in, out := &in.MaxProjectsForServiceAccounts, &out.MaxProjectsForServiceAccounts
			*out = new(int)
			**out = **in
		}
		return nil
	}
}
