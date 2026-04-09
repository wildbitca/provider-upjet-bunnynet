# Provider Bunny.net

`provider-bunnynet` is a [Crossplane](https://crossplane.io/) provider built using [Upjet](https://github.com/crossplane/upjet) that exposes XRM-conformant managed resources for the [Bunny.net API](https://docs.bunny.net/reference/bunnynet-api-overview).

## Available as Provider Family

This provider is distributed as a **provider family**. Install only the Bunny.net services you need instead of all 51 CRDs.

| Package | Description | CRDs |
|---------|-------------|------|
| `provider-family-bunnynet` | ProviderConfig (auto-installed) | 5 |
| `provider-bunnynet-cdn` | Pull Zones, Edge Rules, Hostnames, WAF, Shield, Rate Limits | 16 |
| `provider-bunnynet-dns` | DNS Zones, Records, Scripts | 8 |
| `provider-bunnynet-storage` | Storage Zones, Files | 4 |
| `provider-bunnynet-stream` | Stream Libraries, Collections, Videos | 6 |
| `provider-bunnynet-compute` | Edge Scripts, Container Apps, Database | 12 |

A monolith package (`provider-bunnynet`) with all resources is also available.

## Install

Install a sub-provider. The family provider is pulled automatically as a dependency:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-bunnynet-cdn
spec:
  package: xpkg.upbound.io/wildbitca/provider-bunnynet-cdn:v0.1.1
```

## ProviderConfig

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
      "api_key": "YOUR_BUNNYNET_API_KEY"
    }
---
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

Supported credential key: `api_key`.

## Report a Bug

Open an [issue](https://github.com/wildbitca/provider-upjet-bunnynet/issues).
