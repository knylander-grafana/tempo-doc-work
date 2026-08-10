# Docs PR check: Tempo merged PRs from 2026-08-03 to 2026-08-10

Source repo: `grafana/tempo`

Query used:

```bash
gh pr list --repo grafana/tempo --state merged --search "merged:>2026-08-03" --json number,title,labels,files,author,mergedAt --limit 100
```

Scope: human-authored PRs merged from `2026-08-03T13:03:27Z` through `2026-08-10T13:02:58Z`, the seven days before this automation run. Bot-authored PRs were filtered out per the docs-pr-check skill.

Docs roots: `docs/sources/tempo/` and `docs/sources/helm-charts/tempo-distributed/`

Generated: 2026-08-10

Bot-authored PRs excluded from classification: none in this window.

## Classification table

| PR | Title | Classification | Notes |
|----|-------|----------------|-------|
| [#7710](https://github.com/grafana/tempo/pull/7710) | fix: single binary dashboard | No docs required | Fixes the single-binary docker-compose example: corrects the Service Topology dashboard metric from `traces_service_graph_connection_info` to `traces_service_graph_request_total` and removes the invalid `service-graphs-connection-info` processor from `example/docker-compose/single-binary/tempo.yaml`. Example-only change with no product API or configuration surface. |
| [#7704](https://github.com/grafana/tempo/pull/7704) | docs: update single binary docker compose example | Docs present | Documentation-only PR updating `docs/sources/tempo/docker-example.md` to remove outdated Redpanda/Kafka requirements and refresh the `docker compose ps` output for the single-binary quickstart. |
| [#7703](https://github.com/grafana/tempo/pull/7703) | [backend-scheduler] fix redaction in-flight job-accounting leak + dequeue TOCTOU | No docs required | Internal backend-scheduler bug fix for redaction in-flight counter leaks and a dequeue race. No user-facing configuration, API, or CLI change. |
| [#7700](https://github.com/grafana/tempo/pull/7700) | [backend-scheduler] redaction: dry-run does not block compaction, quiesce, or rescan | Docs update needed | User-visible behavior fix: dry-run redaction (read-only preview) no longer disables tenant compaction/retention, enters quiescence, or arms a rescan. Apply-mode behavior is unchanged. Redaction dry-run is exposed via `tempo-cli redact --dry-run` (added in [#7663](https://github.com/grafana/tempo/pull/7663)) but `docs/sources/tempo/operations/tempo_cli.md` still omits `--dry-run` and does not describe dry-run lifecycle semantics. |
| [#7697](https://github.com/grafana/tempo/pull/7697) | feat(servicegraphs): Add support for `db.system.name` | Docs present | Adds `db.system.name` to default `peer_attributes` and `database_name_attributes` for OpenTelemetry semconv v1.30.0 compatibility. Docs updated in `docs/sources/tempo/configuration/_index.md`, `docs/sources/tempo/configuration/manifest.md`, and `docs/sources/tempo/metrics-from-traces/service_graphs/_index.md`. |
| [#7693](https://github.com/grafana/tempo/pull/7693) | add span_pruning_enabled_by_default per tenant override | Docs present | Adds per-tenant override `span_pruning_enabled` under `overrides` to control the query-frontend's `span_pruning_enabled_by_default` default. Documented in `docs/sources/tempo/configuration/_index.md`; API parameter behavior for `span_pruning` and `span_pruning_enabled_by_default` already covered in `docs/sources/tempo/api_docs/_index.md`. |
| [#7691](https://github.com/grafana/tempo/pull/7691) | enhancement: add configurable Kafka producer compression codec | Docs present | Adds `ingest.kafka.producer_compression` (`none`, `gzip`, `snappy`, `lz4`, `zstd`). Documented in `docs/sources/tempo/configuration/_index.md` and regenerated `docs/sources/tempo/configuration/manifest.md`. |
| [#7593](https://github.com/grafana/tempo/pull/7593) | traces: add optional composed trace diff output | Docs present | Adds `trace-summary-v0-composed` format to the experimental trace diff API and `tempo-cli`. Documented in `docs/sources/tempo/api_docs/_index.md` and `docs/sources/tempo/operations/tempo_cli.md` (format option, 64 KiB patch budget, `patchOmitted` behavior). |
| [#7574](https://github.com/grafana/tempo/pull/7574) | docs: correct dedicated column type/capacity | Docs present | Documentation-only correction in `docs/sources/tempo/configuration/_index.md`: dedicated columns support `int` type and up to 20 string + 5 int attributes per scope (not "10 span + 10 resource string-only"). Aligns config reference with `docs/sources/tempo/operations/dedicated_columns.md` and current vParquet5 capabilities. |

## Gap summary

Prioritized by user impact:

1. **Docs entirely missing** — None. All user-facing changes in this window have at least partial documentation, except the redaction dry-run lifecycle gap noted below (which builds on an existing partial-docs gap from [#7663](https://github.com/grafana/tempo/pull/7663)).

2. **Existing pages that need updates**

   - **#7700 — Redaction dry-run lifecycle**
     - `docs/sources/tempo/operations/tempo_cli.md` (`# Redact traces`): Document `--dry-run` and `--query` flags (still missing from [#7663](https://github.com/grafana/tempo/pull/7663)); clarify that dry-run evaluates match counts without rewriting blocks and does **not** block tenant compaction, retention, or rescan.
     - `docs/sources/tempo/reference-tempo-architecture/components/compaction.md`: Add a short note distinguishing apply-mode redaction (exclusive tenant operation that blocks compaction during quiescence) from dry-run preview mode (non-blocking).
     - Optional follow-up: `docs/sources/tempo/release-notes/v3-0.md` (Trace redaction section) could mention dry-run preview mode for operators discovering the feature outside the CLI reference.

3. **Needs engineering input** — None identified in this window.

## Screenshot inventory

No product UI changes (`.tsx`, `.ts`, `.jsx`, `.js`) in classified PRs. Documentation-only pages with existing screenshots were not affected by layout or label changes that would require refreshed captures.

| PR | Affected doc page | Screenshot count | Image references |
|----|-------------------|------------------|------------------|
| — | — | — | No screenshot refresh needed |
