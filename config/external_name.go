/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"keycloak_realm":                                    config.IdentifierFromProvider,
	"keycloak_required_action":                          config.IdentifierFromProvider,
	"keycloak_realm_keystore_rsa":                       config.IdentifierFromProvider,
	"keycloak_oidc_identity_provider":                   config.IdentifierFromProvider,
	"keycloak_openid_client_client_policy":              config.IdentifierFromProvider,
	"keycloak_openid_client_group_policy":               config.IdentifierFromProvider,
	"keycloak_openid_client_permissions":                config.IdentifierFromProvider,
	"keycloak_openid_client_role_policy":                config.IdentifierFromProvider,
	"keycloak_openid_client_user_policy":                config.IdentifierFromProvider,
	"keycloak_openid_client_default_scopes":             config.IdentifierFromProvider,
	"keycloak_openid_client_scope":                      config.IdentifierFromProvider,
	"keycloak_openid_client":                            config.IdentifierFromProvider,
	"keycloak_openid_client_service_account_realm_role": config.IdentifierFromProvider,
	"keycloak_openid_client_service_account_role":       config.IdentifierFromProvider,
	"keycloak_openid_group_membership_protocol_mapper":  config.IdentifierFromProvider,
	"keycloak_custom_identity_provider_mapper":          config.IdentifierFromProvider,
	"keycloak_generic_protocol_mapper":                  config.IdentifierFromProvider,
	"keycloak_generic_role_mapper":                      config.IdentifierFromProvider,
	"keycloak_user":                                     config.IdentifierFromProvider,
	"keycloak_user_groups":                              config.IdentifierFromProvider,
	"keycloak_users_permissions":                        config.IdentifierFromProvider,
	"keycloak_group":                                    config.IdentifierFromProvider,
	"keycloak_group_roles":                              config.IdentifierFromProvider,
	"keycloak_group_permissions":                        config.IdentifierFromProvider,
	"keycloak_group_memberships":                        config.IdentifierFromProvider,
	"keycloak_role":                                     config.IdentifierFromProvider,
	"keycloak_saml_client":                              config.IdentifierFromProvider,
	"keycloak_saml_client_default_scopes":               config.IdentifierFromProvider,
	"keycloak_saml_client_scope":                        config.IdentifierFromProvider,
	"keycloak_saml_identity_provider":                   config.IdentifierFromProvider,
	"keycloak_default_roles":                            config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"),
	"keycloak_authentication_bindings":                  config.IdentifierFromProvider,
	"keycloak_authentication_execution":                 config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .parameters.parent_flow_alias }}/{{ .external_name }}"),
	"keycloak_authentication_execution_config":          config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .parameters.execution_id }}/{{ .external_name }}"),
	"keycloak_authentication_flow":                      config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"),
	"keycloak_authentication_subflow":                   config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .parameters.parent_flow_alias }}/{{ .external_name }}"),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
