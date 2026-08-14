# Provider BunnyNet

`provider-bunnynet` is a [Crossplane](https://crossplane.io/) provider built using [Upjet](https://github.com/crossplane/upjet) that exposes XRM-conformant managed resources for the [Bunny.net API](https://docs.bunny.net/reference/api-overview).

## Getting Started

This provider uses the **family** distribution model: install only the sub-providers you need. The first sub-provider you install automatically pulls in `provider-family-bunnynet`, which manages the shared `ProviderConfig`.

> For provider families background, see [Scalable Provider Families](https://blog.crossplane.io/crd-scaling-provider-families/) and the [Upbound Provider Families docs](https://docs.upbound.io/manuals/packages/providers/provider-families/).

### Install

Install only the Bunny.net services you need:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-bunnynet-cdn
spec:
  package: xpkg.upbound.io/wildbitca/provider-bunnynet-cdn:v0.2.0
```

Need more services? Add more sub-providers:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-bunnynet-dns
spec:
  package: xpkg.upbound.io/wildbitca/provider-bunnynet-dns:v0.2.0
---
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-bunnynet-storage
spec:
  package: xpkg.upbound.io/wildbitca/provider-bunnynet-storage:v0.2.0
```

### Available sub-providers

| Sub-provider | Resources |
|-------------|-----------|
| `provider-bunnynet-cdn` | Pull zones, hostnames, edge rules, optimizer classes, access lists |
| `provider-bunnynet-dns` | DNS zones and records |
| `provider-bunnynet-storage` | Storage zones |
| `provider-bunnynet-stream` | Stream libraries, collections, videos |
| `provider-bunnynet-compute` | Compute scripts, secrets, variables, container apps, image registries, databases |

### ProviderConfig

Create a Secret with Bunny.net credentials:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: bunnynet-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "access_key": "YOUR_BUNNYNET_API_KEY"
    }
```

Create a ProviderConfig:

```yaml
apiVersion: upjet-bunnynet.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: bunnynet-creds
      namespace: crossplane-system
      key: credentials
```

### Rate limiting

For large deployments, use a `DeploymentRuntimeConfig` to throttle:

```yaml
apiVersion: pkg.crossplane.io/v1beta1
kind: DeploymentRuntimeConfig
metadata:
  name: bunnynet-throttled
spec:
  deploymentTemplate:
    spec:
      selector: {}
      template:
        spec:
          containers:
            - name: package-runtime
              args:
                - --poll=30m
                - --max-reconcile-rate=1
                - --sync=4h
```

## Developing

```bash
make submodules
go install golang.org/x/tools/cmd/goimports@latest
export PATH="$(go env GOPATH)/bin:$PATH"
make generate
```

Run against a Kubernetes cluster:

```bash
make run
```

Build family sub-providers:

```bash
make build.family
make build.family FAMILY_SUBPACKAGES="config cdn dns"
```

## Supported resources

This provider exposes 23 managed resources from the [Bunny.net Terraform Provider v0.18.0](https://registry.terraform.io/providers/BunnyWay/bunnynet/0.18.0/docs), organized in 5 API groups:

- **CDN**: Pull zones, hostnames, edge rules, optimizer classes, access lists
- **DNS**: DNS zones, DNS records
- **Storage**: Storage zones
- **Stream**: Libraries, collections, videos
- **Compute**: Scripts, secrets, variables, container apps, image registries, databases

## Report a Bug

Open an [issue](https://github.com/wildbitca/provider-upjet-bunnynet/issues).
