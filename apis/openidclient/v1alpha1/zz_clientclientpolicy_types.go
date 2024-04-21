// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type ClientClientPolicyInitParameters struct {

	// The clients allowed by this client policy.
	Clients []*string `json:"clients,omitempty" tf:"clients,omitempty"`

	// (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of AFFIRMATIVE, CONSENSUS, or UNANIMOUS. Applies to permissions.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// The description of this client policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Computed) Dictates how the policy decision should be made. Can be either POSITIVE or NEGATIVE. Applies to policies.
	Logic *string `json:"logic,omitempty" tf:"logic,omitempty"`

	// The name of this client policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the resource server this client policy is attached to.
	ResourceServerID *string `json:"resourceServerId,omitempty" tf:"resource_server_id,omitempty"`
}

type ClientClientPolicyObservation struct {

	// The clients allowed by this client policy.
	Clients []*string `json:"clients,omitempty" tf:"clients,omitempty"`

	// (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of AFFIRMATIVE, CONSENSUS, or UNANIMOUS. Applies to permissions.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// The description of this client policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Computed) Dictates how the policy decision should be made. Can be either POSITIVE or NEGATIVE. Applies to policies.
	Logic *string `json:"logic,omitempty" tf:"logic,omitempty"`

	// The name of this client policy.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm this client policy exists within.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// The ID of the resource server this client policy is attached to.
	ResourceServerID *string `json:"resourceServerId,omitempty" tf:"resource_server_id,omitempty"`
}

type ClientClientPolicyParameters struct {

	// The clients allowed by this client policy.
	// +kubebuilder:validation:Optional
	Clients []*string `json:"clients,omitempty" tf:"clients,omitempty"`

	// (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of AFFIRMATIVE, CONSENSUS, or UNANIMOUS. Applies to permissions.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// The description of this client policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Computed) Dictates how the policy decision should be made. Can be either POSITIVE or NEGATIVE. Applies to policies.
	// +kubebuilder:validation:Optional
	Logic *string `json:"logic,omitempty" tf:"logic,omitempty"`

	// The name of this client policy.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The realm this client policy exists within.
	// +crossplane:generate:reference:type=github.com/viletay/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// The ID of the resource server this client policy is attached to.
	// +kubebuilder:validation:Optional
	ResourceServerID *string `json:"resourceServerId,omitempty" tf:"resource_server_id,omitempty"`
}

// ClientClientPolicySpec defines the desired state of ClientClientPolicy
type ClientClientPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClientClientPolicyParameters `json:"forProvider"`
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
	InitProvider ClientClientPolicyInitParameters `json:"initProvider,omitempty"`
}

// ClientClientPolicyStatus defines the observed state of ClientClientPolicy.
type ClientClientPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClientClientPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClientClientPolicy is the Schema for the ClientClientPolicys API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type ClientClientPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clients) || (has(self.initProvider) && has(self.initProvider.clients))",message="spec.forProvider.clients is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceServerId) || (has(self.initProvider) && has(self.initProvider.resourceServerId))",message="spec.forProvider.resourceServerId is a required parameter"
	Spec   ClientClientPolicySpec   `json:"spec"`
	Status ClientClientPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClientClientPolicyList contains a list of ClientClientPolicys
type ClientClientPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClientClientPolicy `json:"items"`
}

// Repository type metadata.
var (
	ClientClientPolicy_Kind             = "ClientClientPolicy"
	ClientClientPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClientClientPolicy_Kind}.String()
	ClientClientPolicy_KindAPIVersion   = ClientClientPolicy_Kind + "." + CRDGroupVersion.String()
	ClientClientPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ClientClientPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ClientClientPolicy{}, &ClientClientPolicyList{})
}
