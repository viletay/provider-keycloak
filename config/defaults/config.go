package defaults

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "default"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_default_roles", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["default_roles"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-keycloak/apis/role/v1alpha1.Role",
		}

	})
}
