# Docs PR check: merged Tempo PRs for 2026-05-11 to 2026-05-18

Evaluated merged PRs in `grafana/tempo` for the previous 7 days from the automation trigger:

- Window start: `2026-05-11T13:02:56Z`
- Window end: `2026-05-18T13:02:56Z`
- Skill: `.claude/skills/docs-pr-check/SKILL.md`

## Table

| PR | Title | Classification | Notes |
|----|-------|----------------|-------|
| [grafana/tempo#7176](https://github.com/grafana/tempo/pull/7176) | overrides: drop merge for processors in user-configurable overrides | Docs present | User-facing breaking behavior change for `metrics_generator.processors` precedence. PR includes docs updates in `docs/sources/tempo/operations/manage-advanced-systems/user-configurable-overrides.md` and `docs/sources/tempo/set-up-for-tracing/setup-tempo/migrate-to-3.md`. |
| [grafana/tempo#7174](https://github.com/grafana/tempo/pull/7174) | chore(livestore): log isLagged inputs as reasons when it returns true | No docs required | Logging-only diagnostic change for live-store internals; no user-facing config, API, or behavior change. |
| [grafana/tempo#7163](https://github.com/grafana/tempo/pull/7163) | enhancement(livestore): expose inspected bytes for metric queries | Docs update needed | Extends inspected-byte accounting for TraceQL metrics queries and makes `tempo_query_frontend_bytes_inspected_total{op="metrics"}` non-zero. Existing query I/O docs mention inspected bytes generally but do not cover the new live-store metric/query behavior. Update `docs/sources/tempo/operations/monitor/query-io-and-timestamp-distance.md` and/or `docs/sources/tempo/reference-tempo-architecture/components/live-store.md`. |
| [grafana/tempo#7162](https://github.com/grafana/tempo/pull/7162) | enhancement(livestore): expose query inspected bytes as a metric | Docs update needed | Adds `tempo_live_store_query_inspected_bytes_total` with `tenant` and `op` labels. Existing live-store metric docs do not list this metric. Update `docs/sources/tempo/reference-tempo-architecture/components/live-store.md`; consider linking from `docs/sources/tempo/operations/monitor/query-io-and-timestamp-distance.md`. |
| [grafana/tempo#7155](https://github.com/grafana/tempo/pull/7155) | fix(livestore): write query-range cache atomically; tolerate cache errors | No docs required | Internal cache correctness fix; no new user-facing configuration or documented workflow. |
| [grafana/tempo#7154](https://github.com/grafana/tempo/pull/7154) | jsonnet: configurable block-builder autoscaling approach (rollout-operator or keda) | Docs present | New Jsonnet/Tanka autoscaling option documented in `docs/sources/tempo/set-up-for-tracing/setup-tempo/deploy/kubernetes/tanka.md`, including `live_store.keda.block_builder_scaling`, `rollout-operator`, and `keda` behavior. |
| [grafana/tempo#7153](https://github.com/grafana/tempo/pull/7153) | [backend-scheduler]: source SubmitRedaction tenant from auth context only | No docs required | Security hardening for an API noted in the PR as not currently externally exposed; no public docs gap identified. |
| [grafana/tempo#7152](https://github.com/grafana/tempo/pull/7152) | feat: add tempodb_cache_store_size_bytes histogram labelled by role | Docs needed | Adds user-visible operational metric `tempodb_cache_store_size_bytes` for backend cache item sizes. No docs found for this metric or cache-size monitoring. Add monitoring documentation with metric type, labels, and tuning use case. |
| [grafana/tempo#7148](https://github.com/grafana/tempo/pull/7148) | enhancement: add TempoDistributorKafkaProduceFailing alert | Docs present | Adds a new Tempo mixin alert and a matching runbook entry in `operations/tempo-mixin/runbook.md`. Existing docs link users to the mixin alerts and runbook. |
| [grafana/tempo#7141](https://github.com/grafana/tempo/pull/7141) | [backendscheduler]: eliminate O(256.M) shard scan in HasJobsForTenant and O(N) prefix scan in BusyBlocksForTenant | No docs required | Internal backend-scheduler performance optimization with no user-facing behavior change. |
| [grafana/tempo#7139](https://github.com/grafana/tempo/pull/7139) | [Backport r251] Fix panic in old tag-style search + integer dedicated columns + vp5 | No docs required | Backport of a bug fix for a panic; no docs gap identified beyond changelog/release-note mention. |
| [grafana/tempo#7138](https://github.com/grafana/tempo/pull/7138) | fix: clean up YAML output on overrides status endpoint | No docs required | User-configurable overrides API serialization bug fix; existing docs describe the schema rather than the prior incorrect YAML output. Optional: add an API response example if YAML output needs explicit documentation. |
| [grafana/tempo#7132](https://github.com/grafana/tempo/pull/7132) | livestore: lock-free block reads via atomic snapshot + crash-safe deletion | Docs present | Adds `live_store.block_reclaim_grace`; PR updates `docs/sources/tempo/configuration/_index.md` and `docs/sources/tempo/configuration/manifest.md`. |
| [grafana/tempo#7129](https://github.com/grafana/tempo/pull/7129) | fix(deps): update module golang.org/x/net to v0.53.0 [security] (release-v2.10) | No docs required | Dependency/security bump only. |
| [grafana/tempo#7128](https://github.com/grafana/tempo/pull/7128) | fix(deps): update module golang.org/x/net to v0.53.0 [security] (release-v2.9) | No docs required | Dependency/security bump only. |
| [grafana/tempo#7127](https://github.com/grafana/tempo/pull/7127) | fix(deps): update module golang.org/x/net to v0.53.0 [security] (release-v2.8) | No docs required | Dependency/security bump only. |
| [grafana/tempo#7122](https://github.com/grafana/tempo/pull/7122) | Tempo 3.0 migration: docker-compose example + doc tweaks | Docs present | Docs-focused PR updating `docs/sources/tempo/set-up-for-tracing/setup-tempo/migrate-to-3.md` and adding `example/docker-compose/migrate-to-3/`. |
| [grafana/tempo#7121](https://github.com/grafana/tempo/pull/7121) | [DOC] Tempo 3.0 release notes | Docs present | Docs-only PR adding Tempo 3.0 release notes and updating upgrade/release-note pages. |
| [grafana/tempo#7118](https://github.com/grafana/tempo/pull/7118) | chore(deps): update module github.com/apache/thrift to v0.23.0 [security] (release-v2.8) | No docs required | Dependency/security bump only. |
| [grafana/tempo#7117](https://github.com/grafana/tempo/pull/7117) | chore(deps): update module github.com/apache/thrift to v0.23.0 [security] (main) | No docs required | Dependency/security bump only. |
| [grafana/tempo#7110](https://github.com/grafana/tempo/pull/7110) | fix(deps): update opentelemetry-collector (main) | No docs required | Dependency update only; no Tempo-facing docs change identified. |
| [grafana/tempo#6607](https://github.com/grafana/tempo/pull/6607) | Make search and metrics streaming easier | Docs update needed | Adds `query_frontend.max_grpc_streaming_packet_size` and changes streaming response behavior. PR updates `docs/sources/tempo/configuration/manifest.md`, but `docs/sources/tempo/configuration/_index.md`, streaming docs, and response-size troubleshooting docs do not explain the new option or default. |
| [grafana/tempo#6302](https://github.com/grafana/tempo/pull/6302) | chore(deps): update github.com/grafana/memberlist digest to 4887d8e (main) | No docs required | Dependency digest update only. |

## Gap summary

Prioritized documentation work:

1. **PR [grafana/tempo#7152](https://github.com/grafana/tempo/pull/7152): document `tempodb_cache_store_size_bytes`.**
   - New operational metric has no docs coverage.
   - Suggested content: metric type (histogram), labels (`role`), what item sizes are measured, and how operators can use it for backend cache or Memcached slab tuning.
   - Suggested location: a Tempo monitoring page such as `docs/sources/tempo/operations/monitor/_index.md`, or a dedicated cache-monitoring section if one exists.

2. **PRs [grafana/tempo#7162](https://github.com/grafana/tempo/pull/7162) and [grafana/tempo#7163](https://github.com/grafana/tempo/pull/7163): update query I/O and live-store metric docs.**
   - Add `tempo_live_store_query_inspected_bytes_total` to live-store operational metrics.
   - Explain labels (`tenant`, `op`) and how it relates to `tempo_query_frontend_bytes_inspected_total`.
   - Note that TraceQL metrics queries now report inspected bytes through the query frontend metric as well.
   - Suggested files: `docs/sources/tempo/reference-tempo-architecture/components/live-store.md` and `docs/sources/tempo/operations/monitor/query-io-and-timestamp-distance.md`.

3. **PR [grafana/tempo#6607](https://github.com/grafana/tempo/pull/6607): expand streaming packet-size configuration docs.**
   - `query_frontend.max_grpc_streaming_packet_size` appears in the generated manifest, but not in the main configuration reference or streaming/troubleshooting narrative docs.
   - Suggested files: `docs/sources/tempo/configuration/_index.md`, `docs/sources/tempo/api_docs/_index.md`, `docs/sources/tempo/traceql/query-editor/_index.md`, and/or `docs/sources/tempo/troubleshooting/querying/response-too-large.md`.

4. **Optional clarification: PR [grafana/tempo#7138](https://github.com/grafana/tempo/pull/7138).**
   - The YAML output bug fix likely does not require standalone docs because the existing API docs describe the schema, not the broken output.
   - If users rely on YAML responses from `/api/overrides`, consider adding a compact response example to `docs/sources/tempo/operations/manage-advanced-systems/user-configurable-overrides.md`.
