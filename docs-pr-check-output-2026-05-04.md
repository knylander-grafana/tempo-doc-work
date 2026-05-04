# Documentation PR check: merged Tempo PRs from 2026-04-27 to 2026-05-04

Assessment window: 2026-04-27 13:00 UTC through 2026-05-04 13:00 UTC.

Source repository: https://github.com/grafana/tempo

## Table

| PR | Title | Classification | Notes |
|----|-------|----------------|-------|
| [#7096](https://github.com/grafana/tempo/pull/7096) | Update Prometheus to v0.311.3 and related dependencies | Docs update needed | Dependency update exposes new user-configurable fields from dskit/memberlist/server config. The PR updates only generated `docs/sources/tempo/configuration/manifest.md`; the hand-edited configuration reference does not mention new fields such as `server.grpc_server_read_buffer_size`, `server.grpc_server_write_buffer_size`, `internal_server.grpc_server_read_buffer_size`, `internal_server.grpc_server_write_buffer_size`, `memberlist.rejoin_seed_nodes`, or `memberlist.propagation_delay_tracker`. Update `docs/sources/tempo/configuration/_index.md`. |
| [#7094](https://github.com/grafana/tempo/pull/7094) | fix(deps): update module github.com/minio/minio-go/v7 to v7.1.0 (main) | No docs required | Dependency bump only; no Tempo user-facing feature, option, API, or behavior change identified. |
| [#7090](https://github.com/grafana/tempo/pull/7090) | chore(deps): update grafana/alloy docker tag to v1.16.0 (main) | No docs required | Example Docker Compose image tag update only; no Tempo documentation change needed. |
| [#7089](https://github.com/grafana/tempo/pull/7089) | chore(deps): update k8s.io to v0.36.0 (main) | No docs required | Dependency bump only; no Tempo user-facing feature, option, API, or behavior change identified. |
| [#7088](https://github.com/grafana/tempo/pull/7088) | Enable RetryInfo by default on ResourceExhausted errors | Docs present | User-facing default behavior change is documented. PR checked "Documentation added" and updated `docs/sources/tempo/configuration/_index.md` and generated manifest. Existing docs now show `distributor.retry_after_on_resource_exhausted` default `5s` and `overrides.defaults.ingestion.retry_info_enabled` default `true`. Link: `/docs/tempo/<TEMPO_VERSION>/configuration/`. |
| [#7086](https://github.com/grafana/tempo/pull/7086) | fix(deps): update module github.com/twmb/franz-go/pkg/kadm to v1.18.0 (main) | No docs required | Dependency bump only; no Tempo user-facing feature, option, API, or behavior change identified. |
| [#7085](https://github.com/grafana/tempo/pull/7085) | fix(deps): update module github.com/twmb/franz-go to v1.21.0 (main) | No docs required | Dependency bump only; no Tempo user-facing feature, option, API, or behavior change identified. |
| [#7084](https://github.com/grafana/tempo/pull/7084) | fix(deps): update module github.com/jedib0t/go-pretty/v6 to v6.7.10 (main) | No docs required | Dependency bump only; no Tempo user-facing feature, option, API, or behavior change identified. |
| [#7078](https://github.com/grafana/tempo/pull/7078) | Revert "Remove rf1_after from query frontend and query paths (#6969)" | Docs update needed | User-facing revert on the `r249` release branch restores the previous `rf1_after` query behavior/configuration path. The PR updates generated manifest/examples but no hand-authored docs were found for `rf1_after` in `docs/sources/tempo`. Add or restore guidance in the appropriate release/upgrade/configuration docs so users know whether `query_frontend.rf1_after` is required for mixed RF3/RF1 stores on that release line. |
| [#7073](https://github.com/grafana/tempo/pull/7073) | docs: cover KEDA autoscaling and partition ring config removal | Docs present | Documentation-only PR. It updates the Tanka deployment guide with KEDA autoscaling details and the upgrade guide for removed `partition_ring_live_store` config. Links: `/docs/tempo/<TEMPO_VERSION>/set-up-for-tracing/setup-tempo/deploy/kubernetes/tanka/` and `/docs/tempo/<TEMPO_VERSION>/set-up-for-tracing/setup-tempo/upgrade/`. |
| [#7072](https://github.com/grafana/tempo/pull/7072) | [DOC] Update network docs for 3.0 changes | Docs present | Documentation-only PR. It updates IPv6, sidecar proxy, and TLS network docs for Tempo 3.0. Links: `/docs/tempo/<TEMPO_VERSION>/configuration/network/ipv6/`, `/docs/tempo/<TEMPO_VERSION>/configuration/network/sidecar-proxy/`, and `/docs/tempo/<TEMPO_VERSION>/configuration/network/tls/`. |
| [#7047](https://github.com/grafana/tempo/pull/7047) | [DOC] Updated configuration blocks, overrides, and caches config doc | Docs present | Documentation-only PR. It updates `docs/sources/tempo/configuration/_index.md` for Tempo 3.0 configuration blocks, overrides, and cache fields. Link: `/docs/tempo/<TEMPO_VERSION>/configuration/`. |
| [#6982](https://github.com/grafana/tempo/pull/6982) | tempo-cli: add migrate config command for 2.x to 3.0 migration | Docs present | New user-facing CLI command is documented. PR added `docs/sources/tempo/operations/tempo_cli.md` content for `tempo-cli migrate config` and linked it from the Tempo 3.0 upgrade guide. Links: `/docs/tempo/<TEMPO_VERSION>/operations/tempo_cli/#migrate-config-command` and `/docs/tempo/<TEMPO_VERSION>/set-up-for-tracing/setup-tempo/upgrade/`. |

## Gap summary

1. **PR [#7096](https://github.com/grafana/tempo/pull/7096): update configuration reference for newly exposed dskit/server/memberlist fields.**
   - Highest priority because these are configuration options visible to operators after the dependency update.
   - Update `docs/sources/tempo/configuration/_index.md`.
   - Fields to verify and document include:
     - `server.grpc_server_read_buffer_size`
     - `server.grpc_server_write_buffer_size`
     - `internal_server.grpc_server_read_buffer_size`
     - `internal_server.grpc_server_write_buffer_size`
     - `memberlist.rejoin_seed_nodes`
     - `memberlist.propagation_delay_tracker.enabled`
     - `memberlist.propagation_delay_tracker.beacon_interval`
     - `memberlist.propagation_delay_tracker.beacon_lifetime`
     - Additional `propagation_delay_tracker` fields from the generated manifest, if present.
   - Also verify whether the `memberlist.join_members` generated-manifest rendering change from `[]` to `""` needs a short note. The PR says existing list-style configs continue to be accepted, so this is not a breaking change.

2. **PR [#7078](https://github.com/grafana/tempo/pull/7078): clarify `rf1_after` behavior for the affected release line.**
   - The revert restores the old query path behavior/configuration after #6969.
   - No hand-authored `rf1_after` documentation was found under `docs/sources/tempo`.
   - Add or restore a release/upgrade/configuration note explaining when users need `query_frontend.rf1_after` for mixed RF3/RF1 block stores on the `r249` release branch.

3. **No entirely new standalone docs pages appear necessary for the remaining PRs.**
   - PRs #7088, #7073, #7072, #7047, and #6982 include adequate documentation.
   - PRs #7094, #7090, #7089, #7086, #7085, and #7084 are dependency/example updates without identified Tempo user-facing documentation needs.
