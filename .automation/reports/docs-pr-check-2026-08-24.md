# Docs PR check: Tempo merged PRs from 2026-08-17 to 2026-08-24

Source repo: `grafana/tempo`

Query used:

```bash
gh pr list --repo grafana/tempo --state merged --search "merged:>2026-08-17" --json number,title,labels,files,author,mergedAt --limit 100
```

Scope: human-authored PRs merged from `2026-08-18T12:55:41Z` through `2026-08-21T17:27:46Z`, the seven days before this automation run. Bot-authored PRs were filtered out per the docs-pr-check skill.

Docs roots: `docs/sources/tempo/` and `docs/sources/helm-charts/tempo-distributed/`

Generated: 2026-08-24

Bot-authored PRs excluded from classification: #7793, #7792, #7791 (Renovate etcd security updates), #7770, #7769, #7768, #7767 (Renovate dependency bumps).

## Classification table

| PR | Title | Classification | Notes |
|----|-------|----------------|-------|
| [#7803](https://github.com/grafana/tempo/pull/7803) | [blockselector]: cache DedicatedColumnsHash per entry | No docs required | Internal compaction block-selector optimization: precomputes `DedicatedColumnsHash()` per entry instead of recomputing it in the `BlocksToCompact` comparison loop. No configuration, API, or CLI change. |
| [#7795](https://github.com/grafana/tempo/pull/7795) | [tempo-mixin] gate Backend Work histogram panels on a latency_metrics variable | No docs required | Updates compiled tempo-mixin Backend Work dashboard JSON to gate histogram panels behind a `latency_metrics` variable (classic vs native). Operator dashboard tooling only; no product configuration or API surface. |
| [#7787](https://github.com/grafana/tempo/pull/7787) | fix: serialize DRAIN pruning with training | No docs required | Internal metrics-generator bug fix: holds the DRAIN sanitizer mutex during periodic maintenance so pruning cannot race with span-name training. No user-visible behavior or configuration change. |
| [#7786](https://github.com/grafana/tempo/pull/7786) | Preserve no-compact flag when copying vParquet5 blocks | No docs required | Internal storage bug fix: `CopyBlock` now propagates the no-compact flag from source to destination blocks so block-builder copies are not compacted or polled prematurely. No operator-facing configuration or API change. |
| [#7785](https://github.com/grafana/tempo/pull/7785) | feat(mcp): add trace diff tool | Docs present | Adds experimental `trace-diff` MCP tool and clarifies composed diff response semantics. Documented in `docs/sources/tempo/api_docs/mcp-server.md`, `docs/sources/tempo/api_docs/_index.md` (trace diff section), and `docs/sources/tempo/introduction/tempo-and-ai.md` (tool count and compare-traces wording). |
| [#7777](https://github.com/grafana/tempo/pull/7777) | Updated release-notes-workflow for changelog process | No docs required | Updates agent skill files under `.claude/skills/` (docs-pr-check, docs-pr-write, release-notes-workflow, release-notes-placement). Internal documentation-agent workflow only; no user-facing product docs. |
| [#7775](https://github.com/grafana/tempo/pull/7775) | Make vParquet5 the default block format | Docs present | Changes default block format from vParquet4 to vParquet5 in code and docs. Updated `docs/sources/tempo/configuration/_index.md`, `docs/sources/tempo/configuration/manifest.md`, `docs/sources/tempo/configuration/parquet.md`, `docs/sources/tempo/operations/schema.md`, and `docs/sources/tempo/reference-tempo-architecture/block-format.md`. |
| [#7773](https://github.com/grafana/tempo/pull/7773) | Add block-size observability at flush and compaction | No docs required | Adds Prometheus histogram metrics `tempo_block_builder_flush_size_bytes` and `tempodb_compaction_output_block_size_bytes`. Observability-only change with no configuration flag or API surface. |
| [#7772](https://github.com/grafana/tempo/pull/7772) | [backend-scheduler] expose pending job depth | Docs present | Adds `tempo_backend_scheduler_jobs_pending` metric and updates tempo-mixin Backend Work dashboard panels. Product doc updated in `docs/sources/tempo/reference-tempo-architecture/components/compaction.md` with metric description and autoscaling context. |
| [#7708](https://github.com/grafana/tempo/pull/7708) | TraceByID V2: add match_depth and ancestor_depth query params | Docs update needed | User-facing TraceByID V2 filtering enhancement. HTTP parameters documented in `docs/sources/tempo/api_docs/_index.md`, and `tempo-cli query api trace-id` gained `--match-depth`, `--ancestor-depth`, `--keep-hierarchy`, and `--q` flags in code, but `docs/sources/tempo/operations/tempo_cli.md` still lists only `--v1`, `--org-id`, and `--header` for that command. |
| [#7702](https://github.com/grafana/tempo/pull/7702) | [backend-scheduler] redaction: caller-specified [start,end] time window | Docs present | Adds `--start`/`--end` time-window scoping (and related `--query`/`--dry-run` clarifications) to `tempo-cli redact`. Documented in `docs/sources/tempo/operations/tempo_cli.md` with examples, overlap semantics, batching guidance, and a worker-version compatibility warning. |
| [#7689](https://github.com/grafana/tempo/pull/7689) | Track TraceQL engine bytes on the query read path | Docs update needed | Adds experimental `engine_bytes_tracking` override under `overrides.defaults.read` (disabled unless enabled). Override documented in `docs/sources/tempo/configuration/_index.md`, but `docs/sources/tempo/api_docs/_index.md` still lists only legacy `additionalMetrics` keys (`cacheHits`, `cacheMisses`, `cacheBytes`, etc.) and does not document the new `engineBytes` key returned on search, metrics, and trace-by-id responses when tracking is enabled. |

## Gap summary

Prioritized by user impact:

1. **Docs entirely missing** — None. All user-facing changes in this window have at least partial documentation.

2. **Existing pages that need updates**

   - **#7708 — TraceByID V2 depth controls in tempo-cli reference**
     - `docs/sources/tempo/operations/tempo_cli.md` (`### Trace ID` under `query api trace-id`): Document `--q`, `--keep-hierarchy`, `--match-depth`, and `--ancestor-depth` with the same semantics as the HTTP API (defaults, ignored-without-`q` behavior, and `ancestor_depth` requiring `keep_hierarchy`).
     - Optional cross-link from the CLI section to the TraceByID V2 parameter list in `docs/sources/tempo/api_docs/_index.md`.

   - **#7689 — Engine bytes additional metric**
     - `docs/sources/tempo/api_docs/_index.md` (`metrics.additionalMetrics`): Add `engineBytes` to the documented key list; note it appears on search, TraceQL metrics, and trace-by-id responses when `overrides.defaults.read.engine_bytes_tracking` (or a per-tenant override) is enabled.
     - `docs/sources/tempo/configuration/_index.md`: Clarify the default is disabled when unset (the current text references a "cluster-wide default" without stating the default value).

3. **Needs engineering input** — None identified in this window.

## Screenshot inventory

No product UI changes (`.tsx`, `.ts`, `.jsx`, `.js`) in classified PRs. PR #7785 updated the inline SVG diagram on `docs/sources/tempo/introduction/tempo-and-ai.md` as part of the same PR; no separate screenshot refresh is required.

| PR | Affected doc page | Screenshot count | Image references |
|----|-------------------|------------------|------------------|
| — | — | — | No screenshot refresh needed |
