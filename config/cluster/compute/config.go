package compute

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds the compute group resource configurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("bunnynet_compute_script", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})
	for _, name := range []string{
		"bunnynet_compute_script_secret", "bunnynet_compute_script_variable",
	} {
		resName := name
		p.AddResourceConfigurator(resName, func(r *config.Resource) {
			r.ShortGroup = "compute"
		})
	}
	p.AddResourceConfigurator("bunnynet_compute_container_app", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})
	p.AddResourceConfigurator("bunnynet_compute_container_imageregistry", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})
	p.AddResourceConfigurator("bunnynet_database", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})
}
