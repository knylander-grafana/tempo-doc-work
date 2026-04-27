# Tempo docs PR check: merged PRs from 2026-04-20 to 2026-04-27

Generated: 2026-04-27

Scope: PRs merged in `grafana/tempo` during the previous 7 days from the automation trigger window, 2026-04-20T13:00:37Z through 2026-04-27T13:00:37Z.

Method: followed `.claude/skills/docs-pr-check/SKILL.md`.

Evidence checked:

- `gh pr list --repo grafana/tempo --state merged --search "merged:2026-04-20..2026-04-27" --limit 100 --json number,title,mergedAt,author,labels,url`
- `gh pr view <PR> --repo grafana/tempo --json title,body,files,labels` for PRs with possible user-facing behavior, configuration, API, deployment, or documentation impact
- Local docs searches under `docs/sources/tempo` for the relevant feature names, configuration keys, metrics, and deployment concepts

## Table

| PR | Title | Classification | Notes |
|----|-------|----------------|-------|
| [#7079](https://github.com/grafana/tempo/pull/7079) | fix(deps): update module github.com/bytedance/sonic to v1.15.1 (main) | No docs required | Dependency patch update. |
| [#7076](https://github.com/grafana/tempo/pull/7076) | chore(deps): lock file maintenance (release-v2.10) | No docs required | Lockfile maintenance. |
| [#7075](https://github.com/grafana/tempo/pull/7075) | chore(deps): lock file maintenance (main) | No docs required | Lockfile maintenance. |
| [#7071](https://github.com/grafana/tempo/pull/7071) | [release-v2.9] doc: added 2.8.4 security fixes to release notes | Docs present | Docs-only release notes update. |
| [#7070](https://github.com/grafana/tempo/pull/7070) | [release-v2.9] doc: added 2.9.2 security fixes to release notes | Docs present | Docs-only release notes update. |
| [#7069](https://github.com/grafana/tempo/pull/7069) | [release-v2.8] doc: added 2.8.4 security fixes to release notes | Docs present | Docs-only release notes update. |
| [#7066](https://github.com/grafana/tempo/pull/7066) | Measure incorrect results | Docs update needed | Adds `tempo_live_store_lagged_requests_total`; no docs found for the new metric. Update live-store/query troubleshooting or monitoring docs. |
| [#7063](https://github.com/grafana/tempo/pull/7063) | feat: add span profiling support | Docs update needed | Adds `-tracing.span-profiling` / `span_profiling` and Pyroscope profile-to-trace correlation. No docs found for `span_profiling`, `pyroscope.profile.id`, or the flag. Update configuration/CLI docs and add usage guidance. |
| [#7062](https://github.com/grafana/tempo/pull/7062) | chore(deps): update grafana/grafana docker tag to v13 (main) | No docs required | Dependency/image update. |
| [#7061](https://github.com/grafana/tempo/pull/7061) | fix(ci): use [0-9]+ in publish-technical-documentation-release regex (release-v2.8) | No docs required | CI workflow fix. |
| [#7060](https://github.com/grafana/tempo/pull/7060) | chore: pinned tempo image to 2.8.4 in docs and jsonnets | No docs required | Release maintenance for docs/jsonnet image tags. |
| [#7059](https://github.com/grafana/tempo/pull/7059) | chore: pinned tempo image to 2.9.2 in docs and jsonnets | No docs required | Release maintenance for docs/jsonnet image tags. |
| [#7058](https://github.com/grafana/tempo/pull/7058) | [release-v2.10] doc: added 2.9.2 security fixes to release notes | Docs present | Docs-only release notes update. |
| [#7057](https://github.com/grafana/tempo/pull/7057) | [release-v2.10] doc: added 2.8.4 security fixes to release notes | Docs present | Docs-only release notes update. |
| [#7055](https://github.com/grafana/tempo/pull/7055) | doc: added 2.8.4 security fixes to release notes | Docs present | Docs-only release notes update. |
| [#7054](https://github.com/grafana/tempo/pull/7054) | doc: added 2.9.2 security fixes to release notes | Docs present | Docs-only release notes update. |
| [#7053](https://github.com/grafana/tempo/pull/7053) | chore: updated tempo image tag in docs and jsonnets | No docs required | Release maintenance for docs/jsonnet image tags. |
| [#7052](https://github.com/grafana/tempo/pull/7052) | chore: cut changelog for v2.8.4 | No docs required | Changelog release maintenance. |
| [#7051](https://github.com/grafana/tempo/pull/7051) | chore: cut changelog for v2.9.2 | No docs required | Changelog release maintenance. |
| [#7050](https://github.com/grafana/tempo/pull/7050) | docs: add v2.10.5 release notes entry | Docs present | Docs-only release notes update. |
| [#7049](https://github.com/grafana/tempo/pull/7049) | chore: cut changelog for v2.10.5 | No docs required | Changelog release maintenance. |
| [#7048](https://github.com/grafana/tempo/pull/7048) | chore: remove deprecated query_live_store config | Docs update needed | Removes a deprecated config option. No current docs mention `query_live_store`, but the removal should be called out in upgrade/migration docs for users carrying older configs. |
| [#7046](https://github.com/grafana/tempo/pull/7046) | [DOC] Update storage config doc for 3.0 changes | Docs present | Docs-only configuration reference update. |
| [#7045](https://github.com/grafana/tempo/pull/7045) | [DOC] Move v2 release notes into version-2 subdirectory | Docs present | Docs-only reorganization. |
| [#7044](https://github.com/grafana/tempo/pull/7044) | chore(deps): update k8s.io to v0.35.4 (main) | No docs required | Dependency update. |
| [#7042](https://github.com/grafana/tempo/pull/7042) | chore(deps): update docker (release-v2.10) | No docs required | Dependency/image update. |
| [#7041](https://github.com/grafana/tempo/pull/7041) | chore(deps): update docker (release-v2.9) | No docs required | Dependency/image update. |
| [#7040](https://github.com/grafana/tempo/pull/7040) | chore(deps): update dependency go to v1.26.2 (release-v2.10) | No docs required | Dependency/toolchain update. |
| [#7039](https://github.com/grafana/tempo/pull/7039) | feat: Add flag for MCP server | Docs present | PR updated `docs/sources/tempo/api_docs/mcp-server.md` and `docs/sources/tempo/set-up-for-tracing/setup-tempo/command-line-flags.md`. |
| [#7038](https://github.com/grafana/tempo/pull/7038) | [release-v2.10] [DOC]Fix typo for 2.10 release notes | Docs present | Docs-only typo fix. |
| [#7037](https://github.com/grafana/tempo/pull/7037) | chore(deps): update dependency go to v1.26.2 (release-v2.9) | No docs required | Dependency/toolchain update. |
| [#7036](https://github.com/grafana/tempo/pull/7036) | chore(deps): update docker (release-v2.8) | No docs required | Dependency/image update. |
| [#7035](https://github.com/grafana/tempo/pull/7035) | chore(deps): update dependency go to v1.26.2 (release-v2.8) | No docs required | Dependency/toolchain update. |
| [#7034](https://github.com/grafana/tempo/pull/7034) | chore: bump ingestion limits | Docs present | PR updated configuration docs, manifest, and max-trace troubleshooting docs. |
| [#7033](https://github.com/grafana/tempo/pull/7033) | chore: make tempo singlebinary example to use the local backend | Docs present | PR updated Docker example docs and example readme alongside the example config. |
| [#7032](https://github.com/grafana/tempo/pull/7032) | Update CI tools image tag to correct format (release-v2.10) | No docs required | CI/tools image maintenance. |
| [#7031](https://github.com/grafana/tempo/pull/7031) | Update CI tools image tag to correct format (release-v2.9) | No docs required | CI/tools image maintenance. |
| [#7030](https://github.com/grafana/tempo/pull/7030) | Allow Go and CI tools image updates on release branches | No docs required | Renovate/CI maintenance. |
| [#7029](https://github.com/grafana/tempo/pull/7029) | [DOC]  3.0 config docs update for query, querier, memberlist | Docs present | Docs-only configuration reference update. |
| [#7028](https://github.com/grafana/tempo/pull/7028) | 3.0 config docs:ingest, block-builder, live-store, metrics-gen | Docs present | Docs-only configuration reference update. |
| [#7027](https://github.com/grafana/tempo/pull/7027) | [DOC]Fix typo for 2.10 release notes | Docs present | Docs-only typo fix. |
| [#7026](https://github.com/grafana/tempo/pull/7026) | [DOC] Distinguish monolithic deploymemt mode (kafkaless) and microservices modes in docs | Docs present | Docs-only architecture/deployment update. |
| [#7025](https://github.com/grafana/tempo/pull/7025) | chore(deps): update grafana/tempo-ci-tools docker tag to v20260421 (main) | No docs required | CI tools image update. |
| [#7024](https://github.com/grafana/tempo/pull/7024) | fix(deps): update module github.com/googleapis/gax-go/v2 to v2.22.0 (main) | No docs required | Dependency update. |
| [#7023](https://github.com/grafana/tempo/pull/7023) | fix(deps): update module google.golang.org/api to v0.276.0 (main) | No docs required | Dependency update. |
| [#7022](https://github.com/grafana/tempo/pull/7022) | fix: Fix golang Renovate updates | No docs required | Renovate configuration fix. |
| [#7021](https://github.com/grafana/tempo/pull/7021) | live-store: better o11y for block creation | No docs required | Adds internal tracing events and error marking for live-store block creation; no stable user-facing config/API/doc surface identified. |
| [#7020](https://github.com/grafana/tempo/pull/7020) | chore(deps): update grafana/grafana docker tag to v12.4.3 (main) | No docs required | Dependency/image update. |
| [#7019](https://github.com/grafana/tempo/pull/7019) | blockbuilder: wait for blocklist poller on test cleanup | No docs required | Test cleanup fix. |
| [#7018](https://github.com/grafana/tempo/pull/7018) | Vulture: fix for query_end_cutoff | No docs required | Internal Vulture check fix; no user-facing docs surface identified. |
| [#7015](https://github.com/grafana/tempo/pull/7015) | Updates for config: server, memory, global strategy overrides, links | Docs present | Docs-only configuration and link update. |
| [#7014](https://github.com/grafana/tempo/pull/7014) | chore(deps): update otel/opentelemetry-collector docker tag to v0.150.1 (main) | No docs required | Dependency/image update. |
| [#7013](https://github.com/grafana/tempo/pull/7013) | chore(deps): update grafana/alloy docker tag to v1.15.1 (main) | No docs required | Dependency/image update. |
| [#7011](https://github.com/grafana/tempo/pull/7011) | fix(deps): update opentelemetry-contrib to v0.150.0 (main) | No docs required | Dependency update. |
| [#7007](https://github.com/grafana/tempo/pull/7007) | chore(deps): replace stale dependency github.com/willf/bloom | No docs required | Dependency replacement. |
| [#7006](https://github.com/grafana/tempo/pull/7006) | chore(deps): replace stale dependency github.com/json-iterator/go | No docs required | Dependency replacement. |
| [#7005](https://github.com/grafana/tempo/pull/7005) | chore(deps): replace stale dependency github.com/golang/snappy | No docs required | Dependency replacement. |
| [#6993](https://github.com/grafana/tempo/pull/6993) | [docs]: document backend scheduler/worker operational details | Docs present | Docs-only API/config/architecture update. |
| [#6988](https://github.com/grafana/tempo/pull/6988) | [docs]: add compaction processs documentation | Docs present | Docs-only operations documentation. |
| [#6975](https://github.com/grafana/tempo/pull/6975) | [DOC] Restore linux and local test docs; updates for 3.0 | Docs present | Docs-only setup/testing update. |
| [#6971](https://github.com/grafana/tempo/pull/6971) | [Bugfix] Live-store: Cut head block if it exceeds MaxBlockBytes | No docs required | Bug fix to enforce existing live-store block-size behavior; no new config or documented workflow change. |
| [#6970](https://github.com/grafana/tempo/pull/6970) | Add KEDA autoscaling to microservices jsonnet | Docs needed | Adds KEDA-based autoscaling support to the OSS microservices Jsonnet library. Docs only contain a brief reference from compaction docs to `operations/jsonnet/microservices/autoscaling.libsonnet`; no user-facing Tanka/Jsonnet setup guidance or `_config.<component>.keda` examples found. |
| [#6954](https://github.com/grafana/tempo/pull/6954) | fix(deps): update module github.com/prometheus/prometheus to v0.311.2 [security] (main) | No docs required | Dependency security update. |
| [#6915](https://github.com/grafana/tempo/pull/6915) | chore(deps): update dependency go to v1.26.2 (main) | No docs required | Dependency/toolchain update. |
| [#6575](https://github.com/grafana/tempo/pull/6575) | [metrics-generator]: configure partition reassignment behaviour on shutdown | Docs present | PR updated `docs/sources/tempo/metrics-from-traces/metrics-generator/_index.md` and `docs/sources/tempo/configuration/manifest.md` for `leave_consumer_group_on_shutdown`; coverage looks adequate. |

## Gap summary

Prioritized documentation work:

1. [#6970](https://github.com/grafana/tempo/pull/6970) - Add docs for KEDA autoscaling in the microservices Jsonnet library.
   - User impact: operators can now enable autoscaling for distributor, metrics-generator, backend-worker, and block-builder through `_config.<component>.keda`, but there is no procedure or example showing how.
   - Suggested docs: update `docs/sources/tempo/set-up-for-tracing/setup-tempo/deploy/kubernetes/tanka.md` or add an operations page for Jsonnet autoscaling. Include prerequisites, disabled-by-default behavior, component examples, and links to `operations/jsonnet/microservices/autoscaling.libsonnet`.

2. [#7063](https://github.com/grafana/tempo/pull/7063) - Document span profiling support.
   - User impact: users need to know how to enable `-tracing.span-profiling` / `span_profiling`, required Pyroscope/OTel environment, expected `pyroscope.profile.id` attribute, and rollout cautions because the PR says it is disabled by default and resource impact should be validated.
   - Suggested docs: update `docs/sources/tempo/configuration/_index.md`, `docs/sources/tempo/set-up-for-tracing/setup-tempo/command-line-flags.md`, and add operational guidance where Tempo self-tracing/profiling is documented.

3. [#7048](https://github.com/grafana/tempo/pull/7048) - Add migration note for removed `query_live_store`.
   - User impact: deployments carrying the deprecated `query_live_store` setting will need to remove it before using the version containing this PR.
   - Suggested docs: update `docs/sources/tempo/set-up-for-tracing/setup-tempo/upgrade.md` or the relevant 3.0 migration section.

4. [#7066](https://github.com/grafana/tempo/pull/7066) - Document `tempo_live_store_lagged_requests_total`.
   - User impact: the new metric helps identify recent-search requests whose requested time range overlaps live-store lag, but it is not described in monitoring or troubleshooting docs.
   - Suggested docs: update live-store monitoring/troubleshooting content, likely under `docs/sources/tempo/operations/monitor/` or `docs/sources/tempo/troubleshooting/querying/`.

Uncertain/no action:

- [#7021](https://github.com/grafana/tempo/pull/7021) adds internal trace events and error status for live-store block creation. I did not classify it as a documentation gap because no stable user-facing config, API, or documented metric surface was identified.
