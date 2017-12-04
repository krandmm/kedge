// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package build

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BinaryBuildRequestOptions, InType: reflect.TypeOf(&BinaryBuildRequestOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BinaryBuildSource, InType: reflect.TypeOf(&BinaryBuildSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BitbucketWebHookCause, InType: reflect.TypeOf(&BitbucketWebHookCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_Build, InType: reflect.TypeOf(&Build{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildConfig, InType: reflect.TypeOf(&BuildConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildConfigList, InType: reflect.TypeOf(&BuildConfigList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildConfigSpec, InType: reflect.TypeOf(&BuildConfigSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildConfigStatus, InType: reflect.TypeOf(&BuildConfigStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildList, InType: reflect.TypeOf(&BuildList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildLog, InType: reflect.TypeOf(&BuildLog{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildLogOptions, InType: reflect.TypeOf(&BuildLogOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildOutput, InType: reflect.TypeOf(&BuildOutput{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildPostCommitSpec, InType: reflect.TypeOf(&BuildPostCommitSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildRequest, InType: reflect.TypeOf(&BuildRequest{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildSource, InType: reflect.TypeOf(&BuildSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildSpec, InType: reflect.TypeOf(&BuildSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildStatus, InType: reflect.TypeOf(&BuildStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildStatusOutput, InType: reflect.TypeOf(&BuildStatusOutput{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildStatusOutputTo, InType: reflect.TypeOf(&BuildStatusOutputTo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildStrategy, InType: reflect.TypeOf(&BuildStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildTriggerCause, InType: reflect.TypeOf(&BuildTriggerCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_BuildTriggerPolicy, InType: reflect.TypeOf(&BuildTriggerPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_CommonSpec, InType: reflect.TypeOf(&CommonSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_CommonWebHookCause, InType: reflect.TypeOf(&CommonWebHookCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_CustomBuildStrategy, InType: reflect.TypeOf(&CustomBuildStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_DockerBuildStrategy, InType: reflect.TypeOf(&DockerBuildStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_DockerStrategyOptions, InType: reflect.TypeOf(&DockerStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_GenericWebHookCause, InType: reflect.TypeOf(&GenericWebHookCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_GenericWebHookEvent, InType: reflect.TypeOf(&GenericWebHookEvent{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_GitBuildSource, InType: reflect.TypeOf(&GitBuildSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_GitHubWebHookCause, InType: reflect.TypeOf(&GitHubWebHookCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_GitInfo, InType: reflect.TypeOf(&GitInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_GitLabWebHookCause, InType: reflect.TypeOf(&GitLabWebHookCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_GitRefInfo, InType: reflect.TypeOf(&GitRefInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_GitSourceRevision, InType: reflect.TypeOf(&GitSourceRevision{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_ImageChangeCause, InType: reflect.TypeOf(&ImageChangeCause{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_ImageChangeTrigger, InType: reflect.TypeOf(&ImageChangeTrigger{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_ImageLabel, InType: reflect.TypeOf(&ImageLabel{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_ImageSource, InType: reflect.TypeOf(&ImageSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_ImageSourcePath, InType: reflect.TypeOf(&ImageSourcePath{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_JenkinsPipelineBuildStrategy, InType: reflect.TypeOf(&JenkinsPipelineBuildStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_ProxyConfig, InType: reflect.TypeOf(&ProxyConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_SecretBuildSource, InType: reflect.TypeOf(&SecretBuildSource{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_SecretSpec, InType: reflect.TypeOf(&SecretSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_SourceBuildStrategy, InType: reflect.TypeOf(&SourceBuildStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_SourceControlUser, InType: reflect.TypeOf(&SourceControlUser{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_SourceRevision, InType: reflect.TypeOf(&SourceRevision{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_SourceStrategyOptions, InType: reflect.TypeOf(&SourceStrategyOptions{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_StageInfo, InType: reflect.TypeOf(&StageInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_StepInfo, InType: reflect.TypeOf(&StepInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_build_WebHookTrigger, InType: reflect.TypeOf(&WebHookTrigger{})},
	)
}

// DeepCopy_build_BinaryBuildRequestOptions is an autogenerated deepcopy function.
func DeepCopy_build_BinaryBuildRequestOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BinaryBuildRequestOptions)
		out := out.(*BinaryBuildRequestOptions)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		return nil
	}
}

// DeepCopy_build_BinaryBuildSource is an autogenerated deepcopy function.
func DeepCopy_build_BinaryBuildSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BinaryBuildSource)
		out := out.(*BinaryBuildSource)
		*out = *in
		return nil
	}
}

// DeepCopy_build_BitbucketWebHookCause is an autogenerated deepcopy function.
func DeepCopy_build_BitbucketWebHookCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BitbucketWebHookCause)
		out := out.(*BitbucketWebHookCause)
		*out = *in
		if err := DeepCopy_build_CommonWebHookCause(&in.CommonWebHookCause, &out.CommonWebHookCause, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_build_Build is an autogenerated deepcopy function.
func DeepCopy_build_Build(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Build)
		out := out.(*Build)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_build_BuildSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_build_BuildStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_build_BuildConfig is an autogenerated deepcopy function.
func DeepCopy_build_BuildConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildConfig)
		out := out.(*BuildConfig)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_build_BuildConfigSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_build_BuildConfigList is an autogenerated deepcopy function.
func DeepCopy_build_BuildConfigList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildConfigList)
		out := out.(*BuildConfigList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]BuildConfig, len(*in))
			for i := range *in {
				if err := DeepCopy_build_BuildConfig(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_build_BuildConfigSpec is an autogenerated deepcopy function.
func DeepCopy_build_BuildConfigSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildConfigSpec)
		out := out.(*BuildConfigSpec)
		*out = *in
		if in.Triggers != nil {
			in, out := &in.Triggers, &out.Triggers
			*out = make([]BuildTriggerPolicy, len(*in))
			for i := range *in {
				if err := DeepCopy_build_BuildTriggerPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if err := DeepCopy_build_CommonSpec(&in.CommonSpec, &out.CommonSpec, c); err != nil {
			return err
		}
		if in.SuccessfulBuildsHistoryLimit != nil {
			in, out := &in.SuccessfulBuildsHistoryLimit, &out.SuccessfulBuildsHistoryLimit
			*out = new(int32)
			**out = **in
		}
		if in.FailedBuildsHistoryLimit != nil {
			in, out := &in.FailedBuildsHistoryLimit, &out.FailedBuildsHistoryLimit
			*out = new(int32)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_BuildConfigStatus is an autogenerated deepcopy function.
func DeepCopy_build_BuildConfigStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildConfigStatus)
		out := out.(*BuildConfigStatus)
		*out = *in
		return nil
	}
}

// DeepCopy_build_BuildList is an autogenerated deepcopy function.
func DeepCopy_build_BuildList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildList)
		out := out.(*BuildList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Build, len(*in))
			for i := range *in {
				if err := DeepCopy_build_Build(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_build_BuildLog is an autogenerated deepcopy function.
func DeepCopy_build_BuildLog(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildLog)
		out := out.(*BuildLog)
		*out = *in
		return nil
	}
}

// DeepCopy_build_BuildLogOptions is an autogenerated deepcopy function.
func DeepCopy_build_BuildLogOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildLogOptions)
		out := out.(*BuildLogOptions)
		*out = *in
		if in.SinceSeconds != nil {
			in, out := &in.SinceSeconds, &out.SinceSeconds
			*out = new(int64)
			**out = **in
		}
		if in.SinceTime != nil {
			in, out := &in.SinceTime, &out.SinceTime
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		}
		if in.TailLines != nil {
			in, out := &in.TailLines, &out.TailLines
			*out = new(int64)
			**out = **in
		}
		if in.LimitBytes != nil {
			in, out := &in.LimitBytes, &out.LimitBytes
			*out = new(int64)
			**out = **in
		}
		if in.Version != nil {
			in, out := &in.Version, &out.Version
			*out = new(int64)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_BuildOutput is an autogenerated deepcopy function.
func DeepCopy_build_BuildOutput(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildOutput)
		out := out.(*BuildOutput)
		*out = *in
		if in.To != nil {
			in, out := &in.To, &out.To
			*out = new(api.ObjectReference)
			**out = **in
		}
		if in.PushSecret != nil {
			in, out := &in.PushSecret, &out.PushSecret
			*out = new(api.LocalObjectReference)
			**out = **in
		}
		if in.ImageLabels != nil {
			in, out := &in.ImageLabels, &out.ImageLabels
			*out = make([]ImageLabel, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_build_BuildPostCommitSpec is an autogenerated deepcopy function.
func DeepCopy_build_BuildPostCommitSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildPostCommitSpec)
		out := out.(*BuildPostCommitSpec)
		*out = *in
		if in.Command != nil {
			in, out := &in.Command, &out.Command
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Args != nil {
			in, out := &in.Args, &out.Args
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_build_BuildRequest is an autogenerated deepcopy function.
func DeepCopy_build_BuildRequest(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildRequest)
		out := out.(*BuildRequest)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Revision != nil {
			in, out := &in.Revision, &out.Revision
			*out = new(SourceRevision)
			if err := DeepCopy_build_SourceRevision(*in, *out, c); err != nil {
				return err
			}
		}
		if in.TriggeredByImage != nil {
			in, out := &in.TriggeredByImage, &out.TriggeredByImage
			*out = new(api.ObjectReference)
			**out = **in
		}
		if in.From != nil {
			in, out := &in.From, &out.From
			*out = new(api.ObjectReference)
			**out = **in
		}
		if in.Binary != nil {
			in, out := &in.Binary, &out.Binary
			*out = new(BinaryBuildSource)
			**out = **in
		}
		if in.LastVersion != nil {
			in, out := &in.LastVersion, &out.LastVersion
			*out = new(int64)
			**out = **in
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api.EnvVar, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.TriggeredBy != nil {
			in, out := &in.TriggeredBy, &out.TriggeredBy
			*out = make([]BuildTriggerCause, len(*in))
			for i := range *in {
				if err := DeepCopy_build_BuildTriggerCause(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.DockerStrategyOptions != nil {
			in, out := &in.DockerStrategyOptions, &out.DockerStrategyOptions
			*out = new(DockerStrategyOptions)
			if err := DeepCopy_build_DockerStrategyOptions(*in, *out, c); err != nil {
				return err
			}
		}
		if in.SourceStrategyOptions != nil {
			in, out := &in.SourceStrategyOptions, &out.SourceStrategyOptions
			*out = new(SourceStrategyOptions)
			if err := DeepCopy_build_SourceStrategyOptions(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_build_BuildSource is an autogenerated deepcopy function.
func DeepCopy_build_BuildSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildSource)
		out := out.(*BuildSource)
		*out = *in
		if in.Binary != nil {
			in, out := &in.Binary, &out.Binary
			*out = new(BinaryBuildSource)
			**out = **in
		}
		if in.Dockerfile != nil {
			in, out := &in.Dockerfile, &out.Dockerfile
			*out = new(string)
			**out = **in
		}
		if in.Git != nil {
			in, out := &in.Git, &out.Git
			*out = new(GitBuildSource)
			if err := DeepCopy_build_GitBuildSource(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Images != nil {
			in, out := &in.Images, &out.Images
			*out = make([]ImageSource, len(*in))
			for i := range *in {
				if err := DeepCopy_build_ImageSource(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.SourceSecret != nil {
			in, out := &in.SourceSecret, &out.SourceSecret
			*out = new(api.LocalObjectReference)
			**out = **in
		}
		if in.Secrets != nil {
			in, out := &in.Secrets, &out.Secrets
			*out = make([]SecretBuildSource, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_build_BuildSpec is an autogenerated deepcopy function.
func DeepCopy_build_BuildSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildSpec)
		out := out.(*BuildSpec)
		*out = *in
		if err := DeepCopy_build_CommonSpec(&in.CommonSpec, &out.CommonSpec, c); err != nil {
			return err
		}
		if in.TriggeredBy != nil {
			in, out := &in.TriggeredBy, &out.TriggeredBy
			*out = make([]BuildTriggerCause, len(*in))
			for i := range *in {
				if err := DeepCopy_build_BuildTriggerCause(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_build_BuildStatus is an autogenerated deepcopy function.
func DeepCopy_build_BuildStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildStatus)
		out := out.(*BuildStatus)
		*out = *in
		if in.StartTimestamp != nil {
			in, out := &in.StartTimestamp, &out.StartTimestamp
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		}
		if in.CompletionTimestamp != nil {
			in, out := &in.CompletionTimestamp, &out.CompletionTimestamp
			*out = new(v1.Time)
			**out = (*in).DeepCopy()
		}
		if in.Config != nil {
			in, out := &in.Config, &out.Config
			*out = new(api.ObjectReference)
			**out = **in
		}
		if err := DeepCopy_build_BuildStatusOutput(&in.Output, &out.Output, c); err != nil {
			return err
		}
		if in.Stages != nil {
			in, out := &in.Stages, &out.Stages
			*out = make([]StageInfo, len(*in))
			for i := range *in {
				if err := DeepCopy_build_StageInfo(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_build_BuildStatusOutput is an autogenerated deepcopy function.
func DeepCopy_build_BuildStatusOutput(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildStatusOutput)
		out := out.(*BuildStatusOutput)
		*out = *in
		if in.To != nil {
			in, out := &in.To, &out.To
			*out = new(BuildStatusOutputTo)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_BuildStatusOutputTo is an autogenerated deepcopy function.
func DeepCopy_build_BuildStatusOutputTo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildStatusOutputTo)
		out := out.(*BuildStatusOutputTo)
		*out = *in
		return nil
	}
}

// DeepCopy_build_BuildStrategy is an autogenerated deepcopy function.
func DeepCopy_build_BuildStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildStrategy)
		out := out.(*BuildStrategy)
		*out = *in
		if in.DockerStrategy != nil {
			in, out := &in.DockerStrategy, &out.DockerStrategy
			*out = new(DockerBuildStrategy)
			if err := DeepCopy_build_DockerBuildStrategy(*in, *out, c); err != nil {
				return err
			}
		}
		if in.SourceStrategy != nil {
			in, out := &in.SourceStrategy, &out.SourceStrategy
			*out = new(SourceBuildStrategy)
			if err := DeepCopy_build_SourceBuildStrategy(*in, *out, c); err != nil {
				return err
			}
		}
		if in.CustomStrategy != nil {
			in, out := &in.CustomStrategy, &out.CustomStrategy
			*out = new(CustomBuildStrategy)
			if err := DeepCopy_build_CustomBuildStrategy(*in, *out, c); err != nil {
				return err
			}
		}
		if in.JenkinsPipelineStrategy != nil {
			in, out := &in.JenkinsPipelineStrategy, &out.JenkinsPipelineStrategy
			*out = new(JenkinsPipelineBuildStrategy)
			if err := DeepCopy_build_JenkinsPipelineBuildStrategy(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_build_BuildTriggerCause is an autogenerated deepcopy function.
func DeepCopy_build_BuildTriggerCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildTriggerCause)
		out := out.(*BuildTriggerCause)
		*out = *in
		if in.GenericWebHook != nil {
			in, out := &in.GenericWebHook, &out.GenericWebHook
			*out = new(GenericWebHookCause)
			if err := DeepCopy_build_GenericWebHookCause(*in, *out, c); err != nil {
				return err
			}
		}
		if in.GitHubWebHook != nil {
			in, out := &in.GitHubWebHook, &out.GitHubWebHook
			*out = new(GitHubWebHookCause)
			if err := DeepCopy_build_GitHubWebHookCause(*in, *out, c); err != nil {
				return err
			}
		}
		if in.ImageChangeBuild != nil {
			in, out := &in.ImageChangeBuild, &out.ImageChangeBuild
			*out = new(ImageChangeCause)
			if err := DeepCopy_build_ImageChangeCause(*in, *out, c); err != nil {
				return err
			}
		}
		if in.GitLabWebHook != nil {
			in, out := &in.GitLabWebHook, &out.GitLabWebHook
			*out = new(GitLabWebHookCause)
			if err := DeepCopy_build_GitLabWebHookCause(*in, *out, c); err != nil {
				return err
			}
		}
		if in.BitbucketWebHook != nil {
			in, out := &in.BitbucketWebHook, &out.BitbucketWebHook
			*out = new(BitbucketWebHookCause)
			if err := DeepCopy_build_BitbucketWebHookCause(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_build_BuildTriggerPolicy is an autogenerated deepcopy function.
func DeepCopy_build_BuildTriggerPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BuildTriggerPolicy)
		out := out.(*BuildTriggerPolicy)
		*out = *in
		if in.GitHubWebHook != nil {
			in, out := &in.GitHubWebHook, &out.GitHubWebHook
			*out = new(WebHookTrigger)
			**out = **in
		}
		if in.GenericWebHook != nil {
			in, out := &in.GenericWebHook, &out.GenericWebHook
			*out = new(WebHookTrigger)
			**out = **in
		}
		if in.ImageChange != nil {
			in, out := &in.ImageChange, &out.ImageChange
			*out = new(ImageChangeTrigger)
			if err := DeepCopy_build_ImageChangeTrigger(*in, *out, c); err != nil {
				return err
			}
		}
		if in.GitLabWebHook != nil {
			in, out := &in.GitLabWebHook, &out.GitLabWebHook
			*out = new(WebHookTrigger)
			**out = **in
		}
		if in.BitbucketWebHook != nil {
			in, out := &in.BitbucketWebHook, &out.BitbucketWebHook
			*out = new(WebHookTrigger)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_CommonSpec is an autogenerated deepcopy function.
func DeepCopy_build_CommonSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CommonSpec)
		out := out.(*CommonSpec)
		*out = *in
		if err := DeepCopy_build_BuildSource(&in.Source, &out.Source, c); err != nil {
			return err
		}
		if in.Revision != nil {
			in, out := &in.Revision, &out.Revision
			*out = new(SourceRevision)
			if err := DeepCopy_build_SourceRevision(*in, *out, c); err != nil {
				return err
			}
		}
		if err := DeepCopy_build_BuildStrategy(&in.Strategy, &out.Strategy, c); err != nil {
			return err
		}
		if err := DeepCopy_build_BuildOutput(&in.Output, &out.Output, c); err != nil {
			return err
		}
		if err := api.DeepCopy_api_ResourceRequirements(&in.Resources, &out.Resources, c); err != nil {
			return err
		}
		if err := DeepCopy_build_BuildPostCommitSpec(&in.PostCommit, &out.PostCommit, c); err != nil {
			return err
		}
		if in.CompletionDeadlineSeconds != nil {
			in, out := &in.CompletionDeadlineSeconds, &out.CompletionDeadlineSeconds
			*out = new(int64)
			**out = **in
		}
		if in.NodeSelector != nil {
			in, out := &in.NodeSelector, &out.NodeSelector
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopy_build_CommonWebHookCause is an autogenerated deepcopy function.
func DeepCopy_build_CommonWebHookCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CommonWebHookCause)
		out := out.(*CommonWebHookCause)
		*out = *in
		if in.Revision != nil {
			in, out := &in.Revision, &out.Revision
			*out = new(SourceRevision)
			if err := DeepCopy_build_SourceRevision(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_build_CustomBuildStrategy is an autogenerated deepcopy function.
func DeepCopy_build_CustomBuildStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*CustomBuildStrategy)
		out := out.(*CustomBuildStrategy)
		*out = *in
		if in.PullSecret != nil {
			in, out := &in.PullSecret, &out.PullSecret
			*out = new(api.LocalObjectReference)
			**out = **in
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api.EnvVar, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Secrets != nil {
			in, out := &in.Secrets, &out.Secrets
			*out = make([]SecretSpec, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_build_DockerBuildStrategy is an autogenerated deepcopy function.
func DeepCopy_build_DockerBuildStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DockerBuildStrategy)
		out := out.(*DockerBuildStrategy)
		*out = *in
		if in.From != nil {
			in, out := &in.From, &out.From
			*out = new(api.ObjectReference)
			**out = **in
		}
		if in.PullSecret != nil {
			in, out := &in.PullSecret, &out.PullSecret
			*out = new(api.LocalObjectReference)
			**out = **in
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api.EnvVar, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.BuildArgs != nil {
			in, out := &in.BuildArgs, &out.BuildArgs
			*out = make([]api.EnvVar, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.ImageOptimizationPolicy != nil {
			in, out := &in.ImageOptimizationPolicy, &out.ImageOptimizationPolicy
			*out = new(ImageOptimizationPolicy)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_DockerStrategyOptions is an autogenerated deepcopy function.
func DeepCopy_build_DockerStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*DockerStrategyOptions)
		out := out.(*DockerStrategyOptions)
		*out = *in
		if in.BuildArgs != nil {
			in, out := &in.BuildArgs, &out.BuildArgs
			*out = make([]api.EnvVar, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.NoCache != nil {
			in, out := &in.NoCache, &out.NoCache
			*out = new(bool)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_GenericWebHookCause is an autogenerated deepcopy function.
func DeepCopy_build_GenericWebHookCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GenericWebHookCause)
		out := out.(*GenericWebHookCause)
		*out = *in
		if in.Revision != nil {
			in, out := &in.Revision, &out.Revision
			*out = new(SourceRevision)
			if err := DeepCopy_build_SourceRevision(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_build_GenericWebHookEvent is an autogenerated deepcopy function.
func DeepCopy_build_GenericWebHookEvent(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GenericWebHookEvent)
		out := out.(*GenericWebHookEvent)
		*out = *in
		if in.Git != nil {
			in, out := &in.Git, &out.Git
			*out = new(GitInfo)
			if err := DeepCopy_build_GitInfo(*in, *out, c); err != nil {
				return err
			}
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api.EnvVar, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.DockerStrategyOptions != nil {
			in, out := &in.DockerStrategyOptions, &out.DockerStrategyOptions
			*out = new(DockerStrategyOptions)
			if err := DeepCopy_build_DockerStrategyOptions(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_build_GitBuildSource is an autogenerated deepcopy function.
func DeepCopy_build_GitBuildSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GitBuildSource)
		out := out.(*GitBuildSource)
		*out = *in
		if err := DeepCopy_build_ProxyConfig(&in.ProxyConfig, &out.ProxyConfig, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_build_GitHubWebHookCause is an autogenerated deepcopy function.
func DeepCopy_build_GitHubWebHookCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GitHubWebHookCause)
		out := out.(*GitHubWebHookCause)
		*out = *in
		if in.Revision != nil {
			in, out := &in.Revision, &out.Revision
			*out = new(SourceRevision)
			if err := DeepCopy_build_SourceRevision(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopy_build_GitInfo is an autogenerated deepcopy function.
func DeepCopy_build_GitInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GitInfo)
		out := out.(*GitInfo)
		*out = *in
		if err := DeepCopy_build_GitBuildSource(&in.GitBuildSource, &out.GitBuildSource, c); err != nil {
			return err
		}
		if in.Refs != nil {
			in, out := &in.Refs, &out.Refs
			*out = make([]GitRefInfo, len(*in))
			for i := range *in {
				if err := DeepCopy_build_GitRefInfo(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_build_GitLabWebHookCause is an autogenerated deepcopy function.
func DeepCopy_build_GitLabWebHookCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GitLabWebHookCause)
		out := out.(*GitLabWebHookCause)
		*out = *in
		if err := DeepCopy_build_CommonWebHookCause(&in.CommonWebHookCause, &out.CommonWebHookCause, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_build_GitRefInfo is an autogenerated deepcopy function.
func DeepCopy_build_GitRefInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GitRefInfo)
		out := out.(*GitRefInfo)
		*out = *in
		if err := DeepCopy_build_GitBuildSource(&in.GitBuildSource, &out.GitBuildSource, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_build_GitSourceRevision is an autogenerated deepcopy function.
func DeepCopy_build_GitSourceRevision(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GitSourceRevision)
		out := out.(*GitSourceRevision)
		*out = *in
		return nil
	}
}

// DeepCopy_build_ImageChangeCause is an autogenerated deepcopy function.
func DeepCopy_build_ImageChangeCause(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageChangeCause)
		out := out.(*ImageChangeCause)
		*out = *in
		if in.FromRef != nil {
			in, out := &in.FromRef, &out.FromRef
			*out = new(api.ObjectReference)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_ImageChangeTrigger is an autogenerated deepcopy function.
func DeepCopy_build_ImageChangeTrigger(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageChangeTrigger)
		out := out.(*ImageChangeTrigger)
		*out = *in
		if in.From != nil {
			in, out := &in.From, &out.From
			*out = new(api.ObjectReference)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_ImageLabel is an autogenerated deepcopy function.
func DeepCopy_build_ImageLabel(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageLabel)
		out := out.(*ImageLabel)
		*out = *in
		return nil
	}
}

// DeepCopy_build_ImageSource is an autogenerated deepcopy function.
func DeepCopy_build_ImageSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageSource)
		out := out.(*ImageSource)
		*out = *in
		if in.Paths != nil {
			in, out := &in.Paths, &out.Paths
			*out = make([]ImageSourcePath, len(*in))
			copy(*out, *in)
		}
		if in.PullSecret != nil {
			in, out := &in.PullSecret, &out.PullSecret
			*out = new(api.LocalObjectReference)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_ImageSourcePath is an autogenerated deepcopy function.
func DeepCopy_build_ImageSourcePath(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ImageSourcePath)
		out := out.(*ImageSourcePath)
		*out = *in
		return nil
	}
}

// DeepCopy_build_JenkinsPipelineBuildStrategy is an autogenerated deepcopy function.
func DeepCopy_build_JenkinsPipelineBuildStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*JenkinsPipelineBuildStrategy)
		out := out.(*JenkinsPipelineBuildStrategy)
		*out = *in
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api.EnvVar, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_build_ProxyConfig is an autogenerated deepcopy function.
func DeepCopy_build_ProxyConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ProxyConfig)
		out := out.(*ProxyConfig)
		*out = *in
		if in.HTTPProxy != nil {
			in, out := &in.HTTPProxy, &out.HTTPProxy
			*out = new(string)
			**out = **in
		}
		if in.HTTPSProxy != nil {
			in, out := &in.HTTPSProxy, &out.HTTPSProxy
			*out = new(string)
			**out = **in
		}
		if in.NoProxy != nil {
			in, out := &in.NoProxy, &out.NoProxy
			*out = new(string)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_SecretBuildSource is an autogenerated deepcopy function.
func DeepCopy_build_SecretBuildSource(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SecretBuildSource)
		out := out.(*SecretBuildSource)
		*out = *in
		return nil
	}
}

// DeepCopy_build_SecretSpec is an autogenerated deepcopy function.
func DeepCopy_build_SecretSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SecretSpec)
		out := out.(*SecretSpec)
		*out = *in
		return nil
	}
}

// DeepCopy_build_SourceBuildStrategy is an autogenerated deepcopy function.
func DeepCopy_build_SourceBuildStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SourceBuildStrategy)
		out := out.(*SourceBuildStrategy)
		*out = *in
		if in.PullSecret != nil {
			in, out := &in.PullSecret, &out.PullSecret
			*out = new(api.LocalObjectReference)
			**out = **in
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]api.EnvVar, len(*in))
			for i := range *in {
				if err := api.DeepCopy_api_EnvVar(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		if in.Incremental != nil {
			in, out := &in.Incremental, &out.Incremental
			*out = new(bool)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_SourceControlUser is an autogenerated deepcopy function.
func DeepCopy_build_SourceControlUser(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SourceControlUser)
		out := out.(*SourceControlUser)
		*out = *in
		return nil
	}
}

// DeepCopy_build_SourceRevision is an autogenerated deepcopy function.
func DeepCopy_build_SourceRevision(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SourceRevision)
		out := out.(*SourceRevision)
		*out = *in
		if in.Git != nil {
			in, out := &in.Git, &out.Git
			*out = new(GitSourceRevision)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_SourceStrategyOptions is an autogenerated deepcopy function.
func DeepCopy_build_SourceStrategyOptions(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SourceStrategyOptions)
		out := out.(*SourceStrategyOptions)
		*out = *in
		if in.Incremental != nil {
			in, out := &in.Incremental, &out.Incremental
			*out = new(bool)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_build_StageInfo is an autogenerated deepcopy function.
func DeepCopy_build_StageInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StageInfo)
		out := out.(*StageInfo)
		*out = *in
		out.StartTime = in.StartTime.DeepCopy()
		if in.Steps != nil {
			in, out := &in.Steps, &out.Steps
			*out = make([]StepInfo, len(*in))
			for i := range *in {
				if err := DeepCopy_build_StepInfo(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_build_StepInfo is an autogenerated deepcopy function.
func DeepCopy_build_StepInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*StepInfo)
		out := out.(*StepInfo)
		*out = *in
		out.StartTime = in.StartTime.DeepCopy()
		return nil
	}
}

// DeepCopy_build_WebHookTrigger is an autogenerated deepcopy function.
func DeepCopy_build_WebHookTrigger(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*WebHookTrigger)
		out := out.(*WebHookTrigger)
		*out = *in
		return nil
	}
}
