package group

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "group"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_group", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["parent_id"] = ujconfig.Reference{
			Type: "Group",
		}
	})
	p.AddResourceConfigurator("keycloak_group_memberships", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["group_id"] = ujconfig.Reference{
			Type: "Group",
		}

	})
	p.AddResourceConfigurator("keycloak_group_roles", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["group_id"] = ujconfig.Reference{
			Type: "Group",
		}
		r.References["role_ids"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-keycloak/apis/role/v1alpha1.Role",
		}
	})
	p.AddResourceConfigurator("keycloak_group_permissions", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["group_id"] = ujconfig.Reference{
			Type: "Group",
		}
	})
}
