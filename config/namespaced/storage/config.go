package storage

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds the storage group resource configurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("bunnynet_storage_zone", func(r *config.Resource) {
		r.ShortGroup = "storage"
	})
	p.AddResourceConfigurator("bunnynet_storage_file", func(r *config.Resource) {
		r.ShortGroup = "storage"
		r.References["storage_zone_id"] = config.Reference{
			TerraformName: "bunnynet_storage_zone",
		}
	})
}
