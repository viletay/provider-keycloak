/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ExecutionConfigInitParameters struct {

	// The name of the configuration.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// The authentication execution this configuration is attached to.
	ExecutionID *string `json:"executionId,omitempty" tf:"execution_id,omitempty"`

	// The realm the authentication execution exists in.
	// +crossplane:generate:reference:type=github.com/viletay/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

type ExecutionConfigObservation struct {

	// The name of the configuration.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// The authentication execution this configuration is attached to.
	ExecutionID *string `json:"executionId,omitempty" tf:"execution_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The realm the authentication execution exists in.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`
}

type ExecutionConfigParameters struct {

	// The name of the configuration.
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// The authentication execution this configuration is attached to.
	// +kubebuilder:validation:Optional
	ExecutionID *string `json:"executionId,omitempty" tf:"execution_id,omitempty"`

	// The realm the authentication execution exists in.
	// +crossplane:generate:reference:type=github.com/viletay/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

// ExecutionConfigSpec defines the desired state of ExecutionConfig
type ExecutionConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExecutionConfigParameters `json:"forProvider"`
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
	InitProvider ExecutionConfigInitParameters `json:"initProvider,omitempty"`
}

// ExecutionConfigStatus defines the observed state of ExecutionConfig.
type ExecutionConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExecutionConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ExecutionConfig is the Schema for the ExecutionConfigs API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type ExecutionConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.alias) || (has(self.initProvider) && has(self.initProvider.alias))",message="spec.forProvider.alias is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || (has(self.initProvider) && has(self.initProvider.config))",message="spec.forProvider.config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.executionId) || (has(self.initProvider) && has(self.initProvider.executionId))",message="spec.forProvider.executionId is a required parameter"
	Spec   ExecutionConfigSpec   `json:"spec"`
	Status ExecutionConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExecutionConfigList contains a list of ExecutionConfigs
type ExecutionConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExecutionConfig `json:"items"`
}

// Repository type metadata.
var (
	ExecutionConfig_Kind             = "ExecutionConfig"
	ExecutionConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExecutionConfig_Kind}.String()
	ExecutionConfig_KindAPIVersion   = ExecutionConfig_Kind + "." + CRDGroupVersion.String()
	ExecutionConfig_GroupVersionKind = CRDGroupVersion.WithKind(ExecutionConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&ExecutionConfig{}, &ExecutionConfigList{})
}
