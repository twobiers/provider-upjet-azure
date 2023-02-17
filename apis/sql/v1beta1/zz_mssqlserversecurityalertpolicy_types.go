/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MSSQLServerSecurityAlertPolicyObservation struct {

	// The ID of the MS SQL Server Security Alert Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MSSQLServerSecurityAlertPolicyParameters struct {

	// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action.
	// +kubebuilder:validation:Optional
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to false.
	// +kubebuilder:validation:Optional
	EmailAccountAdmins *bool `json:"emailAccountAdmins,omitempty" tf:"email_account_admins,omitempty"`

	// Specifies an array of email addresses to which the alert is sent.
	// +kubebuilder:validation:Optional
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// The name of the resource group that contains the MS SQL Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to 0.
	// +kubebuilder:validation:Optional
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// Specifies the name of the MS SQL Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLServer
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Reference to a MSSQLServer to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database server. Possible values are Disabled, Enabled and New.
	// +kubebuilder:validation:Required
	State *string `json:"state" tf:"state,omitempty"`

	// Specifies the identifier key of the Threat Detection audit storage account. This is mandatory when you use storage_endpoint to specify a storage account blob endpoint.
	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// Specifies the blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("primary_blob_endpoint",true)
	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`

	// Reference to a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointRef *v1.Reference `json:"storageEndpointRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointSelector *v1.Selector `json:"storageEndpointSelector,omitempty" tf:"-"`
}

// MSSQLServerSecurityAlertPolicySpec defines the desired state of MSSQLServerSecurityAlertPolicy
type MSSQLServerSecurityAlertPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLServerSecurityAlertPolicyParameters `json:"forProvider"`
}

// MSSQLServerSecurityAlertPolicyStatus defines the observed state of MSSQLServerSecurityAlertPolicy.
type MSSQLServerSecurityAlertPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLServerSecurityAlertPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerSecurityAlertPolicy is the Schema for the MSSQLServerSecurityAlertPolicys API. Manages a Security Alert Policy for a MS SQL Server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLServerSecurityAlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MSSQLServerSecurityAlertPolicySpec   `json:"spec"`
	Status            MSSQLServerSecurityAlertPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLServerSecurityAlertPolicyList contains a list of MSSQLServerSecurityAlertPolicys
type MSSQLServerSecurityAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLServerSecurityAlertPolicy `json:"items"`
}

// Repository type metadata.
var (
	MSSQLServerSecurityAlertPolicy_Kind             = "MSSQLServerSecurityAlertPolicy"
	MSSQLServerSecurityAlertPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLServerSecurityAlertPolicy_Kind}.String()
	MSSQLServerSecurityAlertPolicy_KindAPIVersion   = MSSQLServerSecurityAlertPolicy_Kind + "." + CRDGroupVersion.String()
	MSSQLServerSecurityAlertPolicy_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLServerSecurityAlertPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLServerSecurityAlertPolicy{}, &MSSQLServerSecurityAlertPolicyList{})
}