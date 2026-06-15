# Docs PR check: merged Tempo PRs, 2026-06-08 to 2026-06-15

Evaluation window: 2026-06-08T13:00:31Z through 2026-06-15T13:00:31Z.

Source query: `gh pr list --repo grafana/tempo --state merged --search "merged:>=2026-06-08" --limit 200 --json number,title,mergedAt,author,labels,url`.

Scope: PRs merged in `grafana/tempo` during the previous seven days. Classifications follow `.claude/skills/docs-pr-check/SKILL.md`.

## Table

| PR | Title | Classification | Notes |
|----|-------|----------------|-------|
| [#7376](https://github.com/grafana/tempo/pull/7376) | jsonnet: use AverageValue for the live-store autoscaler trigger | Docs present | User-facing KEDA autoscaling is already described in `docs/sources/tempo/set-up-for-tracing/setup-tempo/deploy/kubernetes/tanka.md`; this fixes generated trigger behavior without adding a new option. |
| [#7421](https://github.com/grafana/tempo/pull/7421) | [release-v2.9] [chore] upgrade tempo to go version 1.26.3 | Docs update needed | `docs/sources/tempo/release-notes/version-2/v2-9.md` has no v2.9.3/Go 1.26.3 release-note entry in this checkout. |
| [#7422](https://github.com/grafana/tempo/pull/7422) | [release-v2.10] [chore] upgrade tempo to go version 1.26.3 | Docs present | Covered by `docs/sources/tempo/release-notes/version-2/v2-10.md` security fixes: "Built Tempo with Go 1.26.3". |
| [#7423](https://github.com/grafana/tempo/pull/7423) | [release-v3.0] [chore] upgrade tempo to go version 1.26.3 | Docs present | Covered by `docs/sources/tempo/release-notes/v3-0.md` v3.0.2 security fixes. |
| [#7425](https://github.com/grafana/tempo/pull/7425) | chore: prepare v2.10.6 release | Docs present | PR updated docs/example image tags; release-note coverage added by #7442 and backfilled by #7441. |
| [#7426](https://github.com/grafana/tempo/pull/7426) | chore: prepare v2.9.3 release | Docs update needed | PR body says v2.9.3 includes user-facing release items, but `docs/sources/tempo/release-notes/version-2/v2-9.md` has no v2.9.3 section in this checkout. |
| [#7411](https://github.com/grafana/tempo/pull/7411) | chore: prepare v3.0.1 release | Docs present | Release prep updated docs/example image tags; no separate v3.0.1 release-note content identified as required. |
| [#7432](https://github.com/grafana/tempo/pull/7432) | chore: stop publishing 32-bit ARM binary archives | Docs update needed | Generic install docs now say 32-bit ARM binaries are not published, but v2.9 release notes do not appear to mention the v2.9.3 artifact change. |
| [#7431](https://github.com/grafana/tempo/pull/7431) | chore: stop publishing 32-bit ARM binary archives | Docs present | Covered by Linux install docs and v2.10.7 release-note backfill (#7479/#7480). |
| [#7429](https://github.com/grafana/tempo/pull/7429) | chore: publish release images to GAR via OIDC | Docs update needed | Versioned Docker Hub tags remain mirrored, but the PR also notes `:latest` is intentionally not pushed. Existing docs/release notes do not describe this image tag policy. |
| [#7428](https://github.com/grafana/tempo/pull/7428) | chore: stop pushing the :latest image tag on release | Docs update needed | User-visible release artifact behavior; no existing docs found for the mutable `:latest` image tag being discontinued. |
| [#7430](https://github.com/grafana/tempo/pull/7430) | chore: publish release images to GAR via OIDC | Docs update needed | Same image publication policy as #7429, plus v2.9 release-note coverage is absent in this checkout. |
| [#7409](https://github.com/grafana/tempo/pull/7409) | Docs: scalars in math | Docs present | PR updates `docs/sources/tempo/metrics-from-traces/metrics-queries/_index.md` and `functions.md`. |
| [#7434](https://github.com/grafana/tempo/pull/7434) | chore: add @knylander-grafana to codeowners for .chloggen | No docs required | CODEOWNERS-only maintenance. |
| [#7288](https://github.com/grafana/tempo/pull/7288) | distributor: bound rebatching preallocation | No docs required | Internal allocation/performance fix with no documented option or user workflow change. |
| [#7246](https://github.com/grafana/tempo/pull/7246) | chore(deps) upgrade grpc-ecosystem/go-grpc-middleware from v1 to v2 | No docs required | Dependency/internal middleware upgrade. |
| [#7296](https://github.com/grafana/tempo/pull/7296) | chore(deps) replace go-test/deep with google/go-cmp | No docs required | Test dependency replacement. |
| [#7427](https://github.com/grafana/tempo/pull/7427) | chore(renovate): group go.mod patch updates into a single PR | No docs required | Renovate configuration maintenance. |
| [#7438](https://github.com/grafana/tempo/pull/7438) | chore: prepare v3.0.2 release | Docs present | Release-note coverage added by #7443 and backfilled by #7441. |
| [#7397](https://github.com/grafana/tempo/pull/7397) | chore(deps): update grafana/tempo-ci-tools docker tag to v20260603 (release-v2.10) | No docs required | CI/tooling dependency update. |
| [#7441](https://github.com/grafana/tempo/pull/7441) | docs: backfill v2.10.5/v2.10.6/v3.0.2 changelog and release notes to main | Docs present | Docs-only release-note backfill. |
| [#7440](https://github.com/grafana/tempo/pull/7440) | docs: fix Go version typo in v2.10.5 changelog | Docs present | Docs-only correction. |
| [#7442](https://github.com/grafana/tempo/pull/7442) | docs: add v2.10.6 release notes and backfill 2.10.1/2.10.5 Go versions | Docs present | Docs-only release-note update. |
| [#7443](https://github.com/grafana/tempo/pull/7443) | docs: add v3.0.2 release notes | Docs present | Docs-only release-note update. |
| [#7379](https://github.com/grafana/tempo/pull/7379) | fix(deps): update module github.com/jaegertracing/jaeger-idl to v0.9.0 (main) | No docs required | Dependency update. |
| [#7380](https://github.com/grafana/tempo/pull/7380) | fix(deps): update module github.com/minio/minio-go/v7 to v7.2.0 (main) | No docs required | Dependency update. |
| [#7361](https://github.com/grafana/tempo/pull/7361) | chore(deps): update otel/opentelemetry-collector docker tag to v0.153.0 (main) | No docs required | Example/test dependency image update. |
| [#7355](https://github.com/grafana/tempo/pull/7355) | chore(deps): update grafana/shared-workflows/create-github-app-token action to v0.3.1 (main) | No docs required | CI dependency update. |
| [#7354](https://github.com/grafana/tempo/pull/7354) | chore(deps): update actions/setup-go action to v6.4.0 (main) | No docs required | CI dependency update. |
| [#7453](https://github.com/grafana/tempo/pull/7453) | Upgrade goreleaser to v2.16.0 | No docs required | Release tooling dependency update. |
| [#7358](https://github.com/grafana/tempo/pull/7358) | [backendscheduler] redaction: honor batch barrier in retention; surface missing redaction blocks | Docs update needed | Adds operator-facing metric `tempo_backend_worker_redaction_block_missing_total` and retention/redaction behavior. `docs/sources/tempo/reference-tempo-architecture/components/compaction.md` key metrics table does not include the new metric. |
| [#7388](https://github.com/grafana/tempo/pull/7388) | fix(deps): update opentelemetry-contrib (main) | No docs required | Dependency update. |
| [#7357](https://github.com/grafana/tempo/pull/7357) | chore(deps): update module github.com/golangci/golangci-lint/v2 to v2.12.2 (main) | No docs required | Lint tooling dependency update. |
| [#7419](https://github.com/grafana/tempo/pull/7419) | chore: add changelog entry skill | No docs required | Internal agent/tooling documentation workflow. |
| [#7456](https://github.com/grafana/tempo/pull/7456) | chore(deps): update module github.com/mark3labs/mcp-go to v0.54.1 (main) | No docs required | Dependency update only. |
| [#7439](https://github.com/grafana/tempo/pull/7439) | chore(ci): automate Tempo releases (RC, minor, patch) | Docs present | Maintainer-facing release automation is documented in `RELEASES.MD`; no product docs required. |
| [#7457](https://github.com/grafana/tempo/pull/7457) | [DOC] Add line for 32-bit ARM binary. | Docs present | Updates `docs/sources/tempo/set-up-for-tracing/setup-tempo/deploy/locally/linux.md`. |
| [#7471](https://github.com/grafana/tempo/pull/7471) | docs: Add blog post and video for 3.0 release | Docs present | Updates `docs/sources/tempo/release-notes/v3-0.md`. |
| [#7470](https://github.com/grafana/tempo/pull/7470) | docs: Update blog links for instrumentation page | Docs present | Updates `docs/sources/tempo/set-up-for-tracing/instrument-send/set-up-instrumentation/_index.md`. |
| [#7473](https://github.com/grafana/tempo/pull/7473) | [release-v3.0] docs: Update blog links for instrumentation page | Docs present | Backport of docs update #7470. |
| [#7472](https://github.com/grafana/tempo/pull/7472) | [release-v3.0] docs: Add blog post and video for 3.0 release | Docs present | Backport of docs update #7471. |
| [#7446](https://github.com/grafana/tempo/pull/7446) | docs: relocate shared doc-agent resources to .claude/skills/shared/ | No docs required | Internal doc-agent resource relocation, not a Tempo product documentation change. |
| [#7467](https://github.com/grafana/tempo/pull/7467) | [Bugfix] limit compare() | Docs update needed | `compare()` now rejects `topN` values outside 1-1000. `docs/sources/tempo/metrics-from-traces/metrics-queries/functions.md` documents `topN` but not the new maximum. |
| [#7463](https://github.com/grafana/tempo/pull/7463) | util: stack-buffer fast path for TraceIDToHexString | No docs required | Internal performance optimization. |
| [#7469](https://github.com/grafana/tempo/pull/7469) | fix: incorrect build version reported by tempo binaries | Docs present | PR updates `RELEASES.MD`; v2.10.7 release-note coverage added by #7479/#7480. |
| [#7474](https://github.com/grafana/tempo/pull/7474) | chore(ci): backport release automation (#7439) and build version fix (#7469) to release-v2.10 | Docs present | Backports maintainer release docs/build-version fix; v2.10.7 release notes cover the user-visible version-reporting fix. |
| [#7475](https://github.com/grafana/tempo/pull/7475) | chore: prepare release v2.10.7 | Docs present | Release-note coverage added by #7480 and backfilled to main by #7479. |
| [#7477](https://github.com/grafana/tempo/pull/7477) | chore(ci): backport release automation (#7439) and build version fix (#7469) to release-v3.0 | Docs present | Backports maintainer release docs/build-version fix; no additional product docs identified. |
| [#7476](https://github.com/grafana/tempo/pull/7476) | chore: use large runner for release workflow | No docs required | CI/release infrastructure change. |
| [#7478](https://github.com/grafana/tempo/pull/7478) | [release-v3.0] chore: use large runner for release workflow | No docs required | Backport of CI/release infrastructure change. |
| [#7479](https://github.com/grafana/tempo/pull/7479) | docs: backfill v2.10.7 changelog and release notes to main | Docs present | Docs-only changelog/release-note backfill. |
| [#7480](https://github.com/grafana/tempo/pull/7480) | docs: add v2.10.7 release notes | Docs present | Docs-only release-note update. |

## Gap summary

Prioritized documentation work, ordered by user impact:

1. **Document changed release artifact/image publication behavior.**
   - PRs: #7428, #7429, #7430.
   - Gap: Releases no longer push the mutable `grafana/tempo:latest` image tag; versioned tags continue. Existing docs do not appear to state this policy.
   - Suggested location: release notes for affected supported versions and, if there is a canonical image-installation page, that page as a general policy note.

2. **Backfill v2.9.3 release notes.**
   - PRs: #7421, #7426, #7430, #7432.
   - Gap: `docs/sources/tempo/release-notes/version-2/v2-9.md` has no v2.9.3 section in this checkout. The v2.9.3 release prep says the changelog included Go 1.26.3, busybox removal, security dependency updates, OpenCensus receiver removal, and related release changes.
   - Suggested location: `docs/sources/tempo/release-notes/version-2/v2-9.md` and the corresponding release-branch source, if maintained separately.

3. **Add the new redaction missing-block metric to compaction/backend-worker docs.**
   - PR: #7358.
   - Gap: `tempo_backend_worker_redaction_block_missing_total` is a new operator-facing metric and warning signal, but `docs/sources/tempo/reference-tempo-architecture/components/compaction.md` does not list it in "Key metrics".
   - Suggested location: `docs/sources/tempo/reference-tempo-architecture/components/compaction.md`; optionally also mention that retention now respects the redaction batch barrier while a redaction batch awaits rescan.

4. **Document the `compare()` `topN` cap.**
   - PR: #7467.
   - Gap: `docs/sources/tempo/metrics-from-traces/metrics-queries/functions.md` says the second `compare()` parameter is optional `topN` and defaults to 10, but does not state the accepted range is 1-1000.
   - Suggested location: `docs/sources/tempo/metrics-from-traces/metrics-queries/functions.md`, in the `compare` parameters section.

5. **Low-risk watch item: live-store KEDA trigger behavior.**
   - PR: #7376.
   - Assessment: Existing Tanka docs cover enabling live-store KEDA autoscaling and its user-facing knobs. The PR changes generated KEDA `metricType` behavior, not documented configuration. No immediate docs update required unless operators need an upgrade note for unexpected live-store scaling behavior before this fix.

## Existing docs coverage found

- `docs/sources/tempo/set-up-for-tracing/setup-tempo/deploy/kubernetes/tanka.md` covers KEDA autoscaling knobs for live-store/block-builder.
- `docs/sources/tempo/set-up-for-tracing/setup-tempo/deploy/locally/linux.md` states Tempo publishes AMD64 and arm64 binaries and does not publish 32-bit ARM binaries.
- `docs/sources/tempo/release-notes/version-2/v2-10.md` covers v2.10.6/v2.10.7 release-note items including Go 1.26.3, OpenCensus removal, build-version reporting, and 32-bit ARM binary removal.
- `docs/sources/tempo/release-notes/v3-0.md` covers v3.0.2 Go 1.26.3 and v3.0 platform notes.
- `docs/sources/tempo/metrics-from-traces/metrics-queries/functions.md` covers scalar math and the existing `compare()` function semantics, but needs the new `topN` cap.
- `docs/sources/tempo/reference-tempo-architecture/components/compaction.md` covers backend scheduler/worker jobs and metrics, but needs the new redaction missing-block metric.
