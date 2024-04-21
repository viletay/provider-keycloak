package openidclient

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	// Group is the short group for this provider.
	Group = "openidclient"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("keycloak_openid_client", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_openid_client_default_scopes", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_openid_client_scope", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_openid_client_service_account_role", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_openid_client_service_account_realm_role", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_openid_client_client_policy", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_openid_client_group_policy", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_openid_client_role_policy", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_openid_client_user_policy", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})

	p.AddResourceConfigurator("keycloak_openid_client_permissions", func(r *ujconfig.Resource) {
		r.ShortGroup = Group
	})
}
