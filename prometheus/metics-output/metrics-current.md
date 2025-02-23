

```markdown
# HELP external_dns_controller_last_reconcile_timestamp_seconds Timestamp of last attempted sync with the DNS provider
# TYPE external_dns_controller_last_reconcile_timestamp_seconds gauge
external_dns_controller_last_reconcile_timestamp_seconds 1.740246918175597e+09
# HELP external_dns_controller_last_sync_timestamp_seconds Timestamp of last successful sync with the DNS provider
# TYPE external_dns_controller_last_sync_timestamp_seconds gauge
external_dns_controller_last_sync_timestamp_seconds 1.740246932026549e+09
# HELP external_dns_controller_no_op_runs_total Number of reconcile loops ending up with no changes on the DNS provider side.
# TYPE external_dns_controller_no_op_runs_total counter
external_dns_controller_no_op_runs_total 0
# HELP external_dns_controller_verified_a_records Number of DNS A-records that exists both in source and registry.
# TYPE external_dns_controller_verified_a_records gauge
external_dns_controller_verified_a_records 0
# HELP external_dns_controller_verified_aaaa_records Number of DNS AAAA-records that exists both in source and registry.
# TYPE external_dns_controller_verified_aaaa_records gauge
external_dns_controller_verified_aaaa_records 0
# HELP external_dns_registry_a_records Number of Registry A records.
# TYPE external_dns_registry_a_records gauge
external_dns_registry_a_records 10
# HELP external_dns_registry_aaaa_records Number of Registry AAAA records.
# TYPE external_dns_registry_aaaa_records gauge
external_dns_registry_aaaa_records 0
# HELP external_dns_registry_endpoints_total Number of Endpoints in the registry
# TYPE external_dns_registry_endpoints_total gauge
external_dns_registry_endpoints_total 11
# HELP external_dns_registry_errors_total Number of Registry errors.
# TYPE external_dns_registry_errors_total counter
external_dns_registry_errors_total 0
# HELP external_dns_source_a_records Number of Source A records.
# TYPE external_dns_source_a_records gauge
external_dns_source_a_records 10
# HELP external_dns_source_aaaa_records Number of Source AAAA records.
# TYPE external_dns_source_aaaa_records gauge
external_dns_source_aaaa_records 0
# HELP external_dns_source_endpoints_total Number of Endpoints in all sources
# TYPE external_dns_source_endpoints_total gauge
external_dns_source_endpoints_total 10
# HELP external_dns_source_errors_total Number of Source errors.
# TYPE external_dns_source_errors_total counter
external_dns_source_errors_total 0
# HELP external_dns_webhook_provider_adjustendpoints_errors_total Errors with AdjustEndpoints method
# TYPE external_dns_webhook_provider_adjustendpoints_errors_total gauge
external_dns_webhook_provider_adjustendpoints_errors_total 0
# HELP external_dns_webhook_provider_adjustendpoints_requests_total Requests with AdjustEndpoints method
# TYPE external_dns_webhook_provider_adjustendpoints_requests_total gauge
external_dns_webhook_provider_adjustendpoints_requests_total 0
# HELP external_dns_webhook_provider_applychanges_errors_total Errors with ApplyChanges method
# TYPE external_dns_webhook_provider_applychanges_errors_total gauge
external_dns_webhook_provider_applychanges_errors_total 0
# HELP external_dns_webhook_provider_applychanges_requests_total Requests with ApplyChanges method
# TYPE external_dns_webhook_provider_applychanges_requests_total gauge
external_dns_webhook_provider_applychanges_requests_total 0
# HELP external_dns_webhook_provider_records_errors_total Errors with Records method
# TYPE external_dns_webhook_provider_records_errors_total gauge
external_dns_webhook_provider_records_errors_total 0
# HELP external_dns_webhook_provider_records_requests_total Requests with Records method
# TYPE external_dns_webhook_provider_records_requests_total gauge
external_dns_webhook_provider_records_requests_total 0
# HELP go_gc_duration_seconds A summary of the wall-time pause (stop-the-world) duration in garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 5.881e-05
go_gc_duration_seconds{quantile="0.25"} 6.9309e-05
go_gc_duration_seconds{quantile="0.5"} 7.599e-05
go_gc_duration_seconds{quantile="0.75"} 0.000166216
go_gc_duration_seconds{quantile="1"} 0.000197552
go_gc_duration_seconds_sum 0.000800889
go_gc_duration_seconds_count 8
# HELP go_gc_gogc_percent Heap size target percentage configured by the user, otherwise 100. This value is set by the GOGC environment variable, and the runtime/debug.SetGCPercent function. Sourced from /gc/gogc:percent.
# TYPE go_gc_gogc_percent gauge
go_gc_gogc_percent 100
# HELP go_gc_gomemlimit_bytes Go runtime memory limit configured by the user, otherwise math.MaxInt64. This value is set by the GOMEMLIMIT environment variable, and the runtime/debug.SetMemoryLimit function. Sourced from /gc/gomemlimit:bytes.
# TYPE go_gc_gomemlimit_bytes gauge
go_gc_gomemlimit_bytes 9.223372036854776e+18
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 9
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.24.0"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated in heap and currently in use. Equals to /memory/classes/heap/objects:bytes.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 5.68044e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated in heap until now, even if released already. Equals to /gc/heap/allocs:bytes.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 2.8783904e+07
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table. Equals to /memory/classes/profiling/buckets:bytes.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 10715
# HELP go_memstats_frees_total Total number of heap objects frees. Equals to /gc/heap/frees:objects + /gc/heap/tiny/allocs:objects.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 482536
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata. Equals to /memory/classes/metadata/other:bytes.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 3.660824e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and currently in use, same as go_memstats_alloc_bytes. Equals to /memory/classes/heap/objects:bytes.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 5.68044e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used. Equals to /memory/classes/heap/released:bytes + /memory/classes/heap/free:bytes.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 6.905856e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use. Equals to /memory/classes/heap/objects:bytes + /memory/classes/heap/unused:bytes
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 8.822784e+06
# HELP go_memstats_heap_objects Number of currently allocated objects. Equals to /gc/heap/objects:objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 32088
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS. Equals to /memory/classes/heap/released:bytes.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 2.998272e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system. Equals to /memory/classes/heap/objects:bytes + /memory/classes/heap/unused:bytes + /memory/classes/heap/released:bytes + /memory/classes/heap/free:bytes.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 1.572864e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.74024692988661e+09
# HELP go_memstats_mallocs_total Total number of heap objects allocated, both live and gc-ed. Semantically a counter version for go_memstats_heap_objects gauge. Equals to /gc/heap/allocs:objects + /gc/heap/tiny/allocs:objects.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 514624
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures. Equals to /memory/classes/metadata/mcache/inuse:bytes.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 19328
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system. Equals to /memory/classes/metadata/mcache/inuse:bytes + /memory/classes/metadata/mcache/free:bytes.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 31408
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures. Equals to /memory/classes/metadata/mspan/inuse:bytes.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 240640
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system. Equals to /memory/classes/metadata/mspan/inuse:bytes + /memory/classes/metadata/mspan/free:bytes.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 310080
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place. Equals to /gc/heap/goal:bytes.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 1.0913642e+07
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations. Equals to /memory/classes/other:bytes.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 2.301997e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes obtained from system for stack allocator in non-CGO environments. Equals to /memory/classes/heap/stacks:bytes.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 1.048576e+06
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator. Equals to /memory/classes/heap/stacks:bytes + /memory/classes/os-stacks:bytes.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 1.048576e+06
# HELP go_memstats_sys_bytes Number of bytes obtained from system. Equals to /memory/classes/total:byte.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 2.309224e+07
# HELP go_sched_gomaxprocs_threads The current runtime.GOMAXPROCS setting, or the number of operating system threads that can execute user-level Go code simultaneously. Sourced from /sched/gomaxprocs:threads.
# TYPE go_sched_gomaxprocs_threads gauge
go_sched_gomaxprocs_threads 16
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 23
# HELP http_request_duration_seconds The HTTP request latencies in seconds.
# TYPE http_request_duration_seconds summary
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="hostedzone",query="",scheme="https",status="200",quantile="0.5"} 0.174236335
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="hostedzone",query="",scheme="https",status="200",quantile="0.9"} 0.539617333
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="hostedzone",query="",scheme="https",status="200",quantile="0.99"} 0.885243225
http_request_duration_seconds_sum{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="hostedzone",query="",scheme="https",status="200"} 3.131792749
http_request_duration_seconds_count{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="hostedzone",query="",scheme="https",status="200"} 12
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="rrset",query="",scheme="https",status="200",quantile="0.5"} 0.135893023
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="rrset",query="",scheme="https",status="200",quantile="0.9"} 0.138746565
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="rrset",query="",scheme="https",status="200",quantile="0.99"} 0.138746565
http_request_duration_seconds_sum{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="rrset",query="",scheme="https",status="200"} 0.274639588
http_request_duration_seconds_count{handler="instrumented_http",host="route53.amazonaws.com",method="GET",path="rrset",query="",scheme="https",status="200"} 2
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="POST",path="hostedzone",query="",scheme="https",status="200",quantile="0.5"} 0.187784949
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="POST",path="hostedzone",query="",scheme="https",status="200",quantile="0.9"} 0.215379492
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="POST",path="hostedzone",query="",scheme="https",status="200",quantile="0.99"} 0.335407224
http_request_duration_seconds_sum{handler="instrumented_http",host="route53.amazonaws.com",method="POST",path="hostedzone",query="",scheme="https",status="200"} 20.023025360999995
http_request_duration_seconds_count{handler="instrumented_http",host="route53.amazonaws.com",method="POST",path="hostedzone",query="",scheme="https",status="200"} 102
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="POST",path="rrset",query="",scheme="https",status="200",quantile="0.5"} 0.17372937
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="POST",path="rrset",query="",scheme="https",status="200",quantile="0.9"} 0.365078338
http_request_duration_seconds{handler="instrumented_http",host="route53.amazonaws.com",method="POST",path="rrset",query="",scheme="https",status="200",quantile="0.99"} 0.365078338
http_request_duration_seconds_sum{handler="instrumented_http",host="route53.amazonaws.com",method="POST",path="rrset",query="",scheme="https",status="200"} 0.538807708
http_request_duration_seconds_count{handler="instrumented_http",host="route53.amazonaws.com",method="POST",path="rrset",query="",scheme="https",status="200"} 2
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0.341346
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 122880
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 11
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 4.653056e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.740246856e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 3.779629056e+10
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes 9.223372036854776e+18
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 1
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP registry_errors_total Number of Registry errors.
# TYPE registry_errors_total counter
registry_errors_total 0
# HELP source_errors_total Number of Source errors.
# TYPE source_errors_total counter
source_errors_total 0
```
