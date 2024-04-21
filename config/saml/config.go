package saml

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "saml"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_saml_identity_provider", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["realm"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
	})

}
