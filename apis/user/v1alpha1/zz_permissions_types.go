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

type ImpersonateScopeInitParameters struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ImpersonateScopeObservation struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ImpersonateScopeParameters struct {

	// Decision strategy of the permission.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageGroupMembershipScopeInitParameters struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageGroupMembershipScopeObservation struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageGroupMembershipScopeParameters struct {

	// Decision strategy of the permission.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageScopeInitParameters struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageScopeObservation struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ManageScopeParameters struct {

	// Decision strategy of the permission.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type MapRolesScopeInitParameters struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type MapRolesScopeObservation struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type MapRolesScopeParameters struct {

	// Decision strategy of the permission.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type PermissionsInitParameters struct {

	// When specified, set the scope based impersonate permission.
	ImpersonateScope []ImpersonateScopeInitParameters `json:"impersonateScope,omitempty" tf:"impersonate_scope,omitempty"`

	// When specified, set the scope based manage_group_membership permission.
	ManageGroupMembershipScope []ManageGroupMembershipScopeInitParameters `json:"manageGroupMembershipScope,omitempty" tf:"manage_group_membership_scope,omitempty"`

	// When specified, set the scope based manage permission.
	ManageScope []ManageScopeInitParameters `json:"manageScope,omitempty" tf:"manage_scope,omitempty"`

	// When specified, set the scope based map_roles permission.
	MapRolesScope []MapRolesScopeInitParameters `json:"mapRolesScope,omitempty" tf:"map_roles_scope,omitempty"`

	// The realm in which to manage fine-grained user permissions.
	// +crossplane:generate:reference:type=github.com/viletay/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// When specified, set the scope based user_impersonated permission.
	UserImpersonatedScope []UserImpersonatedScopeInitParameters `json:"userImpersonatedScope,omitempty" tf:"user_impersonated_scope,omitempty"`

	// When specified, set the scope based view permission.
	ViewScope []ViewScopeInitParameters `json:"viewScope,omitempty" tf:"view_scope,omitempty"`
}

type PermissionsObservation struct {

	// Resource server id representing the realm management client on which these permissions are managed.
	// Resource server id representing the realm management client on which this permission is managed
	AuthorizationResourceServerID *string `json:"authorizationResourceServerId,omitempty" tf:"authorization_resource_server_id,omitempty"`

	// When true, this indicates that fine-grained user permissions are enabled. This will always be true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When specified, set the scope based impersonate permission.
	ImpersonateScope []ImpersonateScopeObservation `json:"impersonateScope,omitempty" tf:"impersonate_scope,omitempty"`

	// When specified, set the scope based manage_group_membership permission.
	ManageGroupMembershipScope []ManageGroupMembershipScopeObservation `json:"manageGroupMembershipScope,omitempty" tf:"manage_group_membership_scope,omitempty"`

	// When specified, set the scope based manage permission.
	ManageScope []ManageScopeObservation `json:"manageScope,omitempty" tf:"manage_scope,omitempty"`

	// When specified, set the scope based map_roles permission.
	MapRolesScope []MapRolesScopeObservation `json:"mapRolesScope,omitempty" tf:"map_roles_scope,omitempty"`

	// The realm in which to manage fine-grained user permissions.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// When specified, set the scope based user_impersonated permission.
	UserImpersonatedScope []UserImpersonatedScopeObservation `json:"userImpersonatedScope,omitempty" tf:"user_impersonated_scope,omitempty"`

	// When specified, set the scope based view permission.
	ViewScope []ViewScopeObservation `json:"viewScope,omitempty" tf:"view_scope,omitempty"`
}

type PermissionsParameters struct {

	// When specified, set the scope based impersonate permission.
	// +kubebuilder:validation:Optional
	ImpersonateScope []ImpersonateScopeParameters `json:"impersonateScope,omitempty" tf:"impersonate_scope,omitempty"`

	// When specified, set the scope based manage_group_membership permission.
	// +kubebuilder:validation:Optional
	ManageGroupMembershipScope []ManageGroupMembershipScopeParameters `json:"manageGroupMembershipScope,omitempty" tf:"manage_group_membership_scope,omitempty"`

	// When specified, set the scope based manage permission.
	// +kubebuilder:validation:Optional
	ManageScope []ManageScopeParameters `json:"manageScope,omitempty" tf:"manage_scope,omitempty"`

	// When specified, set the scope based map_roles permission.
	// +kubebuilder:validation:Optional
	MapRolesScope []MapRolesScopeParameters `json:"mapRolesScope,omitempty" tf:"map_roles_scope,omitempty"`

	// The realm in which to manage fine-grained user permissions.
	// +crossplane:generate:reference:type=github.com/viletay/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// When specified, set the scope based user_impersonated permission.
	// +kubebuilder:validation:Optional
	UserImpersonatedScope []UserImpersonatedScopeParameters `json:"userImpersonatedScope,omitempty" tf:"user_impersonated_scope,omitempty"`

	// When specified, set the scope based view permission.
	// +kubebuilder:validation:Optional
	ViewScope []ViewScopeParameters `json:"viewScope,omitempty" tf:"view_scope,omitempty"`
}

type UserImpersonatedScopeInitParameters struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type UserImpersonatedScopeObservation struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type UserImpersonatedScopeParameters struct {

	// Decision strategy of the permission.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ViewScopeInitParameters struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ViewScopeObservation struct {

	// Decision strategy of the permission.
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

type ViewScopeParameters struct {

	// Decision strategy of the permission.
	// +kubebuilder:validation:Optional
	DecisionStrategy *string `json:"decisionStrategy,omitempty" tf:"decision_strategy,omitempty"`

	// Description of the permission.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Assigned policies to the permission. Each element within this list should be a policy ID.
	// +kubebuilder:validation:Optional
	// +listType=set
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`
}

// PermissionsSpec defines the desired state of Permissions
type PermissionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PermissionsParameters `json:"forProvider"`
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
	InitProvider PermissionsInitParameters `json:"initProvider,omitempty"`
}

// PermissionsStatus defines the observed state of Permissions.
type PermissionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PermissionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Permissions is the Schema for the Permissionss API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type Permissions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionsSpec   `json:"spec"`
	Status            PermissionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PermissionsList contains a list of Permissionss
type PermissionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Permissions `json:"items"`
}

// Repository type metadata.
var (
	Permissions_Kind             = "Permissions"
	Permissions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Permissions_Kind}.String()
	Permissions_KindAPIVersion   = Permissions_Kind + "." + CRDGroupVersion.String()
	Permissions_GroupVersionKind = CRDGroupVersion.WithKind(Permissions_Kind)
)

func init() {
	SchemeBuilder.Register(&Permissions{}, &PermissionsList{})
}
