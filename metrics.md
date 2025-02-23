# Monitoring
<!-- THIS FILE MUST NOT BE EDITED BY HAND -->
<!-- ON NEW METRIC ADDED PLEASE RUN 'make generate-metrics-documentation' -->
<!-- markdownlint-disable MD013 -->

All metrics available for scraping are exposed on the `/metrics` endpoint. The metrics are in the Prometheus exposition format.

```sh
curl https://localhost:7979/metrics
```

## Supported Metrics

> Full metric name is constructed as follows:
> `external_dns_<subsystem>_<name>` or `external_dns.<subsystem>.<name>`.

| Name                             | Metric Type | Subsystem   |  Help                                                 |
|:---------------------------------|:------------|:------------|:------------------------------------------------------|
| last_reconcile_timestamp_seconds | Gauge | controller | Timestamp of last attempted sync with the DNS provider |
| last_sync_timestamp_seconds | Gauge | controller | Timestamp of last successful sync with the DNS provider |
| no_op_runs_total | Counter | controller | Number of reconcile loops ending up with no changes on the DNS provider side. |
| verified_a_records | Gauge | controller | Number of DNS A-records that exists both in source and registry. |
| verified_aaaa_records | Gauge | controller | Number of DNS AAAA-records that exists both in source and registry. |
| cache_apply_changes_calls | Counter | provider | Number of calls to the provider cache ApplyChanges. |
| cache_records_calls | Counter | provider | Number of calls to the provider cache Records list. |
| a_records | Gauge | registry | Number of Registry A records. |
| aaaa_records | Gauge | registry | Number of Registry AAAA records. |
| endpoints_total | Gauge | registry | Number of Endpoints in the registry |
| errors_total | Counter | registry | Number of Registry errors. |
| a_records | Gauge | source | Number of Source A records. |
| aaaa_records | Gauge | source | Number of Source AAAA records. |
| endpoints_total | Gauge | source | Number of Endpoints in all sources |
| errors_total | Counter | source | Number of Source errors. |
| adjustendpoints_errors_total | Gauge | webhook_provider | Errors with AdjustEndpoints method |
| adjustendpoints_requests_total | Gauge | webhook_provider | Requests with AdjustEndpoints method |
| applychanges_errors_total | Gauge | webhook_provider | Errors with ApplyChanges method |
| applychanges_requests_total | Gauge | webhook_provider | Requests with ApplyChanges method |
| records_errors_total | Gauge | webhook_provider | Errors with Records method |
| records_requests_total | Gauge | webhook_provider | Requests with Records method |
