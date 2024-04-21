package user

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "user"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_user", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_user_groups", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["user_id"] = ujconfig.Reference{
			Type: "User",
		}
		r.References["group_ids"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-keycloak/apis/group/v1alpha1.Group",
		}
	})

	p.AddResourceConfigurator("keycloak_users_permissions", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})
}
