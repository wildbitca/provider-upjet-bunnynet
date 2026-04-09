package dns

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds the dns group resource configurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("bunnynet_dns_zone", func(r *config.Resource) {
		r.ShortGroup = "dns"
	})
	// DNS zone-scoped resources with zone_id reference to bunnynet_dns_zone.
	for _, name := range []string{
		"bunnynet_dns_record", "bunnynet_dns_script", "bunnynet_dns_script_variable",
	} {
		resName := name
		p.AddResourceConfigurator(resName, func(r *config.Resource) {
			r.ShortGroup = "dns"
			r.References["zone_id"] = config.Reference{
				TerraformName: "bunnynet_dns_zone",
			}
		})
	}
}
