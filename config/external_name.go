package config

import "github.com/crossplane/upjet/v2/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider. BunnyWay/bunnynet uses Terraform Plugin Framework. All 24 managed
// resources use IdentifierFromProvider (provider-generated IDs).
// See https://registry.terraform.io/providers/BunnyWay/bunnynet/latest/docs
var ExternalNameConfigs = map[string]config.ExternalName{
	// Account
	"bunnynet_account_subuser": config.IdentifierFromProvider,

	// CDN - Pull Zones
	"bunnynet_pullzone":                 config.IdentifierFromProvider,
	"bunnynet_pullzone_access_list":     config.IdentifierFromProvider,
	"bunnynet_pullzone_edgerule":        config.IdentifierFromProvider,
	"bunnynet_pullzone_hostname":        config.IdentifierFromProvider,
	"bunnynet_pullzone_optimizer_class": config.IdentifierFromProvider,
	"bunnynet_pullzone_ratelimit_rule":  config.IdentifierFromProvider,
	"bunnynet_pullzone_shield":          config.IdentifierFromProvider,
	"bunnynet_pullzone_waf_rule":        config.IdentifierFromProvider,

	// DNS
	"bunnynet_dns_zone":            config.IdentifierFromProvider,
	"bunnynet_dns_record":          config.IdentifierFromProvider,
	"bunnynet_dns_script":          config.IdentifierFromProvider,
	"bunnynet_dns_script_variable": config.IdentifierFromProvider,

	// Storage
	"bunnynet_storage_zone": config.IdentifierFromProvider,
	"bunnynet_storage_file": config.IdentifierFromProvider,

	// Stream (Video)
	"bunnynet_stream_library":    config.IdentifierFromProvider,
	"bunnynet_stream_collection": config.IdentifierFromProvider,
	"bunnynet_stream_video":      config.IdentifierFromProvider,

	// Compute / Edge
	"bunnynet_compute_script":                  config.IdentifierFromProvider,
	"bunnynet_compute_script_secret":           config.IdentifierFromProvider,
	"bunnynet_compute_script_variable":         config.IdentifierFromProvider,
	"bunnynet_compute_container_app":           config.IdentifierFromProvider,
	"bunnynet_compute_container_imageregistry": config.IdentifierFromProvider,

	// Database
	"bunnynet_database": config.IdentifierFromProvider,
}

// ExternalNameConfigurations returns the list of all ExternalName configured resources.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns a list of all resources with external name configured.
func ExternalNameConfigured() []string {
	l := make([]string, 0, len(ExternalNameConfigs))
	for name := range ExternalNameConfigs {
		l = append(l, name)
	}
	return l
}
