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

// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AccessKey.
func (mg *AccessKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.Username,
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UsernameRef,
		Selector:     mg.Spec.ForProvider.UsernameSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Username")
	}
	mg.Spec.ForProvider.Username = rsp.ResolvedValue
	mg.Spec.ForProvider.UsernameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupPolicyAttachment.
func (mg *GroupPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.PolicyARN,
		Extract:      PolicyARN(),
		Reference:    mg.Spec.ForProvider.PolicyARNRef,
		Selector:     mg.Spec.ForProvider.PolicyARNSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyARN")
	}
	mg.Spec.ForProvider.PolicyARN = rsp.ResolvedValue
	mg.Spec.ForProvider.PolicyARNRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.GroupName,
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupNameRef,
		Selector:     mg.Spec.ForProvider.GroupNameSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupName")
	}
	mg.Spec.ForProvider.GroupName = rsp.ResolvedValue
	mg.Spec.ForProvider.GroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupUserMembership.
func (mg *GroupUserMembership) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.GroupName,
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupNameRef,
		Selector:     mg.Spec.ForProvider.GroupNameSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupName")
	}
	mg.Spec.ForProvider.GroupName = rsp.ResolvedValue
	mg.Spec.ForProvider.GroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.UserName,
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserNameRef,
		Selector:     mg.Spec.ForProvider.UserNameSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserName")
	}
	mg.Spec.ForProvider.UserName = rsp.ResolvedValue
	mg.Spec.ForProvider.UserNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RolePolicyAttachment.
func (mg *RolePolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: mg.Spec.ForProvider.PolicyARNs,
		Extract:       PolicyARN(),
		References:    mg.Spec.ForProvider.PolicyARNRefs,
		Selector:      mg.Spec.ForProvider.PolicyARNSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyARNs")
	}
	mg.Spec.ForProvider.PolicyARNs = mrsp.ResolvedValues
	mg.Spec.ForProvider.PolicyARNRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.RoleName,
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RoleNameRef,
		Selector:     mg.Spec.ForProvider.RoleNameSelector,
		To: reference.To{
			List:    &RoleList{},
			Managed: &Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleName")
	}
	mg.Spec.ForProvider.RoleName = rsp.ResolvedValue
	mg.Spec.ForProvider.RoleNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPolicyAttachment.
func (mg *UserPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.PolicyARN,
		Extract:      PolicyARN(),
		Reference:    mg.Spec.ForProvider.PolicyARNRef,
		Selector:     mg.Spec.ForProvider.PolicyARNSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyARN")
	}
	mg.Spec.ForProvider.PolicyARN = rsp.ResolvedValue
	mg.Spec.ForProvider.PolicyARNRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.UserName,
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserNameRef,
		Selector:     mg.Spec.ForProvider.UserNameSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.UserName")
	}
	mg.Spec.ForProvider.UserName = rsp.ResolvedValue
	mg.Spec.ForProvider.UserNameRef = rsp.ResolvedReference

	return nil
}
