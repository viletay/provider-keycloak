package samlclient

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "samlclient"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_saml_client", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_saml_client_default_scopes", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_saml_client_scope", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})
}
