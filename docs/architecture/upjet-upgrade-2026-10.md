# Upgrade: upjet v2.5.1, crossplane-runtime v2.4.0, bunnynet 0.19.2, OpenTofu 1.13.2

| Component | Before | After |
|-----------|--------|-------|
| upjet | fork v2.2.5 | upstream v2.5.1 via fork tag v2.5.1-wb.1 |
| crossplane-runtime | v2.1.0 | v2.4.0 |
| crossplane/apis/v2 | - | v2.4.1 |
| controller-runtime | v0.22.4 | v0.25.2 |
| controller-tools | v0.19.0 | v0.20.0 |
| crossplane-tools (angryjet) | 2025-10-17 | 2026-07-19 |
| k8s.io/* | v0.34.3 | v0.36.2 |
| Go | 1.26.0 | 1.26.8 |
| BunnyWay/bunnynet | 0.19.1 | 0.19.2 |
| OpenTofu | 1.11.5 | 1.13.2 |

## Notes

- `crossplane-tools` had to be bumped: the old angryjet emitted
  `crossplane-runtime/v2/apis/common/...` imports that no longer exist.
  Its k8s 0.36 requirement forced controller-runtime v0.25.x
  (v0.23 does not compile against client-go 0.36).
- Terraform provider 0.19.2 changes: Optimizer HTML to Markdown on `pullzone`,
  `.` accepted for `SRV` records. The resource set is unchanged (24 resources,
  9 data sources, none wired).
- OpenTofu 1.13.2 initialised the provider schema and generated all CRDs;
  the registry source stays `registry.opentofu.org/BunnyWay/bunnynet`.
- Single-nested blocks (`nesting_mode: single`) are now emitted as CRD
  objects: `Pullzone.routing` and `Shield.botDetection` were arrays before.
- `examples-generated` is derived from provider docs and still shows
  single blocks as one-element lists; use object form in real manifests.
