---
aliases:
  - ../../operations/traceql-metrics/ # /docs/tempo/next/operations/traceql-metrics/
title: Configure TraceQL metrics
menuTitle: Configure TraceQL metrics
description: Learn about configuring TraceQL metrics.
weight: 400
keywords:
  - Prometheus
  - TraceQL
  - TraceQL metrics
---

# Configure TraceQL metrics

TraceQL language provides metrics queries as a feature.
Metric queries extend trace queries by applying a function to trace query results.
This powerful feature creates metrics from traces, much in the same way that LogQL metric queries create metrics from logs.

## Before you begin

To use the metrics generated from traces, you need to:

- Configure a Tempo data source in Grafana or Grafana Cloud ([documentation](/docs/grafana/<GRAFANA_VERSION>/datasources/tempo/configure-tempo-data-source/))
- Access Grafana Cloud or Grafana version 10.4 or later

## Evaluate query timeouts

Because of their expensive nature, these queries can take a long time to run.
As such, consider increasing the timeouts in various places of
the system to allow enough time for the data to be returned.

Consider these areas when raising timeouts:

- Any proxy in front of Grafana
- Grafana data source for Prometheus pointing at Tempo
- Tempo configuration
  - `querier.search.query_timeout`
  - `server.http_server_read_timeout`
  - `server.http_server_write_timeout`

## Set TraceQL metrics query options

The `query_frontend.metrics` configuration block controls all TraceQL metrics queries.
The configuration depends on the environment.

{{< admonition type="note" >}}
The default maximum time range for a metrics query is 24 hours, configured using the `query_frontend.metrics.max_duration` parameter.

This is different to the default TraceQL maximum time range of 168 hours (7 days).

{{< /admonition >}}

The `query_frontend.metrics.query_backend_after` parameter controls the boundary between querying the live-store and backend storage.
Time ranges older than `query_backend_after` (default `15m`) are searched in backend/object storage only, while more recent data is queried from the live-store.

### Understand metrics query range limits

Tempo validates TraceQL metrics query ranges before it applies internal step alignment.
This means `query_frontend.metrics.max_duration` is checked against the time range that you request, not against the larger range Tempo might create internally while sharding the query.

The effective metrics query range limit is selected in this order:

1. The read override `max_metrics_duration`, when it is non-zero. You can set it under `overrides.defaults.read` for the default override or under a tenant-specific override.
1. The `query_frontend.metrics.max_duration` value.

The default `query_frontend.metrics.max_duration` is `24h`.
A value of `0` disables the limit.
Tempo applies the same limit to range metrics queries and instant metrics queries.

If `query_frontend.query_end_cutoff` is set, Tempo clamps the query end time to the cutoff before it checks `max_duration`.
This helps requests with an end time in the future, for example, because of client clock skew, pass validation when the effective query range is within the limit.
If the whole requested query window falls inside `query_end_cutoff`, Tempo rejects the request with a bad request error instead of returning a partial or empty result.

For example, use a global 24 hour limit and a 12 hour default read override:

```yaml
query_frontend:
  query_end_cutoff: 30s
  metrics:
    max_duration: 24h

overrides:
  defaults:
    read:
      max_metrics_duration: 12h
```

For example, in a cloud environment, smaller jobs with more concurrency may be
desired due to the nature of scale on the backend.

```yaml
query_frontend:
  metrics:
    concurrent_jobs: 1000
    target_bytes_per_job: 2.25e+08 # ~225MB
    interval: 30m0s
```

For an on-prem backend, you can improve query times by lowering the concurrency,
while increasing the job size.

```yaml
query_frontend:
  metrics:
    concurrent_jobs: 8
    target_bytes_per_job: 1.25e+09 # ~1.25GB
```

## Sampling and performance optimization

TraceQL metrics queries support sampling hints to improve performance on large datasets. Refer to the [TraceQL metrics sampling](/docs/tempo/<TEMPO_VERSION>/metrics-from-traces/metrics-queries/sampling-guide/) documentation for more information.

When using sampling in your TraceQL metrics queries, consider:

- **Timeout settings:** Sampled queries run faster but may still benefit from adequate timeouts
- **Concurrent jobs:** Sampling reduces per-job processing time, allowing higher concurrency
