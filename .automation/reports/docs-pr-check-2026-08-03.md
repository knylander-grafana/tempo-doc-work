# Docs PR check: Tempo merged PRs from 2026-07-27 to 2026-08-03

Source repo: `grafana/tempo`

Query used:

```bash
gh pr list --repo grafana/tempo --state merged --search "merged:>2026-07-27" --json number,title,labels,files,author,mergedAt --limit 100
```

Scope: human-authored PRs merged from `2026-07-27T13:00:08Z` through `2026-08-03T13:00:08Z`, the seven days before this automation run. Bot-authored PRs were filtered out per the docs-pr-check skill.

Docs roots: `docs/sources/tempo/` and `docs/sources/helm-charts/tempo-distributed/`

Generated: 2026-08-03

Bot-authored PRs excluded from classification: #7688 (backport bot), #7684, #7655, #7647 (Renovate).

## Classification table

| PR | Title | Classification | Notes |
|----|-------|----------------|-------|
| [#7687](https://github.com/grafana/tempo/pull/7687) | docs: Add docs for PR 7671 | Docs present | Documents memcached connection tuning from [#7671](https://github.com/grafana/tempo/pull/7671) in `docs/sources/tempo/configuration/_index.md` (`connect_timeout`, `max_idle_conns`, `min_idle_conns_headroom_percentage`) and `docs/sources/tempo/set-up-for-tracing/setup-tempo/upgrade.md`. |
| [#7663](https://github.com/grafana/tempo/pull/7663) | [backendscheduler]: TraceQL query selector for redaction | Docs update needed | Adds `--query` (TraceQL selector) and `--dry-run` to `tempo-cli redact`, mutually exclusive with `--trace-id`. `docs/sources/tempo/operations/tempo_cli.md` still documents `--trace-id` as required and omits the new flags. |
| [#7642](https://github.com/grafana/tempo/pull/7642) | docs: restructure introduction architecture page | Docs present | Documentation-only PR reorganizing `docs/sources/tempo/introduction/_index.md`, `architecture.md`, and `telemetry.md`. |
| [#7608](https://github.com/grafana/tempo/pull/7608) | chore(parquetquery): apply metrics sampling via a SyncIterator option | No docs required | Internal iterator refactor in `pkg/parquetquery/` and vParquet4/vParquet5 block search code; no user-visible API or configuration change. |
| [#7483](https://github.com/grafana/tempo/pull/7483) | feat: Support TraceQL filter in TraceByID V2 endpoint | Docs update needed | API parameters `q` and `keep_hierarchy` documented in `docs/sources/tempo/api_docs/_index.md`. `tempo-cli query api trace-id` gained matching `--q` and `--keep-hierarchy` flags but `docs/sources/tempo/operations/tempo_cli.md` Trace ID section does not document them. |

## Gap summary

Prioritized by user impact:

1. **Docs entirely missing** — None. All user-facing changes in this window have at least partial documentation.

2. **Existing pages that need updates**

   - **#7663 — Redaction TraceQL query selector**
     - `docs/sources/tempo/operations/tempo_cli.md` (`# Redact traces`): Update syntax to show `--query` as an alternative to `--trace-id` (mutually exclusive); document `--dry-run`; add examples for query-based redaction and dry-run mode.
     - Optional follow-up: `docs/sources/tempo/release-notes/v3-0.md` (Trace redaction section) and `docs/sources/tempo/reference-tempo-architecture/components/compaction.md` could mention query-driven redaction jobs for operators who discover the feature outside the CLI reference.

   - **#7483 — TraceByID V2 TraceQL filter**
     - `docs/sources/tempo/operations/tempo_cli.md` (Trace ID under Query API): Add `--q` and `--keep-hierarchy` options with examples; note v2-only behavior and that `--keep-hierarchy` requires `--q`.

   - **#7687 — Memcached connection defaults (optional completeness)**
     - `docs/sources/tempo/operations/caching.md`: Operational guidance on memcached connection limits could mention the new default `max_idle_conns` (`100`) and `min_idle_conns_headroom_percentage` (`-1`) so operators sizing pools see the change outside the configuration reference and upgrade guide.

3. **Needs engineering input** — None identified in this window.

## Screenshot inventory

No product UI changes in classified PRs. Documentation-only restructuring in #7642 does not trigger screenshot validation under the docs-pr-check skill (no frontend `.tsx`/`.ts` changes).

---

Downstream: Run `docs-pr-write` for #7663 and #7483 to update `tempo_cli.md`.
