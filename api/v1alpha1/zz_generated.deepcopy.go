//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.FunctionRef != nil {
		in, out := &in.FunctionRef, &out.FunctionRef
		*out = new(string)
		**out = **in
	}
	if in.EventRef != nil {
		in, out := &in.EventRef, &out.EventRef
		*out = new(EventRef)
		(*in).DeepCopyInto(*out)
	}
	if in.SubFlowRef != nil {
		in, out := &in.SubFlowRef, &out.SubFlowRef
		*out = new(string)
		**out = **in
	}
	if in.RetryRef != nil {
		in, out := &in.RetryRef, &out.RetryRef
		*out = new(string)
		**out = **in
	}
	if in.NonRetryableErrors != nil {
		in, out := &in.NonRetryableErrors, &out.NonRetryableErrors
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.RetryableErrors != nil {
		in, out := &in.RetryableErrors, &out.RetryableErrors
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.ActionDataFilter != nil {
		in, out := &in.ActionDataFilter, &out.ActionDataFilter
		*out = new(ActionDataFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.Sleep != nil {
		in, out := &in.Sleep, &out.Sleep
		*out = new(Sleep)
		(*in).DeepCopyInto(*out)
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionDataFilter) DeepCopyInto(out *ActionDataFilter) {
	*out = *in
	if in.FromStateData != nil {
		in, out := &in.FromStateData, &out.FromStateData
		*out = new(string)
		**out = **in
	}
	if in.UseResults != nil {
		in, out := &in.UseResults, &out.UseResults
		*out = new(bool)
		**out = **in
	}
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = new(string)
		**out = **in
	}
	if in.ToStateData != nil {
		in, out := &in.ToStateData, &out.ToStateData
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionDataFilter.
func (in *ActionDataFilter) DeepCopy() *ActionDataFilter {
	if in == nil {
		return nil
	}
	out := new(ActionDataFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auth) DeepCopyInto(out *Auth) {
	*out = *in
	in.Properties.DeepCopyInto(&out.Properties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auth.
func (in *Auth) DeepCopy() *Auth {
	if in == nil {
		return nil
	}
	out := new(Auth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthProperties) DeepCopyInto(out *AuthProperties) {
	*out = *in
	in.Basic.DeepCopyInto(&out.Basic)
	in.Bearer.DeepCopyInto(&out.Bearer)
	in.Oauth2.DeepCopyInto(&out.Oauth2)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthProperties.
func (in *AuthProperties) DeepCopy() *AuthProperties {
	if in == nil {
		return nil
	}
	out := new(AuthProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicAuthProperties) DeepCopyInto(out *BasicAuthProperties) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]Metadata, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicAuthProperties.
func (in *BasicAuthProperties) DeepCopy() *BasicAuthProperties {
	if in == nil {
		return nil
	}
	out := new(BasicAuthProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BearerAuthProperties) DeepCopyInto(out *BearerAuthProperties) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]Metadata, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BearerAuthProperties.
func (in *BearerAuthProperties) DeepCopy() *BearerAuthProperties {
	if in == nil {
		return nil
	}
	out := new(BearerAuthProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompletionType) DeepCopyInto(out *CompletionType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompletionType.
func (in *CompletionType) DeepCopy() *CompletionType {
	if in == nil {
		return nil
	}
	out := new(CompletionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Constant) DeepCopyInto(out *Constant) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Constant.
func (in *Constant) DeepCopy() *Constant {
	if in == nil {
		return nil
	}
	out := new(Constant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCondition) DeepCopyInto(out *DataCondition) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]Metadata, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCondition.
func (in *DataCondition) DeepCopy() *DataCondition {
	if in == nil {
		return nil
	}
	out := new(DataCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Error) DeepCopyInto(out *Error) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Error.
func (in *Error) DeepCopy() *Error {
	if in == nil {
		return nil
	}
	out := new(Error)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Event) DeepCopyInto(out *Event) {
	*out = *in
	if in.Correlation != nil {
		in, out := &in.Correlation, &out.Correlation
		*out = make([]EventCorrelationRule, len(*in))
		copy(*out, *in)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]Metadata, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Event.
func (in *Event) DeepCopy() *Event {
	if in == nil {
		return nil
	}
	out := new(Event)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventCorrelationRule) DeepCopyInto(out *EventCorrelationRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventCorrelationRule.
func (in *EventCorrelationRule) DeepCopy() *EventCorrelationRule {
	if in == nil {
		return nil
	}
	out := new(EventCorrelationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventDataFilter) DeepCopyInto(out *EventDataFilter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventDataFilter.
func (in *EventDataFilter) DeepCopy() *EventDataFilter {
	if in == nil {
		return nil
	}
	out := new(EventDataFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventRef) DeepCopyInto(out *EventRef) {
	*out = *in
	if in.ConsumeEventRef != nil {
		in, out := &in.ConsumeEventRef, &out.ConsumeEventRef
		*out = new(string)
		**out = **in
	}
	if in.ConsumeEventTimeout != nil {
		in, out := &in.ConsumeEventTimeout, &out.ConsumeEventTimeout
		*out = new(string)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
	if in.ContextAttributes != nil {
		in, out := &in.ContextAttributes, &out.ContextAttributes
		*out = new(map[string]unstructured.Unstructured)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]unstructured.Unstructured, len(*in))
			for key, val := range *in {
				(*out)[key] = *val.DeepCopy()
			}
		}
	}
	if in.Invoke != nil {
		in, out := &in.Invoke, &out.Invoke
		*out = new(InvokeTye)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventRef.
func (in *EventRef) DeepCopy() *EventRef {
	if in == nil {
		return nil
	}
	out := new(EventRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]Metadata, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IterationMode) DeepCopyInto(out *IterationMode) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IterationMode.
func (in *IterationMode) DeepCopy() *IterationMode {
	if in == nil {
		return nil
	}
	out := new(IterationMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoServerlessWorkflow) DeepCopyInto(out *KogitoServerlessWorkflow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoServerlessWorkflow.
func (in *KogitoServerlessWorkflow) DeepCopy() *KogitoServerlessWorkflow {
	if in == nil {
		return nil
	}
	out := new(KogitoServerlessWorkflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoServerlessWorkflow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoServerlessWorkflowList) DeepCopyInto(out *KogitoServerlessWorkflowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KogitoServerlessWorkflow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoServerlessWorkflowList.
func (in *KogitoServerlessWorkflowList) DeepCopy() *KogitoServerlessWorkflowList {
	if in == nil {
		return nil
	}
	out := new(KogitoServerlessWorkflowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KogitoServerlessWorkflowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoServerlessWorkflowSpec) DeepCopyInto(out *KogitoServerlessWorkflowSpec) {
	*out = *in
	if in.Constants != nil {
		in, out := &in.Constants, &out.Constants
		*out = make([]Constant, len(*in))
		copy(*out, *in)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new([]v1.Secret)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.Secret, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = make([]Timeout, len(*in))
		copy(*out, *in)
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]Error, len(*in))
		copy(*out, *in)
	}
	in.Auth.DeepCopyInto(&out.Auth)
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = new([]Event)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Event, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Functions != nil {
		in, out := &in.Functions, &out.Functions
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Retries = in.Retries
	if in.States != nil {
		in, out := &in.States, &out.States
		*out = make([]State, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoServerlessWorkflowSpec.
func (in *KogitoServerlessWorkflowSpec) DeepCopy() *KogitoServerlessWorkflowSpec {
	if in == nil {
		return nil
	}
	out := new(KogitoServerlessWorkflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KogitoServerlessWorkflowStatus) DeepCopyInto(out *KogitoServerlessWorkflowStatus) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]Endpoint, len(*in))
		copy(*out, *in)
	}
	in.Address.DeepCopyInto(&out.Address)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KogitoServerlessWorkflowStatus.
func (in *KogitoServerlessWorkflowStatus) DeepCopy() *KogitoServerlessWorkflowStatus {
	if in == nil {
		return nil
	}
	out := new(KogitoServerlessWorkflowStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuth2Properties) DeepCopyInto(out *OAuth2Properties) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Audiences != nil {
		in, out := &in.Audiences, &out.Audiences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]Metadata, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuth2Properties.
func (in *OAuth2Properties) DeepCopy() *OAuth2Properties {
	if in == nil {
		return nil
	}
	out := new(OAuth2Properties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProduceEvents) DeepCopyInto(out *ProduceEvents) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProduceEvents.
func (in *ProduceEvents) DeepCopy() *ProduceEvents {
	if in == nil {
		return nil
	}
	out := new(ProduceEvents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Retry) DeepCopyInto(out *Retry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Retry.
func (in *Retry) DeepCopy() *Retry {
	if in == nil {
		return nil
	}
	out := new(Retry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sleep) DeepCopyInto(out *Sleep) {
	*out = *in
	if in.Before != nil {
		in, out := &in.Before, &out.Before
		*out = new(string)
		**out = **in
	}
	if in.After != nil {
		in, out := &in.After, &out.After
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sleep.
func (in *Sleep) DeepCopy() *Sleep {
	if in == nil {
		return nil
	}
	out := new(Sleep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *State) DeepCopyInto(out *State) {
	*out = *in
	if in.Exclusive != nil {
		in, out := &in.Exclusive, &out.Exclusive
		*out = new(bool)
		**out = **in
	}
	if in.ActionMode != nil {
		in, out := &in.ActionMode, &out.ActionMode
		*out = new(ActionMode)
		**out = **in
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = new([]Action)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Action, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.DataConditions != nil {
		in, out := &in.DataConditions, &out.DataConditions
		*out = new([]DataCondition)
		if **in != nil {
			in, out := *in, *out
			*out = make([]DataCondition, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.EventConditions != nil {
		in, out := &in.EventConditions, &out.EventConditions
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.DefaultCondition != nil {
		in, out := &in.DefaultCondition, &out.DefaultCondition
		*out = new(string)
		**out = **in
	}
	if in.OnEvents != nil {
		in, out := &in.OnEvents, &out.OnEvents
		*out = new([]Event)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Event, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.Branches != nil {
		in, out := &in.Branches, &out.Branches
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.CompletionType != nil {
		in, out := &in.CompletionType, &out.CompletionType
		*out = new(CompletionType)
		**out = **in
	}
	if in.NumCompleted != nil {
		in, out := &in.NumCompleted, &out.NumCompleted
		*out = new(int)
		**out = **in
	}
	if in.InputCollection != nil {
		in, out := &in.InputCollection, &out.InputCollection
		*out = new(string)
		**out = **in
	}
	if in.OutputCollection != nil {
		in, out := &in.OutputCollection, &out.OutputCollection
		*out = new(string)
		**out = **in
	}
	if in.IterationParam != nil {
		in, out := &in.IterationParam, &out.IterationParam
		*out = new(string)
		**out = **in
	}
	if in.BatchSize != nil {
		in, out := &in.BatchSize, &out.BatchSize
		*out = new(int)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(IterationMode)
		**out = **in
	}
	if in.EventRef != nil {
		in, out := &in.EventRef, &out.EventRef
		*out = new(EventRef)
		(*in).DeepCopyInto(*out)
	}
	if in.EventDataFilter != nil {
		in, out := &in.EventDataFilter, &out.EventDataFilter
		*out = new(EventDataFilter)
		**out = **in
	}
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(Timeouts)
		**out = **in
	}
	if in.StateDataFilter != nil {
		in, out := &in.StateDataFilter, &out.StateDataFilter
		*out = new(StateDataFilter)
		**out = **in
	}
	if in.Transition != nil {
		in, out := &in.Transition, &out.Transition
		*out = new(string)
		**out = **in
	}
	if in.OnErrors != nil {
		in, out := &in.OnErrors, &out.OnErrors
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.End != nil {
		in, out := &in.End, &out.End
		*out = new(bool)
		**out = **in
	}
	if in.CompensatedBy != nil {
		in, out := &in.CompensatedBy, &out.CompensatedBy
		*out = new(string)
		**out = **in
	}
	if in.UsedForCompensation != nil {
		in, out := &in.UsedForCompensation, &out.UsedForCompensation
		*out = new(bool)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new([]Metadata)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Metadata, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new State.
func (in *State) DeepCopy() *State {
	if in == nil {
		return nil
	}
	out := new(State)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateDataFilter) DeepCopyInto(out *StateDataFilter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateDataFilter.
func (in *StateDataFilter) DeepCopy() *StateDataFilter {
	if in == nil {
		return nil
	}
	out := new(StateDataFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timeout) DeepCopyInto(out *Timeout) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timeout.
func (in *Timeout) DeepCopy() *Timeout {
	if in == nil {
		return nil
	}
	out := new(Timeout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timeouts) DeepCopyInto(out *Timeouts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timeouts.
func (in *Timeouts) DeepCopy() *Timeouts {
	if in == nil {
		return nil
	}
	out := new(Timeouts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transition) DeepCopyInto(out *Transition) {
	*out = *in
	if in.ProduceEvents != nil {
		in, out := &in.ProduceEvents, &out.ProduceEvents
		*out = new([]ProduceEvents)
		if **in != nil {
			in, out := *in, *out
			*out = make([]ProduceEvents, len(*in))
			copy(*out, *in)
		}
	}
	if in.Compensate != nil {
		in, out := &in.Compensate, &out.Compensate
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transition.
func (in *Transition) DeepCopy() *Transition {
	if in == nil {
		return nil
	}
	out := new(Transition)
	in.DeepCopyInto(out)
	return out
}
