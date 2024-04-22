/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/viletay/provider-keycloak/config/authentication"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/viletay/provider-keycloak/config/common"
	"github.com/viletay/provider-keycloak/config/defaults"
	"github.com/viletay/provider-keycloak/config/group"
	"github.com/viletay/provider-keycloak/config/mapper"
	"github.com/viletay/provider-keycloak/config/oidc"
	"github.com/viletay/provider-keycloak/config/openidclient"
	"github.com/viletay/provider-keycloak/config/openidgroup"
	"github.com/viletay/provider-keycloak/config/realm"
	"github.com/viletay/provider-keycloak/config/role"
	"github.com/viletay/provider-keycloak/config/saml"
	"github.com/viletay/provider-keycloak/config/samlclient"
	"github.com/viletay/provider-keycloak/config/user"
)

const (
	resourcePrefix = "keycloak"
	modulePath     = "github.com/viletay/provider-keycloak"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("keycloak.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			KnownReferencers(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		authentication.Configure,
		realm.Configure,

		user.Configure,
		group.Configure,
		role.Configure,

		oidc.Configure,
		openidclient.Configure,
		openidgroup.Configure,

		saml.Configure,
		samlclient.Configure,

		mapper.Configure,
		defaults.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// KnownReferencers adds referencers for fields that are known and common among
// more than a few resources.
func KnownReferencers() ujconfig.ResourceOption { //nolint:gocyclo
	return func(r *ujconfig.Resource) {
		for k, s := range r.TerraformResource.Schema {
			// We shouldn't add referencers for status fields and sensitive fields
			// since they already have secret referencer.
			if (s.Computed && !s.Optional) || s.Sensitive {
				continue
			}
			switch k {
			case "realm_id":
				r.References["realm_id"] = ujconfig.Reference{
					Type: "github.com/viletay/provider-keycloak/apis/realm/v1alpha1.Realm",
				}
			case "client_id":
				r.References["client_id"] = ujconfig.Reference{
					Type: "github.com/viletay/provider-keycloak/apis/openidclient/v1alpha1.Client",
				}
			case "service_account_user_id":
				r.References["service_account_user_id"] = ujconfig.Reference{
					Type:              "github.com/viletay/provider-keycloak/apis/openidclient/v1alpha1.Client",
					Extractor:         common.PathServiceAccountRoleIDExtractor,
					RefFieldName:      "ServiceAccountUserClientIDRef",
					SelectorFieldName: "ServiceAccountUserClientIDSelector",
				}
				r.LateInitializer = ujconfig.LateInitializer{
					IgnoredFields: []string{"service_account_user_id"},
				}
			}

		}
	}
}
