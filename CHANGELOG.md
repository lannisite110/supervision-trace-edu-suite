# Changelog

All notable changes to `supervision-trace-edu-suite` are documented here.

## [0.2.0] — 2026-06-19

### Added

- Fabric sandbox K8s overlay (`fabric-sandbox-configmap.yaml`) shared across 3 plugins
- Joint-debug smoke script covering evaluate / simulate / status / report for all labs
- Dedicated Vue dev-frontend preview (`make dev-frontend` on `:5174`)
- Compliance footers on all tutorial docs

### Changed

- Plugin manifests bumped to **0.2.0**; `coreVersion` aligned to `>=0.3.0 <0.4.0`
- Integration docs updated for **web3-edu-platform-core** v0.3 (plugin registry, gateway, scheduler)

### Compliance

- `fabric-local` only — no mainnet
- Fictional data only — no real regulatory, payment, or EMR integration

## [0.1.0] — 2026-06-19

Initial release: 3 CN trace labs (food / medical / charity) with chaincode, rules, Vue panels, and K8s job templates.
