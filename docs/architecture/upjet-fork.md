# Why this provider pins the wildbitca/upjet fork

`go.mod` replaces `github.com/crossplane/upjet/v2` with
`github.com/wildbitca/upjet/v2 v2.5.1-wb.1` (upstream `v2.5.1` plus a small
patch set). The replace is required because of one patch.

## Numeric Terraform state ids (required)

`BunnyWay/bunnynet` is a Terraform Plugin Framework provider. 16 of its 24
resources declare `id` as a **number**, so the state written by the CLI runtime
contains `{"id": 577654}`. Upstream upjet v2.5.1 only accepts string ids:

- `config.IdentifierFromProvider.GetExternalNameFn` returns
  `cannot find id in tfstate` for a numeric id;
- `terraform.FileProducer.isStateEmpty` fails with
  `cannot work with a non-string id`;
- the next reconcile then fails decoding the observation (`ReadString`).

`config/external_name_test.go` (`TestNumericStateIDAsExternalName`) reproduces
this: it passes with the fork and fails against plain upstream v2.5.1
(verified by dropping the replace temporarily). The fork adds
`resource.NormalizeStateID` and float64-aware `IDAsExternalName` /
`isStateEmpty`.

## Other fork patches

| Patch | Used here |
|-------|-----------|
| Skip the CEL rule for required dynamic parameters | No (no dynamic fields) |
| Skip the synthetic TF state when the external ID is empty | Yes, replaces `patches/apply.sh` |

`GetImportIDFn` was dropped from the fork: upstream has it since v2.4.0.

## Removing the replace

Drop the replace once upstream accepts numeric state ids. Keep
`TestNumericStateIDAsExternalName` as the gate: when it passes against plain
upstream, the replace is no longer needed.
