package cdn

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds the cdn group resource configurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("bunnynet_pullzone", func(r *config.Resource) {
		r.ShortGroup = "cdn"
	})
	// Pull zone-scoped resources with pull_zone_id reference to bunnynet_pullzone.
	for _, name := range []string{
		"bunnynet_pullzone_access_list", "bunnynet_pullzone_edgerule",
		"bunnynet_pullzone_hostname", "bunnynet_pullzone_optimizer_class",
		"bunnynet_pullzone_ratelimit_rule", "bunnynet_pullzone_shield",
		"bunnynet_pullzone_waf_rule",
	} {
		resName := name
		p.AddResourceConfigurator(resName, func(r *config.Resource) {
			r.ShortGroup = "cdn"
			r.References["pull_zone_id"] = config.Reference{
				TerraformName: "bunnynet_pullzone",
			}
		})
	}
}
