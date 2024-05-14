package mapper

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "mapper"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_generic_protocol_mapper", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["client_scope_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-keycloak/apis/openidclient/v1alpha1.ClientScope",
		}
		r.References["client_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-keycloak/apis/openidclient/v1alpha1.Client",
		}
	})
	p.AddResourceConfigurator("keycloak_generic_role_mapper", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["role_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-keycloak/apis/role/v1alpha1.Role",
		}
	})
	p.AddResourceConfigurator("keycloak_custom_identity_provider_mapper", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["realm"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
	})
}
