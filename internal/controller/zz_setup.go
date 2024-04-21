// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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
	providerconfig "github.com/viletay/provider-keycloak/internal/controller/providerconfig"
	realm "github.com/viletay/provider-keycloak/internal/controller/realm/realm"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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
		providerconfig.Setup,
		realm.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
