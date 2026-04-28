# Docs PR check: Tempo PRs merged 2026-04-21 through 2026-04-28

Source repository: `grafana/tempo`

Date range: previous 7 days from the automation run on 2026-04-28.

Method: Followed `.claude/skills/docs-pr-check/SKILL.md`. For each merged PR, checked PR title/body, labels, changed files, documentation checklist state, docs changes under `docs/`, and searched `docs/sources/tempo` for existing coverage of user-facing features.

## Table

| PR | Title | Classification | Notes |
|----|-------|----------------|-------|
| [#7084](https://github.com/grafana/tempo/pull/7084) | fix(deps): update module github.com/jedib0t/go-pretty/v6 to v6.7.10 (main) | No docs required | Dependency update only. |
| [#7079](https://github.com/grafana/tempo/pull/7079) | fix(deps): update module github.com/bytedance/sonic to v1.15.1 (main) | No docs required | Dependency update only. |
| [#7078](https://github.com/grafana/tempo/pull/7078) | Revert "Remove rf1_after from query frontend and query paths (#6969)" | Docs update needed | Reintroduces `query_frontend.rf1_after` behavior and vulture flag support. Only `configuration/manifest.md` was updated; no coverage found in `configuration/_index.md` or upgrade guidance. |
| [#7076](https://github.com/grafana/tempo/pull/7076) | chore(deps): lock file maintenance (release-v2.10) | No docs required | Lock file maintenance. |
| [#7075](https://github.com/grafana/tempo/pull/7075) | chore(deps): lock file maintenance (main) | No docs required | Lock file maintenance. |
| [#7071](https://github.com/grafana/tempo/pull/7071) | [release-v2.9] doc: added 2.8.4 security fixes to release notes | Docs present | Updates `docs/sources/tempo/release-notes/v2-8.md`. |
| [#7070](https://github.com/grafana/tempo/pull/7070) | [release-v2.9] doc: added 2.9.2 security fixes to release notes | Docs present | Updates `docs/sources/tempo/release-notes/v2-9.md`. |
| [#7069](https://github.com/grafana/tempo/pull/7069) | [release-v2.8] doc: added 2.8.4 security fixes to release notes | Docs present | Updates `docs/sources/tempo/release-notes/v2-8.md`. |
| [#7066](https://github.com/grafana/tempo/pull/7066) | Measure incorrect results | Docs needed | Adds `tempo_live_store_lagged_requests_total` metric for requests affected by live-store Kafka lag. No docs found for the metric or how operators should use it. |
| [#7063](https://github.com/grafana/tempo/pull/7063) | feat: add span profiling support | Docs needed | Adds top-level `span_profiling` YAML option and `-span-profiling` flag. No docs found in configuration or command-line flag docs. |
| [#7062](https://github.com/grafana/tempo/pull/7062) | chore(deps): update grafana/grafana docker tag to v13 (main) | No docs required | Docker example dependency tag update only. |
| [#7061](https://github.com/grafana/tempo/pull/7061) | fix(ci): use [0-9]+ in publish-technical-documentation-release regex (release-v2.8) | No docs required | CI/docs publishing workflow fix. |
| [#7060](https://github.com/grafana/tempo/pull/7060) | chore: pinned tempo image to 2.8.4 in docs and jsonnets | Docs present | Updates docs and examples to pin Tempo image version. |
| [#7059](https://github.com/grafana/tempo/pull/7059) | chore: pinned tempo image to 2.9.2 in docs and jsonnets | Docs present | Updates docs and examples to pin Tempo image version. |
| [#7058](https://github.com/grafana/tempo/pull/7058) | [release-v2.10] doc: added 2.9.2 security fixes to release notes | Docs present | Updates `docs/sources/tempo/release-notes/v2-9.md`. |
| [#7057](https://github.com/grafana/tempo/pull/7057) | [release-v2.10] doc: added 2.8.4 security fixes to release notes | Docs present | Updates `docs/sources/tempo/release-notes/v2-8.md`. |
| [#7055](https://github.com/grafana/tempo/pull/7055) | doc: added 2.8.4 security fixes to release notes | Docs present | Updates `docs/sources/tempo/release-notes/version-2/v2-8.md`. |
| [#7054](https://github.com/grafana/tempo/pull/7054) | doc: added 2.9.2 security fixes to release notes | Docs present | Updates `docs/sources/tempo/release-notes/version-2/v2-9.md`. |
| [#7053](https://github.com/grafana/tempo/pull/7053) | chore: updated tempo image tag in docs and jsonnets | Docs present | Updates docs/examples/jsonnet image references for Tempo 2.10.5. |
| [#7052](https://github.com/grafana/tempo/pull/7052) | chore: cut changelog for v2.8.4 | No docs required | Release maintenance. |
| [#7051](https://github.com/grafana/tempo/pull/7051) | chore: cut changelog for v2.9.2 | No docs required | Release maintenance. |
| [#7050](https://github.com/grafana/tempo/pull/7050) | docs: add v2.10.5 release notes entry | Docs present | Updates `docs/sources/tempo/release-notes/version-2/v2-10.md`. |
| [#7049](https://github.com/grafana/tempo/pull/7049) | chore: cut changelog for v2.10.5 | No docs required | Release maintenance. |
| [#7048](https://github.com/grafana/tempo/pull/7048) | chore: remove deprecated query_live_store config | Docs update needed | Breaking removal of `querier.query_live_store`; docs search found no migration/upgrade guidance. Add upgrade consideration telling users to remove the field. |
| [#7046](https://github.com/grafana/tempo/pull/7046) | [DOC] Update storage config doc for 3.0 changes | Docs present | Updates `docs/sources/tempo/configuration/_index.md`. |
| [#7045](https://github.com/grafana/tempo/pull/7045) | [DOC] Move v2 release notes into version-2 subdirectory | Docs present | Documentation structure change. |
| [#7044](https://github.com/grafana/tempo/pull/7044) | chore(deps): update k8s.io to v0.35.4 (main) | No docs required | Dependency update only. |
| [#7042](https://github.com/grafana/tempo/pull/7042) | chore(deps): update docker (release-v2.10) | No docs required | Dependency/Docker image update only. |
| [#7041](https://github.com/grafana/tempo/pull/7041) | chore(deps): update docker (release-v2.9) | No docs required | Dependency/Docker image update only. |
| [#7040](https://github.com/grafana/tempo/pull/7040) | chore(deps): update dependency go to v1.26.2 (release-v2.10) | No docs required | Dependency update only. |
| [#7039](https://github.com/grafana/tempo/pull/7039) | feat: Add flag for MCP server | Docs present | Updates MCP server docs and command-line flag docs for `--query-frontend.mcp-server.enabled`. |
| [#7038](https://github.com/grafana/tempo/pull/7038) | [release-v2.10] [DOC]Fix typo for 2.10 release notes | Docs present | Documentation-only typo fix. |
| [#7037](https://github.com/grafana/tempo/pull/7037) | chore(deps): update dependency go to v1.26.2 (release-v2.9) | No docs required | Dependency update only. |
| [#7036](https://github.com/grafana/tempo/pull/7036) | chore(deps): update docker (release-v2.8) | No docs required | Dependency/Docker image update only. |
| [#7035](https://github.com/grafana/tempo/pull/7035) | chore(deps): update dependency go to v1.26.2 (release-v2.8) | No docs required | Dependency update only. |
| [#7034](https://github.com/grafana/tempo/pull/7034) | chore: bump ingestion limits | Docs present | Updates configuration reference, generated manifest, and troubleshooting example for new ingestion defaults. |
| [#7033](https://github.com/grafana/tempo/pull/7033) | chore: make tempo singlebinary example to use the local backend | Docs present | Updates `docs/sources/tempo/docker-example.md`. |
| [#7032](https://github.com/grafana/tempo/pull/7032) | Update CI tools image tag to correct format (release-v2.10) | No docs required | CI/tooling update. |
| [#7031](https://github.com/grafana/tempo/pull/7031) | Update CI tools image tag to correct format (release-v2.9) | No docs required | CI/tooling update. |
| [#7030](https://github.com/grafana/tempo/pull/7030) | Allow Go and CI tools image updates on release branches | No docs required | Renovate/CI configuration update. |
| [#7029](https://github.com/grafana/tempo/pull/7029) | [DOC] 3.0 config docs update for query, querier, memberlist | Docs present | Updates configuration reference. |
| [#7028](https://github.com/grafana/tempo/pull/7028) | 3.0 config docs:ingest, block-builder, live-store, metrics-gen | Docs present | Updates configuration reference. |
| [#7027](https://github.com/grafana/tempo/pull/7027) | [DOC]Fix typo for 2.10 release notes | Docs present | Documentation-only typo fix. |
| [#7026](https://github.com/grafana/tempo/pull/7026) | [DOC] Distinguish monolithic deploymemt mode (kafkaless) and microservices modes in docs | Docs present | Updates configuration and architecture/deployment docs. |
| [#7025](https://github.com/grafana/tempo/pull/7025) | chore(deps): update grafana/tempo-ci-tools docker tag to v20260421 (main) | No docs required | CI tools dependency update. |
| [#7024](https://github.com/grafana/tempo/pull/7024) | fix(deps): update module github.com/googleapis/gax-go/v2 to v2.22.0 (main) | No docs required | Dependency update only. |
| [#7023](https://github.com/grafana/tempo/pull/7023) | fix(deps): update module google.golang.org/api to v0.276.0 (main) | No docs required | Dependency update only. |
| [#7022](https://github.com/grafana/tempo/pull/7022) | fix: Fix golang Renovate updates | No docs required | Renovate/dependency maintenance. |
| [#7021](https://github.com/grafana/tempo/pull/7021) | live-store: better o11y for block creation | No docs required | Adds internal tracing attributes/events for live-store block creation; no user-facing config or behavior change. |
| [#7020](https://github.com/grafana/tempo/pull/7020) | chore(deps): update grafana/grafana docker tag to v12.4.3 (main) | No docs required | Docker example dependency tag update only. |
| [#7019](https://github.com/grafana/tempo/pull/7019) | blockbuilder: wait for blocklist poller on test cleanup | No docs required | Test cleanup only. |
| [#7018](https://github.com/grafana/tempo/pull/7018) | Vulture: fix for query_end_cutoff | No docs required | Vulture bug fix; no docs surface changed. |
| [#7015](https://github.com/grafana/tempo/pull/7015) | Updates for config: server, memory, global strategy overrides, links | Docs present | Updates API, configuration, and operations docs. |
| [#7007](https://github.com/grafana/tempo/pull/7007) | chore(deps): replace stale dependency github.com/willf/bloom | No docs required | Dependency replacement only. |
| [#7006](https://github.com/grafana/tempo/pull/7006) | chore(deps): replace stale dependency github.com/json-iterator/go | No docs required | Dependency replacement only. |
| [#7005](https://github.com/grafana/tempo/pull/7005) | chore(deps): replace stale dependency github.com/golang/snappy | No docs required | Dependency replacement only. |
| [#6993](https://github.com/grafana/tempo/pull/6993) | [docs]: document backend scheduler/worker operational details | Docs present | Updates API, configuration, and compaction component docs. |
| [#6988](https://github.com/grafana/tempo/pull/6988) | [docs]: add compaction processs documentation | Docs present | Adds/updates compaction and polling docs. |
| [#6975](https://github.com/grafana/tempo/pull/6975) | [DOC] Restore linux and local test docs; updates for 3.0 | Docs present | Restores and updates setup/deployment docs. |
| [#6971](https://github.com/grafana/tempo/pull/6971) | [Bugfix] Live-store: Cut head block if it exceeds MaxBlockBytes | No docs required | Bug fix to honor existing `MaxBlockBytes` behavior; no new user-facing option. |
| [#6970](https://github.com/grafana/tempo/pull/6970) | Add KEDA autoscaling to microservices jsonnet | Docs update needed | Adds disabled-by-default KEDA autoscaling support to microservices Jsonnet. Existing docs only mention the autoscaling file from compaction docs; add deployment/Jsonnet docs explaining enablement and supported scalers. |
| [#6954](https://github.com/grafana/tempo/pull/6954) | fix(deps): update module github.com/prometheus/prometheus to v0.311.2 [security] (main) | No docs required | Dependency security update; no direct user-facing docs change required by docs-pr-check criteria. |
| [#6575](https://github.com/grafana/tempo/pull/6575) | [metrics-generator]: configure partition reassignment behaviour on shutdown | Docs present | Updates metrics-generator docs for `leave_consumer_group_on_shutdown` and generated configuration manifest. |

## Gap summary

Prioritized documentation work:

1. **PR [#7063](https://github.com/grafana/tempo/pull/7063): Span profiling support**
   - Add documentation for `span_profiling: true` and the `-span-profiling` CLI flag.
   - Suggested locations:
     - `docs/sources/tempo/configuration/_index.md`
     - `docs/sources/tempo/set-up-for-tracing/setup-tempo/command-line-flags.md`
     - Consider an operations/tracing note explaining that OTel tracing must be configured and that this adds Pyroscope pprof labels to spans.

2. **PR [#7066](https://github.com/grafana/tempo/pull/7066): live-store lagged request metric**
   - Document `tempo_live_store_lagged_requests_total` and when it increments.
   - Suggested location: operations/monitoring or live-store operational docs.
   - Include guidance for interpreting the `method` label and correlating with Kafka lag / incomplete recent query results.

3. **PR [#7048](https://github.com/grafana/tempo/pull/7048): `querier.query_live_store` removal**
   - Add upgrade guidance that this deprecated field is removed and must be deleted from configs before upgrading.
   - Suggested location: Tempo 3.0 upgrade considerations / release notes, with an example before/after config snippet.

4. **PR [#7078](https://github.com/grafana/tempo/pull/7078): `query_frontend.rf1_after` reintroduced**
   - Update manually maintained config docs for `query_frontend.rf1_after`.
   - Clarify current status after the revert and any relationship to migration from RF3 to RF1 blocks.
   - Suggested location: `docs/sources/tempo/configuration/_index.md`; release notes may also need correction if they still state the field is ignored or deprecated.

5. **PR [#6970](https://github.com/grafana/tempo/pull/6970): KEDA autoscaling for microservices Jsonnet**
   - Add documentation explaining how to enable the disabled-by-default KEDA autoscaling Jsonnet and which components/scalers are supported:
     - distributor CPU
     - metrics-generator CPU
     - backend-worker Prometheus outstanding blocks
     - block-builder Kubernetes workload scaler
   - Suggested locations:
     - `docs/sources/tempo/set-up-for-tracing/setup-tempo/deploy/kubernetes/tanka.md`
     - relevant operations Jsonnet docs
     - optionally expand the existing compaction docs cross-reference to point to a dedicated autoscaling section.

Uncertain / needs engineering confirmation:

- **PR [#7078](https://github.com/grafana/tempo/pull/7078)**: Confirm whether `query_frontend.rf1_after` should be documented as an active advanced migration option, and whether earlier docs or release notes from #6969 need to be reverted/updated.
- **PR [#6970](https://github.com/grafana/tempo/pull/6970)**: Confirm the recommended public docs surface for operations Jsonnet/KEDA examples, since the implementation is in `operations/jsonnet/microservices` rather than Helm chart docs.
