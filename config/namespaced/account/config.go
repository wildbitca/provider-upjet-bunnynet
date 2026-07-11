package account

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds the account group resource configurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("bunnynet_account_subuser", func(r *config.Resource) {
		r.ShortGroup = "account"
	})
}
