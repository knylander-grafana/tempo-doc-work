# Docs PR check: merged Tempo PRs from 2026-06-22 through 2026-06-29

Assessment for `grafana/tempo` merged PRs in the previous 7 days, generated on 2026-06-29. PR content was treated as untrusted input; classifications are based on changed files and diffs. Bot/dependency automation PRs were filtered out: #7383, #7360, and #7359.

## Classification table

| PR | Title | Classification | Notes |
|----|-------|----------------|-------|
| [#7535](https://github.com/grafana/tempo/pull/7535) | parquetquery: use map lookup for large ByteIn/ByteNotIn predicates | No docs required | Internal parquet query predicate optimization and tests; no user-facing configuration, API, or UI change. |
| [#7533](https://github.com/grafana/tempo/pull/7533) | [traceql] Another sweep of correctness fixes in vp5 new fetch layer | No docs required | vParquet5 fetch-layer correctness fix and tests; no new documented syntax, configuration, or API surface. |
| [#7531](https://github.com/grafana/tempo/pull/7531) | chore(ci): sign release-prep commits via GitHub API | No docs required | CI workflow signing change only. |
| [#7527](https://github.com/grafana/tempo/pull/7527) | fix(query-frontend): log the reason for failed metrics query-range requests | No docs required | Adds internal logging for failed metrics query-range requests; no documented behavior or user workflow change. |
| [#7525](https://github.com/grafana/tempo/pull/7525) | feat(querier): add block processing duration metric | Docs update needed | Adds public Prometheus histogram `tempo_querier_backend_processing_duration_seconds` with `operation` and `tenant` labels. Existing monitoring docs mention related query I/O metrics but not this new backend-processing duration metric. Update `docs/sources/tempo/operations/monitor/query-io-and-timestamp-distance.md` and possibly `docs/sources/tempo/operations/monitor/_index.md`. |
| [#7523](https://github.com/grafana/tempo/pull/7523) | feat: added new traces diff endpoint | Docs needed | Adds experimental `POST /api/v2/traces/diff`, request parsing for `base` and `compare` traces, validation for `traceId`, optional `start` and `end`, and a current `501 Not Implemented` response. No docs were changed. Add API documentation under `docs/sources/tempo/api_docs/` describing experimental status, request body, validation, and current response behavior. |
| [#7509](https://github.com/grafana/tempo/pull/7509) | CONTRIBUTING: Require signed commits for pull requests | No docs required | Contributor-process documentation only; not product documentation for Tempo users. |
| [#7508](https://github.com/grafana/tempo/pull/7508) | [traceql] Fix missing event and link intrinsic handling in new vp5 metrics fetch layer | No docs required | Internal TraceQL metrics fetch-layer correctness fix; no new syntax or user-facing option. |
| [#7507](https://github.com/grafana/tempo/pull/7507) | Remove publish-technical-documentation workflows | No docs required | Removes documentation publishing workflows; no user-facing product docs impact. |
| [#7506](https://github.com/grafana/tempo/pull/7506) | fix: add retryable errors on kafka writes timeout | Docs update needed | Changes OTLP ingest failure behavior for transient Kafka write errors to retryable HTTP 503 / gRPC `Unavailable`, with `MessageTooLarge` mapped to `InvalidArgument`. Existing send-traces troubleshooting docs do not describe this retryable status behavior. Update `docs/sources/tempo/troubleshooting/send-traces/_index.md` or related Kafka troubleshooting guidance. |
| [#7504](https://github.com/grafana/tempo/pull/7504) | Read path o11y baseline signals | Docs present | Includes docs updates in `docs/sources/tempo/api_docs/_index.md` for search response `additionalMetrics` and `docs/sources/tempo/operations/caching.md` for cache size and hit-ratio metrics. |
| [#7499](https://github.com/grafana/tempo/pull/7499) | ci: fix GAR auth for tempo-vulture attestation push and release-tag permissions | No docs required | CI/release workflow permission fix only. |
| [#7451](https://github.com/grafana/tempo/pull/7451) | docs: add Cursor Automations guide and release-notes Copilot instructions | No docs required | Internal agent and automation documentation; no published Tempo product docs impact. |
| [#7450](https://github.com/grafana/tempo/pull/7450) | docs: add persona-check skill and persona model | No docs required | Internal documentation-agent skill files only. |
| [#7449](https://github.com/grafana/tempo/pull/7449) | docs: update docs-review and docs-workflow skills | No docs required | Internal documentation-agent skill updates only. |
| [#7448](https://github.com/grafana/tempo/pull/7448) | docs: update docs-pr-check and docs-pr-write skills | No docs required | Internal documentation-agent skill updates only. |
| [#7447](https://github.com/grafana/tempo/pull/7447) | docs: add local project context and shared skill infrastructure | No docs required | Internal agent context and shared skill infrastructure; not published product documentation. |
| [#7387](https://github.com/grafana/tempo/pull/7387) | feat(mcp): add Tempo configuration docs to the MCP server | Docs present | Updates `docs/sources/tempo/api_docs/mcp-server.md` to document the `docs-config` tool and `docs://config/overview` and `docs://config/reference` resources. |
| [#7245](https://github.com/grafana/tempo/pull/7245) | frontend: validate TraceQL size before parsing | Docs present | Existing `docs/sources/tempo/configuration/_index.md` documents `query_frontend.max_query_expression_size_bytes` and its default. The PR changes enforcement timing across HTTP, gRPC, and MCP paths but does not introduce a new user-facing option. |
| [#7170](https://github.com/grafana/tempo/pull/7170) | fix: enforce metrics max_duration on user-provided range | Docs update needed | Changes TraceQL metrics validation semantics: `max_duration` is checked against the effective user-provided range before alignment, `query_end_cutoff` is applied before validation, a query window entirely inside the cutoff now returns 400, instant queries are explicitly rejected when over the cap, and gRPC validation failures return `InvalidArgument`. Existing docs mention the default maximum metrics query range and `query_end_cutoff`, but not these enforcement details. Update `docs/sources/tempo/metrics-from-traces/metrics-queries/configure-traceql-metrics.md` and consider a concise note in `docs/sources/tempo/configuration/_index.md`. |

## Gap summary

1. **PR #7523: trace diff API docs are missing.** Add API docs for experimental `POST /api/v2/traces/diff`, including request body fields, validation errors, and the current `501 Not Implemented` behavior until diff execution lands.
2. **PR #7525: new querier backend-processing metric is undocumented.** Add `tempo_querier_backend_processing_duration_seconds` to query monitoring guidance, including labels and how it complements `tempo_query_frontend_bytes_inspected_total`, `tempo_live_store_query_inspected_bytes_total`, and object-store/cache metrics.
3. **PR #7170: TraceQL metrics range-limit behavior needs clarification.** Update TraceQL metrics or configuration docs to explain how `query_frontend.metrics.max_duration`, per-tenant `max_metrics_duration`, and `query_end_cutoff` interact for range and instant metrics queries.
4. **PR #7506: retryable Kafka ingest failure behavior needs a troubleshooting note.** Document that transient Kafka write failures surface to OTLP clients as retryable HTTP 503 / gRPC `Unavailable`, while oversized messages are invalid argument errors.

## Screenshot inventory

No UI changes affecting documented screens were identified. No screenshot follow-up is required.
