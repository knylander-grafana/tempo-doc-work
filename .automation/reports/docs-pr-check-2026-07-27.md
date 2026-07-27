# Docs PR check: Tempo merged PRs from 2026-07-20 to 2026-07-27

Source repo: `grafana/tempo`

Query used:

```bash
gh pr list --repo grafana/tempo --state merged --search "merged:>2026-07-20" --json number,title,labels,files,author,mergedAt --limit 100
```

Scope: human-authored PRs merged from `2026-07-20T13:00:24Z` through `2026-07-27T13:00:24Z`, the seven days before this automation run. Bot-authored PRs were filtered out per the docs-pr-check skill.

Docs roots: `docs/sources/tempo/` and `docs/sources/helm-charts/tempo-distributed/`

Generated: 2026-07-27

Bot-authored PRs excluded from classification: #7678, #7673, #7638, #7607.

## Classification table

| PR | Title | Classification | Notes |
|----|-------|----------------|-------|
| [#7679](https://github.com/grafana/tempo/pull/7679) | Add EOL note for 2.9 (release-v2.9) | Docs present | Adds the Tempo 2.9 EOL admonition to `docs/sources/tempo/release-notes/v2-9.md` and `docs/sources/tempo/_index.md` on the 2.9 release branch. |
| [#7677](https://github.com/grafana/tempo/pull/7677) | docs: Add EOL note for 2.9 (2.10 and 2.9) | Docs present | Same EOL note backported to the 2.10 and 2.9 release branches. |
| [#7676](https://github.com/grafana/tempo/pull/7676) | docs: Add EOL note for 2.9 (main and 3.0) | Docs present | Same EOL note on `main` and the 3.0 release branch at `docs/sources/tempo/release-notes/version-2/v2-9.md`. |
| [#7675](https://github.com/grafana/tempo/pull/7675) | docs: Add 2.9.4 rel notes to Version 2.9 (release-v2.9) | Docs present | Release notes for v2.9.4 added to `docs/sources/tempo/release-notes/v2-9.md`. |
| [#7487](https://github.com/grafana/tempo/pull/7487) | docs: Add 2.9.4 rel notes to Version 2.9 (release-v2.10) | Docs present | Same v2.9.4 release notes backported to the 2.10 release branch. |
| [#7486](https://github.com/grafana/tempo/pull/7486) | docs: Draft 2.9.4 release notes for main and 3.0 branches | Docs present | v2.9.4 release notes added to `docs/sources/tempo/release-notes/version-2/v2-9.md` on `main` and 3.0. |
| [#7671](https://github.com/grafana/tempo/pull/7671) | Keep memcached connection pools warm and add connection tuning options | Docs update needed | Adds `connect_timeout` and `min_idle_conns_headroom_percentage`, raises the default `max_idle_conns` from 16 to 100, and stops closing idle connections by default. `docs/sources/tempo/configuration/_index.md` still lists `max_idle_conns` default 16 and omits the new fields; `docs/sources/tempo/operations/caching.md` discusses connection-limit troubleshooting but not the new tuning knobs. |
| [#7665](https://github.com/grafana/tempo/pull/7665) | ingest: fix stale per-partition lag metrics after partition handoff | No docs required | Correctness fix for `tempo_ingest_group_partition_lag` gauges left behind after partition reassignment; no configuration or API contract changed. |
| [#7664](https://github.com/grafana/tempo/pull/7664) | frontend/v1: fix requestBatch reuse race with the doneChan watcher | No docs required | Internal concurrency bug fix in the query-frontend v1 batching path; no user-visible behavior change. |
| [#7662](https://github.com/grafana/tempo/pull/7662) | fix(overrides): fall back to cluster default for unset retry_info_enabled | No docs required | Bug fix so per-tenant overrides that omit `retry_info_enabled` inherit the cluster default instead of silently forcing `false`; restores documented override semantics in `docs/sources/tempo/configuration/_index.md`. |
| [#7660](https://github.com/grafana/tempo/pull/7660) | Fix intrinsic filtering in tag-value autocomplete; simplify lenient parser | No docs required | Restores expected tag-value autocomplete filtering when the incomplete matcher targets an intrinsic; no API or configuration change. |
| [#7628](https://github.com/grafana/tempo/pull/7628) | add span prune by default and check if span prune is already done on the write path | Docs update needed | PR documents `span_pruning_enabled_by_default` in `docs/sources/tempo/configuration/_index.md` and `docs/sources/tempo/api_docs/_index.md`, but the trace-by-ID v2 response `message` field for write-path vs read-path pruning status is undocumented. |
| [#7617](https://github.com/grafana/tempo/pull/7617) | Tag values: cache empty per-block results so tag-less blocks aren't re-scanned | No docs required | Performance optimization that caches empty per-block tag-value results; no API or configuration change. |
| [#7611](https://github.com/grafana/tempo/pull/7611) | metrics-generator: skip stale backlog on startup | Docs present | Opt-in `skip_stale_backlog_on_startup` is documented in `docs/sources/tempo/configuration/_index.md` and `docs/sources/tempo/configuration/manifest.md`. |
| [#7669](https://github.com/grafana/tempo/pull/7669) | chore: prepare v2.9.4 release | No docs required | Release preparation (changelog collation, image tag bump, jsonnet regen); no user-facing product behavior change. |
| [#7668](https://github.com/grafana/tempo/pull/7668) | Add agent guidance: PR template usage and changelog note style | No docs required | Contributor and agent workflow guidance in `.agents/guidance/precommit.md` and `.chloggen/README.md`; no Tempo product-doc gap. |
| [#7661](https://github.com/grafana/tempo/pull/7661) | Point agent guidance at .chloggen instead of CHANGELOG.md | No docs required | Contributor workflow update in `AGENTS.md` and `.agents/guidance/precommit.md`; no Tempo product-doc gap. |
| [#7659](https://github.com/grafana/tempo/pull/7659) | chore: backport Go 1.26.5 CVE fix to release-v2.9 | No docs required | Security dependency bump on the 2.9 release branch; no user-facing behavior change. |
| [#7641](https://github.com/grafana/tempo/pull/7641) | chore: update Go to 1.26.5 and bump x/net, x/text to fix CVEs | No docs required | Security dependency bump on `main`; no user-facing behavior change. |
| [#7643](https://github.com/grafana/tempo/pull/7643) | chore: bump tools image tag to sync gofumpt/golangci-lint/goyacc | No docs required | CI and developer tooling sync; no Tempo product behavior change. |
| [#7639](https://github.com/grafana/tempo/pull/7639) | fix(mixin): damp TempoUserConfigurableOverridesReloadFailing sensitivity | No docs required | Adds a 15-minute `for` duration to a mixin alert rule; operational tuning with no product configuration or API change. |

## Gap summary

1. Update memcached client configuration reference for PR #7671. In `docs/sources/tempo/configuration/_index.md`, document `connect_timeout` and `min_idle_conns_headroom_percentage`, update the `max_idle_conns` default from 16 to 100, and explain that idle connections are no longer closed by default (set `min_idle_conns_headroom_percentage` to a positive value to re-enable reaping). Consider adding a short tuning note to `docs/sources/tempo/operations/caching.md` under the existing connection-limit troubleshooting section.
2. Document span-pruning response messages for PR #7628. In `docs/sources/tempo/api_docs/_index.md` under Query V2, describe the response `message` field and the two pruning-status strings: traces already pruned before ingestion ("already pruned before it reached Tempo") versus traces pruned while serving the request ("pruned by Tempo while serving this request"), including how to retrieve the unpruned trace.

No remaining documentation gap was uncertain enough to require engineering input before these updates can begin.

## Screenshot inventory

No screenshot inventory. None of the human-authored PRs in scope modify frontend UI files or user-visible UI elements. The documentation gaps do not require updates to existing screenshots.
