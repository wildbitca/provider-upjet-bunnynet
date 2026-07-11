# Changelog

All notable changes to this provider family are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
