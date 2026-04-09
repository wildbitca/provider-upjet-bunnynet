package stream

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds the stream group resource configurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("bunnynet_stream_library", func(r *config.Resource) {
		r.ShortGroup = "stream"
	})
	// Stream library-scoped resources with library_id reference to bunnynet_stream_library.
	for _, name := range []string{
		"bunnynet_stream_collection", "bunnynet_stream_video",
	} {
		resName := name
		p.AddResourceConfigurator(resName, func(r *config.Resource) {
			r.ShortGroup = "stream"
			r.References["library_id"] = config.Reference{
				TerraformName: "bunnynet_stream_library",
			}
		})
	}
}
