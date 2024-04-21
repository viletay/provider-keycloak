package role

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "role"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_role", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["composite_roles"] = ujconfig.Reference{
			Type: "Role",
		}
	})
}
