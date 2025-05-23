// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Account.
	apisresolver "github.com/upbound/provider-azure/internal/apis"
)

func (mg *Account) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	if mg.Spec.ForProvider.NetworkRules != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.NetworkRules.VirtualNetworkSubnetIds),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.ForProvider.NetworkRules.VirtualNetworkSubnetIdsRefs,
				Selector:      mg.Spec.ForProvider.NetworkRules.VirtualNetworkSubnetIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkRules.VirtualNetworkSubnetIds")
		}
		mg.Spec.ForProvider.NetworkRules.VirtualNetworkSubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.NetworkRules.VirtualNetworkSubnetIdsRefs = mrsp.ResolvedReferences

	}
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.NetworkRules != nil {
		{
			m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.NetworkRules.VirtualNetworkSubnetIds),
				Extract:       resource.ExtractResourceID(),
				References:    mg.Spec.InitProvider.NetworkRules.VirtualNetworkSubnetIdsRefs,
				Selector:      mg.Spec.InitProvider.NetworkRules.VirtualNetworkSubnetIdsSelector,
				To:            reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.NetworkRules.VirtualNetworkSubnetIds")
		}
		mg.Spec.InitProvider.NetworkRules.VirtualNetworkSubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.InitProvider.NetworkRules.VirtualNetworkSubnetIdsRefs = mrsp.ResolvedReferences

	}

	return nil
}

// ResolveReferences of this AccountLocalUser.
func (mg *AccountLocalUser) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PermissionScope); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PermissionScope[i3].ResourceName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.PermissionScope[i3].ResourceNameRef,
				Selector:     mg.Spec.ForProvider.PermissionScope[i3].ResourceNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PermissionScope[i3].ResourceName")
		}
		mg.Spec.ForProvider.PermissionScope[i3].ResourceName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PermissionScope[i3].ResourceNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta2", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.StorageAccountIDRef,
			Selector:     mg.Spec.ForProvider.StorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountID")
	}
	mg.Spec.ForProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.PermissionScope); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PermissionScope[i3].ResourceName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.PermissionScope[i3].ResourceNameRef,
				Selector:     mg.Spec.InitProvider.PermissionScope[i3].ResourceNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PermissionScope[i3].ResourceName")
		}
		mg.Spec.InitProvider.PermissionScope[i3].ResourceName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PermissionScope[i3].ResourceNameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this BlobInventoryPolicy.
func (mg *BlobInventoryPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Rules); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Rules[i3].StorageContainerName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Rules[i3].StorageContainerNameRef,
				Selector:     mg.Spec.ForProvider.Rules[i3].StorageContainerNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Rules[i3].StorageContainerName")
		}
		mg.Spec.ForProvider.Rules[i3].StorageContainerName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Rules[i3].StorageContainerNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta2", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.StorageAccountIDRef,
			Selector:     mg.Spec.ForProvider.StorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountID")
	}
	mg.Spec.ForProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Rules); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta1", "Container", "ContainerList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Rules[i3].StorageContainerName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.Rules[i3].StorageContainerNameRef,
				Selector:     mg.Spec.InitProvider.Rules[i3].StorageContainerNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Rules[i3].StorageContainerName")
		}
		mg.Spec.InitProvider.Rules[i3].StorageContainerName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Rules[i3].StorageContainerNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta2", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageAccountID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.StorageAccountIDRef,
			Selector:     mg.Spec.InitProvider.StorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StorageAccountID")
	}
	mg.Spec.InitProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StorageAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ManagementPolicy.
func (mg *ManagementPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta2", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageAccountID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.StorageAccountIDRef,
			Selector:     mg.Spec.ForProvider.StorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.StorageAccountID")
	}
	mg.Spec.ForProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.StorageAccountIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("storage.azure.upbound.io", "v1beta2", "Account", "AccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageAccountID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.StorageAccountIDRef,
			Selector:     mg.Spec.InitProvider.StorageAccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.StorageAccountID")
	}
	mg.Spec.InitProvider.StorageAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.StorageAccountIDRef = rsp.ResolvedReference

	return nil
}
