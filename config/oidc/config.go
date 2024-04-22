package oidc

import ujconfig "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "oidc"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_oidc_identity_provider", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
		r.References["realm"] = ujconfig.Reference{
			Type: "github.com/viletay/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
		if s, ok := r.TerraformResource.Schema["client_id"]; ok {
			s.Sensitive = true
		}
	})

}
