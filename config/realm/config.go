package realm

import ujconfig "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_realm", func(r *ujconfig.Resource) {
		r.ShortGroup = "realm"
	})

	p.AddResourceConfigurator("keycloak_required_action", func(r *ujconfig.Resource) {
		r.ShortGroup = "realm"
		r.Kind = "RequiredAction"
	})

	p.AddResourceConfigurator("keycloak_realm_keystore_rsa", func(r *ujconfig.Resource) {
		r.ShortGroup = "realm"
		if s, ok := r.TerraformResource.Schema["private_key"]; ok {
			s.Sensitive = true
		}
		if s, ok := r.TerraformResource.Schema["certificate"]; ok {
			s.Sensitive = true
		}
	})
}
