package compute

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "compute"

// Configure adds the compute group resource configurators.
func Configure(p *config.Provider) {
	for _, name := range []string{
		"bunnynet_compute_script",
		"bunnynet_compute_script_secret",
		"bunnynet_compute_script_variable",
		"bunnynet_compute_container_app",
		"bunnynet_compute_container_imageregistry",
		"bunnynet_database",
	} {
		resName := name
		p.AddResourceConfigurator(resName, func(r *config.Resource) {
			r.ShortGroup = shortGroup
		})
	}
}
