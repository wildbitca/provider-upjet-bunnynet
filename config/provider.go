package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/wildbitca/provider-upjet-bunnynet/hack"

	cdnCluster "github.com/wildbitca/provider-upjet-bunnynet/config/cluster/cdn"
	computeCluster "github.com/wildbitca/provider-upjet-bunnynet/config/cluster/compute"
	dnsCluster "github.com/wildbitca/provider-upjet-bunnynet/config/cluster/dns"
	storageCluster "github.com/wildbitca/provider-upjet-bunnynet/config/cluster/storage"
	streamCluster "github.com/wildbitca/provider-upjet-bunnynet/config/cluster/stream"
	cdnNamespaced "github.com/wildbitca/provider-upjet-bunnynet/config/namespaced/cdn"
	computeNamespaced "github.com/wildbitca/provider-upjet-bunnynet/config/namespaced/compute"
	dnsNamespaced "github.com/wildbitca/provider-upjet-bunnynet/config/namespaced/dns"
	storageNamespaced "github.com/wildbitca/provider-upjet-bunnynet/config/namespaced/storage"
	streamNamespaced "github.com/wildbitca/provider-upjet-bunnynet/config/namespaced/stream"
)

const (
	resourcePrefix = "upjet-bunnynet"
	modulePath     = "github.com/wildbitca/provider-upjet-bunnynet"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("upjet-bunnynet.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithMainTemplate(hack.MainTemplate),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		cdnCluster.Configure,
		dnsCluster.Configure,
		storageCluster.Configure,
		streamCluster.Configure,
		computeCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("upjet-bunnynet.m.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithMainTemplate(hack.MainTemplate),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		cdnNamespaced.Configure,
		dnsNamespaced.Configure,
		storageNamespaced.Configure,
		streamNamespaced.Configure,
		computeNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
