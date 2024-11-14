//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/envoyproxy/gateway/api/v1alpha1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apisv1 "sigs.k8s.io/gateway-api/apis/v1"
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMBackend) DeepCopyInto(out *LLMBackend) {
	*out = *in
	in.BackendRef.DeepCopyInto(&out.BackendRef)
	if in.ProviderPolicy != nil {
		in, out := &in.ProviderPolicy, &out.ProviderPolicy
		*out = new(LLMProviderPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.TrafficPolicy != nil {
		in, out := &in.TrafficPolicy, &out.TrafficPolicy
		*out = new(LLMTrafficPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMBackend.
func (in *LLMBackend) DeepCopy() *LLMBackend {
	if in == nil {
		return nil
	}
	out := new(LLMBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMPolicyRateLimitHeaderMatch) DeepCopyInto(out *LLMPolicyRateLimitHeaderMatch) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMPolicyRateLimitHeaderMatch.
func (in *LLMPolicyRateLimitHeaderMatch) DeepCopy() *LLMPolicyRateLimitHeaderMatch {
	if in == nil {
		return nil
	}
	out := new(LLMPolicyRateLimitHeaderMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMPolicyRateLimitMetadataMatch) DeepCopyInto(out *LLMPolicyRateLimitMetadataMatch) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultValue != nil {
		in, out := &in.DefaultValue, &out.DefaultValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMPolicyRateLimitMetadataMatch.
func (in *LLMPolicyRateLimitMetadataMatch) DeepCopy() *LLMPolicyRateLimitMetadataMatch {
	if in == nil {
		return nil
	}
	out := new(LLMPolicyRateLimitMetadataMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMPolicyRateLimitModelNameMatch) DeepCopyInto(out *LLMPolicyRateLimitModelNameMatch) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMPolicyRateLimitModelNameMatch.
func (in *LLMPolicyRateLimitModelNameMatch) DeepCopy() *LLMPolicyRateLimitModelNameMatch {
	if in == nil {
		return nil
	}
	out := new(LLMPolicyRateLimitModelNameMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMPolicyRateLimitValue) DeepCopyInto(out *LLMPolicyRateLimitValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMPolicyRateLimitValue.
func (in *LLMPolicyRateLimitValue) DeepCopy() *LLMPolicyRateLimitValue {
	if in == nil {
		return nil
	}
	out := new(LLMPolicyRateLimitValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMProviderAPIKey) DeepCopyInto(out *LLMProviderAPIKey) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(apisv1.SecretObjectReference)
		(*in).DeepCopyInto(*out)
	}
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMProviderAPIKey.
func (in *LLMProviderAPIKey) DeepCopy() *LLMProviderAPIKey {
	if in == nil {
		return nil
	}
	out := new(LLMProviderAPIKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMProviderAWSBedrock) DeepCopyInto(out *LLMProviderAWSBedrock) {
	*out = *in
	if in.SigningAlgorithm != nil {
		in, out := &in.SigningAlgorithm, &out.SigningAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.HostRewrite != nil {
		in, out := &in.HostRewrite, &out.HostRewrite
		*out = new(string)
		**out = **in
	}
	if in.InlineCredential != nil {
		in, out := &in.InlineCredential, &out.InlineCredential
		*out = new(LLMProviderAWSBedrockInlineCredential)
		**out = **in
	}
	if in.CredentialsFile != nil {
		in, out := &in.CredentialsFile, &out.CredentialsFile
		*out = new(LLMProviderAWSBedrockCredentialsFile)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMProviderAWSBedrock.
func (in *LLMProviderAWSBedrock) DeepCopy() *LLMProviderAWSBedrock {
	if in == nil {
		return nil
	}
	out := new(LLMProviderAWSBedrock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMProviderAWSBedrockCredentialsFile) DeepCopyInto(out *LLMProviderAWSBedrockCredentialsFile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMProviderAWSBedrockCredentialsFile.
func (in *LLMProviderAWSBedrockCredentialsFile) DeepCopy() *LLMProviderAWSBedrockCredentialsFile {
	if in == nil {
		return nil
	}
	out := new(LLMProviderAWSBedrockCredentialsFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMProviderAWSBedrockInlineCredential) DeepCopyInto(out *LLMProviderAWSBedrockInlineCredential) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMProviderAWSBedrockInlineCredential.
func (in *LLMProviderAWSBedrockInlineCredential) DeepCopy() *LLMProviderAWSBedrockInlineCredential {
	if in == nil {
		return nil
	}
	out := new(LLMProviderAWSBedrockInlineCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMProviderPolicy) DeepCopyInto(out *LLMProviderPolicy) {
	*out = *in
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(LLMProviderAPIKey)
		(*in).DeepCopyInto(*out)
	}
	if in.AWSBedrock != nil {
		in, out := &in.AWSBedrock, &out.AWSBedrock
		*out = new(LLMProviderAWSBedrock)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMProviderPolicy.
func (in *LLMProviderPolicy) DeepCopy() *LLMProviderPolicy {
	if in == nil {
		return nil
	}
	out := new(LLMProviderPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMRoute) DeepCopyInto(out *LLMRoute) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMRoute.
func (in *LLMRoute) DeepCopy() *LLMRoute {
	if in == nil {
		return nil
	}
	out := new(LLMRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LLMRoute) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMRouteExtProcConfig) DeepCopyInto(out *LLMRouteExtProcConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMRouteExtProcConfig.
func (in *LLMRouteExtProcConfig) DeepCopy() *LLMRouteExtProcConfig {
	if in == nil {
		return nil
	}
	out := new(LLMRouteExtProcConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMRouteList) DeepCopyInto(out *LLMRouteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LLMRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMRouteList.
func (in *LLMRouteList) DeepCopy() *LLMRouteList {
	if in == nil {
		return nil
	}
	out := new(LLMRouteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LLMRouteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMRouteSpec) DeepCopyInto(out *LLMRouteSpec) {
	*out = *in
	if in.ExtProcConfig != nil {
		in, out := &in.ExtProcConfig, &out.ExtProcConfig
		*out = new(LLMRouteExtProcConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetRefs != nil {
		in, out := &in.TargetRefs, &out.TargetRefs
		*out = make([]v1alpha2.LocalPolicyTargetReferenceWithSectionName, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]LLMBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtProc != nil {
		in, out := &in.ExtProc, &out.ExtProc
		*out = make([]apiv1alpha1.ExtProc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMRouteSpec.
func (in *LLMRouteSpec) DeepCopy() *LLMRouteSpec {
	if in == nil {
		return nil
	}
	out := new(LLMRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMTrafficPolicy) DeepCopyInto(out *LLMTrafficPolicy) {
	*out = *in
	if in.RateLimit != nil {
		in, out := &in.RateLimit, &out.RateLimit
		*out = new(LLMTrafficPolicyRateLimit)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMTrafficPolicy.
func (in *LLMTrafficPolicy) DeepCopy() *LLMTrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(LLMTrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMTrafficPolicyRateLimit) DeepCopyInto(out *LLMTrafficPolicyRateLimit) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]LLMTrafficPolicyRateLimitRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMTrafficPolicyRateLimit.
func (in *LLMTrafficPolicyRateLimit) DeepCopy() *LLMTrafficPolicyRateLimit {
	if in == nil {
		return nil
	}
	out := new(LLMTrafficPolicyRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LLMTrafficPolicyRateLimitRule) DeepCopyInto(out *LLMTrafficPolicyRateLimitRule) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]LLMPolicyRateLimitHeaderMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]LLMPolicyRateLimitMetadataMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make([]LLMPolicyRateLimitValue, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LLMTrafficPolicyRateLimitRule.
func (in *LLMTrafficPolicyRateLimitRule) DeepCopy() *LLMTrafficPolicyRateLimitRule {
	if in == nil {
		return nil
	}
	out := new(LLMTrafficPolicyRateLimitRule)
	in.DeepCopyInto(out)
	return out
}
