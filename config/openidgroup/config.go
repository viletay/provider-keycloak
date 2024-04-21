package openidgroup

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "openidgroup"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_openid_group_membership_protocol_mapper", func(r *ujconfig.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = Group
		r.References["client_scope_id"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-keycloak/apis/openidclient/v1alpha1.ClientScope",
		}
	})
}
