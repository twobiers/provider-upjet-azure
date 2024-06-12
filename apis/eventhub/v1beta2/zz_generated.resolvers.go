// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this EventHub.
	apisresolver "github.com/upbound/provider-azure/internal/apis"
)

func (mg *EventHub) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta2", "EventHubNamespace", "EventHubNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NamespaceNameRef,
			Selector:     mg.Spec.ForProvider.NamespaceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceName")
	}
	mg.Spec.ForProvider.NamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceNameRef = rsp.ResolvedReference
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

	return nil
}

// ResolveReferences of this EventHubNamespace.
func (mg *EventHubNamespace) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.NetworkRulesets != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NetworkRulesets.VirtualNetworkRule); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetID),
					Extract:      rconfig.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetIDRef,
					Selector:     mg.Spec.ForProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetID")
			}
			mg.Spec.ForProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetIDRef = rsp.ResolvedReference

		}
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

	if mg.Spec.InitProvider.NetworkRulesets != nil {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.NetworkRulesets.VirtualNetworkRule); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("network.azure.upbound.io", "v1beta2", "Subnet", "SubnetList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetID),
					Extract:      rconfig.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetIDRef,
					Selector:     mg.Spec.InitProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetID")
			}
			mg.Spec.InitProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.NetworkRulesets.VirtualNetworkRule[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}

	return nil
}