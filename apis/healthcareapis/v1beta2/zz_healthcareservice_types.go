// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthenticationConfigurationInitParameters struct {

	// The intended audience to receive authentication tokens for the service. The default value is https://azurehealthcareapis.com
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service.
	// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
	Authority *string `json:"authority,omitempty" tf:"authority,omitempty"`

	// (Boolean) Enables the 'SMART on FHIR' option for mobile and web implementations.
	SmartProxyEnabled *bool `json:"smartProxyEnabled,omitempty" tf:"smart_proxy_enabled,omitempty"`
}

type AuthenticationConfigurationObservation struct {

	// The intended audience to receive authentication tokens for the service. The default value is https://azurehealthcareapis.com
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service.
	// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
	Authority *string `json:"authority,omitempty" tf:"authority,omitempty"`

	// (Boolean) Enables the 'SMART on FHIR' option for mobile and web implementations.
	SmartProxyEnabled *bool `json:"smartProxyEnabled,omitempty" tf:"smart_proxy_enabled,omitempty"`
}

type AuthenticationConfigurationParameters struct {

	// The intended audience to receive authentication tokens for the service. The default value is https://azurehealthcareapis.com
	// +kubebuilder:validation:Optional
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service.
	// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
	// +kubebuilder:validation:Optional
	Authority *string `json:"authority,omitempty" tf:"authority,omitempty"`

	// (Boolean) Enables the 'SMART on FHIR' option for mobile and web implementations.
	// +kubebuilder:validation:Optional
	SmartProxyEnabled *bool `json:"smartProxyEnabled,omitempty" tf:"smart_proxy_enabled,omitempty"`
}

type CorsConfigurationInitParameters struct {

	// (Boolean) If credentials are allowed via CORS.
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// A set of headers to be allowed via CORS.
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// The methods to be allowed via CORS. Possible values are DELETE, GET, HEAD, MERGE, POST, OPTIONS, PATCH and PUT.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// A set of origins to be allowed via CORS.
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// The max age to be allowed via CORS.
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`
}

type CorsConfigurationObservation struct {

	// (Boolean) If credentials are allowed via CORS.
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// A set of headers to be allowed via CORS.
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// The methods to be allowed via CORS. Possible values are DELETE, GET, HEAD, MERGE, POST, OPTIONS, PATCH and PUT.
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// A set of origins to be allowed via CORS.
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// The max age to be allowed via CORS.
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`
}

