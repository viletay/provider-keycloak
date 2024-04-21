/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	bindings "github.com/viletay/provider-keycloak/internal/controller/authentication/bindings"
	execution "github.com/viletay/provider-keycloak/internal/controller/authentication/execution"
	executionconfig "github.com/viletay/provider-keycloak/internal/controller/authentication/executionconfig"
	flow "github.com/viletay/provider-keycloak/internal/controller/authentication/flow"
	subflow "github.com/viletay/provider-keycloak/internal/controller/authentication/subflow"
	roles "github.com/viletay/provider-keycloak/internal/controller/default/roles"
	group "github.com/viletay/provider-keycloak/internal/controller/group/group"
	memberships "github.com/viletay/provider-keycloak/internal/controller/group/memberships"
	permissions "github.com/viletay/provider-keycloak/internal/controller/group/permissions"
	rolesgroup "github.com/viletay/provider-keycloak/internal/controller/group/roles"
	identityprovidermapper "github.com/viletay/provider-keycloak/internal/controller/mapper/identityprovidermapper"
	protocolmapper "github.com/viletay/provider-keycloak/internal/controller/mapper/protocolmapper"
	rolemapper "github.com/viletay/provider-keycloak/internal/controller/mapper/rolemapper"
	identityprovider "github.com/viletay/provider-keycloak/internal/controller/oidc/identityprovider"
	client "github.com/viletay/provider-keycloak/internal/controller/openidclient/client"
	clientclientpolicy "github.com/viletay/provider-keycloak/internal/controller/openidclient/clientclientpolicy"
	clientdefaultscopes "github.com/viletay/provider-keycloak/internal/controller/openidclient/clientdefaultscopes"
	clientgrouppolicy "github.com/viletay/provider-keycloak/internal/controller/openidclient/clientgrouppolicy"
	clientpermissions "github.com/viletay/provider-keycloak/internal/controller/openidclient/clientpermissions"
	clientrolepolicy "github.com/viletay/provider-keycloak/internal/controller/openidclient/clientrolepolicy"
	clientscope "github.com/viletay/provider-keycloak/internal/controller/openidclient/clientscope"
	clientserviceaccountrealmrole "github.com/viletay/provider-keycloak/internal/controller/openidclient/clientserviceaccountrealmrole"
	clientserviceaccountrole "github.com/viletay/provider-keycloak/internal/controller/openidclient/clientserviceaccountrole"
	clientuserpolicy "github.com/viletay/provider-keycloak/internal/controller/openidclient/clientuserpolicy"
	groupmembershipprotocolmapper "github.com/viletay/provider-keycloak/internal/controller/openidgroup/groupmembershipprotocolmapper"
	providerconfig "github.com/viletay/provider-keycloak/internal/controller/providerconfig"
	keystorersa "github.com/viletay/provider-keycloak/internal/controller/realm/keystorersa"
	realm "github.com/viletay/provider-keycloak/internal/controller/realm/realm"
	requiredaction "github.com/viletay/provider-keycloak/internal/controller/realm/requiredaction"
	role "github.com/viletay/provider-keycloak/internal/controller/role/role"
	identityprovidersaml "github.com/viletay/provider-keycloak/internal/controller/saml/identityprovider"
	clientsamlclient "github.com/viletay/provider-keycloak/internal/controller/samlclient/client"
	clientdefaultscopessamlclient "github.com/viletay/provider-keycloak/internal/controller/samlclient/clientdefaultscopes"
	clientscopesamlclient "github.com/viletay/provider-keycloak/internal/controller/samlclient/clientscope"
	groups "github.com/viletay/provider-keycloak/internal/controller/user/groups"
	permissionsuser "github.com/viletay/provider-keycloak/internal/controller/user/permissions"
	user "github.com/viletay/provider-keycloak/internal/controller/user/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		bindings.Setup,
		execution.Setup,
		executionconfig.Setup,
		flow.Setup,
		subflow.Setup,
		roles.Setup,
		group.Setup,
		memberships.Setup,
		permissions.Setup,
		rolesgroup.Setup,
		identityprovidermapper.Setup,
		protocolmapper.Setup,
		rolemapper.Setup,
		identityprovider.Setup,
		client.Setup,
		clientclientpolicy.Setup,
		clientdefaultscopes.Setup,
		clientgrouppolicy.Setup,
		clientpermissions.Setup,
		clientrolepolicy.Setup,
		clientscope.Setup,
		clientserviceaccountrealmrole.Setup,
		clientserviceaccountrole.Setup,
		clientuserpolicy.Setup,
		groupmembershipprotocolmapper.Setup,
		providerconfig.Setup,
		keystorersa.Setup,
		realm.Setup,
		requiredaction.Setup,
		role.Setup,
		identityprovidersaml.Setup,
		clientsamlclient.Setup,
		clientdefaultscopessamlclient.Setup,
		clientscopesamlclient.Setup,
		groups.Setup,
		permissionsuser.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
