package authentication

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "authentication"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_authentication_bindings", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("keycloak_authentication_execution", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("keycloak_authentication_execution_config", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("keycloak_authentication_flow", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})
	p.AddResourceConfigurator("keycloak_authentication_subflow", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})
}