type CorsConfigurationParameters struct {

	// (Boolean) If credentials are allowed via CORS.
	// +kubebuilder:validation:Optional
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// A set of headers to be allowed via CORS.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// The methods to be allowed via CORS. Possible values are DELETE, GET, HEAD, MERGE, POST, OPTIONS, PATCH and PUT.
	// +kubebuilder:validation:Optional
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// A set of origins to be allowed via CORS.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// The max age to be allowed via CORS.
	// +kubebuilder:validation:Optional
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`
}

type HealthcareServiceIdentityInitParameters struct {

	// The type of managed identity to assign. The only possible value is SystemAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthcareServiceIdentityObservation struct {

	// The ID of the Healthcare Service.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The ID of the Healthcare Service.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The type of managed identity to assign. The only possible value is SystemAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthcareServiceIdentityParameters struct {

	// The type of managed identity to assign. The only possible value is SystemAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type HealthcareServiceInitParameters struct {

	// A set of Azure object IDs that are allowed to access the Service.
	// +listType=set
	AccessPolicyObjectIds []*string `json:"accessPolicyObjectIds,omitempty" tf:"access_policy_object_ids,omitempty"`

	// An authentication_configuration block as defined below.
	AuthenticationConfiguration *AuthenticationConfigurationInitParameters `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration,omitempty"`

	// Specifies the name of the storage account which the operation configuration information is exported to.
	ConfigurationExportStorageAccountName *string `json:"configurationExportStorageAccountName,omitempty" tf:"configuration_export_storage_account_name,omitempty"`

	// A cors_configuration block as defined below.
	CorsConfiguration *CorsConfigurationInitParameters `json:"corsConfiguration,omitempty" tf:"cors_configuration,omitempty"`

	// A versionless Key Vault Key ID for CMK encryption of the backing database. Changing this forces a new resource to be created.
	CosmosDBKeyVaultKeyVersionlessID *string `json:"cosmosdbKeyVaultKeyVersionlessId,omitempty" tf:"cosmosdb_key_vault_key_versionless_id,omitempty"`

	// The provisioned throughput for the backing database. Range of 400-100000. Defaults to 1000.
	CosmosDBThroughput *float64 `json:"cosmosdbThroughput,omitempty" tf:"cosmosdb_throughput,omitempty"`

	// An identity block as defined below.
	Identity *HealthcareServiceIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The type of the service. Values at time of publication are: fhir, fhir-Stu3 and fhir-R4. Default value is fhir.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the supported Azure Region where the Service should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether public network access is enabled or disabled for this service instance. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HealthcareServiceObservation struct {

	// A set of Azure object IDs that are allowed to access the Service.
	// +listType=set
	AccessPolicyObjectIds []*string `json:"accessPolicyObjectIds,omitempty" tf:"access_policy_object_ids,omitempty"`

	// An authentication_configuration block as defined below.
	AuthenticationConfiguration *AuthenticationConfigurationObservation `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration,omitempty"`

	// Specifies the name of the storage account which the operation configuration information is exported to.
	ConfigurationExportStorageAccountName *string `json:"configurationExportStorageAccountName,omitempty" tf:"configuration_export_storage_account_name,omitempty"`

	// A cors_configuration block as defined below.
	CorsConfiguration *CorsConfigurationObservation `json:"corsConfiguration,omitempty" tf:"cors_configuration,omitempty"`

	// A versionless Key Vault Key ID for CMK encryption of the backing database. Changing this forces a new resource to be created.
	CosmosDBKeyVaultKeyVersionlessID *string `json:"cosmosdbKeyVaultKeyVersionlessId,omitempty" tf:"cosmosdb_key_vault_key_versionless_id,omitempty"`

	// The provisioned throughput for the backing database. Range of 400-100000. Defaults to 1000.
	CosmosDBThroughput *float64 `json:"cosmosdbThroughput,omitempty" tf:"cosmosdb_throughput,omitempty"`

	// The ID of the Healthcare Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity *HealthcareServiceIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The type of the service. Values at time of publication are: fhir, fhir-Stu3 and fhir-R4. Default value is fhir.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the supported Azure Region where the Service should be created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether public network access is enabled or disabled for this service instance. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the Resource Group in which to create the Service. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HealthcareServiceParameters struct {

	// A set of Azure object IDs that are allowed to access the Service.
	// +kubebuilder:validation:Optional
	// +listType=set
	AccessPolicyObjectIds []*string `json:"accessPolicyObjectIds,omitempty" tf:"access_policy_object_ids,omitempty"`

	// An authentication_configuration block as defined below.
	// +kubebuilder:validation:Optional
	AuthenticationConfiguration *AuthenticationConfigurationParameters `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration,omitempty"`

	// Specifies the name of the storage account which the operation configuration information is exported to.
	// +kubebuilder:validation:Optional
	ConfigurationExportStorageAccountName *string `json:"configurationExportStorageAccountName,omitempty" tf:"configuration_export_storage_account_name,omitempty"`

	// A cors_configuration block as defined below.
	// +kubebuilder:validation:Optional
	CorsConfiguration *CorsConfigurationParameters `json:"corsConfiguration,omitempty" tf:"cors_configuration,omitempty"`

	// A versionless Key Vault Key ID for CMK encryption of the backing database. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CosmosDBKeyVaultKeyVersionlessID *string `json:"cosmosdbKeyVaultKeyVersionlessId,omitempty" tf:"cosmosdb_key_vault_key_versionless_id,omitempty"`

	// The provisioned throughput for the backing database. Range of 400-100000. Defaults to 1000.
	// +kubebuilder:validation:Optional
	CosmosDBThroughput *float64 `json:"cosmosdbThroughput,omitempty" tf:"cosmosdb_throughput,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity *HealthcareServiceIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// The type of the service. Values at time of publication are: fhir, fhir-Stu3 and fhir-R4. Default value is fhir.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Specifies the supported Azure Region where the Service should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Whether public network access is enabled or disabled for this service instance. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the Resource Group in which to create the Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// HealthcareServiceSpec defines the desired state of HealthcareService
type HealthcareServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthcareServiceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider HealthcareServiceInitParameters `json:"initProvider,omitempty"`
}

// HealthcareServiceStatus defines the observed state of HealthcareService.
type HealthcareServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthcareServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// HealthcareService is the Schema for the HealthcareServices API. Manages a Healthcare Service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HealthcareService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   HealthcareServiceSpec   `json:"spec"`
	Status HealthcareServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcareServiceList contains a list of HealthcareServices
type HealthcareServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthcareService `json:"items"`
}

// Repository type metadata.
var (
	HealthcareService_Kind             = "HealthcareService"
	HealthcareService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HealthcareService_Kind}.String()
	HealthcareService_KindAPIVersion   = HealthcareService_Kind + "." + CRDGroupVersion.String()
	HealthcareService_GroupVersionKind = CRDGroupVersion.WithKind(HealthcareService_Kind)
)

func init() {
	SchemeBuilder.Register(&HealthcareService{}, &HealthcareServiceList{})
}
