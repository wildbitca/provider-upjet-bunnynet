# Changelog

All notable changes to this provider family are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

Upgrades the toolchain and the wrapped Terraform provider. Proposed next
version: **0.5.0** (see the PR description for the semver reasoning).

### Changed

- Upstream `BunnyWay/bunnynet` Terraform provider `v0.19.1` -> `v0.19.2`
  (adds `pullzone` Optimizer HTML to Markdown and `.` as an `SRV` record
  value; no resources or data sources were added, removed or renamed).
- OpenTofu `1.11.5` -> `1.13.2` (Makefile, publish workflow, image build).
- `crossplane/upjet` `v2.2.5` (fork) -> upstream `v2.5.1` carried by the fork
  tag `v2.5.1-wb.1`; `crossplane-runtime` `v2.1.0` -> `v2.4.0`;
  `crossplane/apis/v2` `v2.4.1` (new direct dependency);
  `controller-runtime` `v0.22.4` -> `v0.25.2`; `controller-tools` `v0.19.0` ->
  `v0.20.0`; `k8s.io/*` `v0.34.3` -> `v0.36.2`; Go `1.26.8`.
- Common Crossplane API types now come from
  `github.com/crossplane/crossplane/apis/v2/core/v2` (`xpv1`/`xpv2` merged).
- Conversion webhooks are registered by the generated
  `SetupWebhookWithManager*` aggregators; the hand-written `ProviderConfig`
  packages expose a no-op `SetupWebhookWithManager`.
- `apis/generate.go` now also deletes stale `zz_generated.managed*.go` so a
  change of the common API import path regenerates cleanly.

### Backwards compatibility breaks

- CRD shape: `Pullzone.spec.forProvider.routing` and
  `Shield.spec.forProvider.botDetection` change from a one-element **array**
  to an **object** (Terraform `nesting_mode: single`). Manifests must change
  `routing: [{...}]` to `routing: {...}` (same for `botDetection`).

### Added

- Regression tests in `config/` for external-name wiring completeness and for
  numeric Terraform state ids.
- `docs/architecture/upjet-upgrade-2026-10.md` and
  `docs/architecture/upjet-fork.md`.

### Fixed

- README resource count (24, not 23).

## [0.3.0] - 2026-08-14

Upgrades the upstream `BunnyWay/bunnynet` Terraform provider from `v0.15.1` to
`v0.18.0` and regenerates the provider schema, CRDs and controllers.

### Backwards compatibility breaks

- `Shield` (cdn): the `ddos.mode` attribute was **removed** upstream. Shield Zones
  are always in Blocking mode; sensitivity is configured via `ddos.level`. No
  `Shield` resources existed in our control plane at the time of the upgrade, so
  the removal is inert for us — but a manifest carrying `ddos.mode` will now fail
  schema validation.

### Added

- `Shield` (cdn): `uploadScanningAntivirus`, `uploadScanningCsam` — Upload Scanning.
- `Shield` (cdn): `whitelabelBlock`, `whitelabelChallenge`, `whitelabelRateLimit` —
  Whitelabel custom pages.
- `Pullzone` (cdn): `originScriptExecuteBeforeCache`.
- `Zone` (storage): `hostnameS3`, exposed separately now that upstream reverted
  `hostname` to always return the Standard hostname regardless of `type`.

### Fixed (upstream)

- `Record` (dns): changing `zone` now triggers a replacement instead of an
  in-place update that could not succeed.

### Notes

- The **resource list is unchanged** (23 resources). `0.18.0` adds Terraform
  *actions* (`pullzone_cache_purge`, `url_cache_purge`), which upjet does not
  generate CRDs for.
- `stream_library.api_key_readonly`, added upstream in `0.15.2`, is **not** exposed:
  upjet excludes it as a sensitive field.
- `Makefile`: the schema step now runs `tofu init -upgrade`. Without it, a
  `.terraform.lock.hcl` left in the local cache by a previous version makes every
  version bump fail with "locked provider does not match configured version
  constraint". CI never hit this because it always starts cold.

## [0.2.3] - 2026-07-10

Ships the account sub-provider package. The `Subuser` API and CRDs were
introduced in 0.2.2, but the publish matrix did not build a
`provider-bunnynet-account` package, so the resource could not be installed.
This release adds `account` to the publish matrix and repository list.

### Added

- New `provider-bunnynet-account` family sub-provider package, shipping the `Subuser` managed resource (`bunnynet_account_subuser`) introduced in 0.2.2. Install it alongside `provider-family-bunnynet` to manage account subusers.

## [0.2.2] - 2026-07-10

Restores the full code-generation pipeline and exposes a new managed resource
for account subusers. Purely additive — no breaking changes.

### Added

- New `Subuser` managed resource (`account` group, from `bunnynet_account_subuser`) for managing account team members / subusers. Available in both the cluster-scoped (`account.upjet-bunnynet.upbound.io`) and namespaced (`account.upjet-bunnynet.m.upbound.io`) API groups. This brings the provider to 24 managed resources per scope (was 23).

### Fixed

- Restored the full `//go:generate` pipeline in `apis/generate.go`. Previously `make generate` regenerated only the Terraform provider schema; the Upjet generator, `controller-gen` (deepcopy + CRD manifests), and `angryjet` (managed-resource methodsets) were not wired up, so CRDs, deepcopy functions, and methodsets drifted from the committed sources. `make generate` now regenerates everything in one shot and is idempotent (a no-op regen produces byte-identical output).
- As a result of the restored pipeline, `config/provider-metadata.yaml` is now populated from the upstream Terraform docs (previously a placeholder stub), `config/generated.lst` is generated for the `schema-version-diff` CI check, and all CRDs now carry proper field descriptions.
- Added the `github.com/crossplane/crossplane-tools` (angryjet) and `github.com/dave/jennifer` module dependencies required by the restored generator pipeline.
- Excluded the generated-but-unshipped `cmd/provider/monolith/` main from version control so `make check-diff` stays clean after `make generate`.

## [0.2.1] - 2026-07-09

### Changed

- Upgraded the upstream BunnyWay/bunnynet Terraform provider from `v0.14.0` to `v0.15.1`; provider schema, CRDs, and controllers were regenerated.

## [0.2.0] - 2026-04-10

### Added

- Initial family-only release of the Crossplane provider for bunny.net, built with Upjet. Ships cluster-scoped and namespaced managed resources across the `cdn`, `dns`, `storage`, `stream`, and `compute` groups.
