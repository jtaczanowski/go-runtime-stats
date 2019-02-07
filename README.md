# go-runtime-stats [![Build Status](https://travis-ci.org/jtaczanowski/go-runtime-stats.png?branch=master)](https://travis-ci.org/jtaczanowski/go-runtime-stats) [![Coverage Status](https://coveralls.io/repos/github/jtaczanowski/go-runtime-stats/badge.svg?branch=master)](https://coveralls.io/github/jtaczanowski/go-runtime-stats?branch=master)

go-runtime-stats - Golang package providing collecting go runtime stats and sending it to graphite server.

Example usage (also included in _example catalog):

```go
package main

import "github.com/jtaczanowski/go-runtime-stats"

func main() {
	goruntimestats.Start(goruntimestats.Config{
		GraphiteHost:     "127.0.0.1",
		GraphitePort:     2003,
		GraphiteProtocol: "udp",
		GraphitePrefix:   "metrics.prefix",
		Interval:         10,
		HTTPOn:           true,
		HTTPPort:         9999,
	})

	// insted of empty select (used for block) put your code
	select {}
}
```

![alt text](_example/go-runtime-stats.png "Grafana Dashboard")

| Metric                           | Source                           | Description                                                        | Unit               |
|----------------------------------|----------------------------------|--------------------------------------------------------------------|--------------------|
| cpu.count                        | runtime.runtime.NumCPU()         | Number of machine CPU                                              | number of CPUs     |
| cpu.goroutines_number            | runtime.NumGoroutine()           | Number of goroutines                                               | number             |
| cpu.cgo_calls_number_delta       | runtime.NumCgoCall()             | Delta of Cgo calls from last collect interval                      | number             |
| cpu.cgo_calls_number_total       | runtime.NumCgoCall()             | Summ of Cgo calls                                                  | number             |
| mem.general.alloc_bytes          | runtime.MemStats                 | Alloc is bytes of allocated heap objects.                          | bytes              |
| mem.general.total_bytes          | runtime.TotalAlloc               | TotalAlloc is cumulative bytes allocated for heap objects.         | bytes              |
| mem.general.sys_bytes            | runtime.Sys                      | Sys is the total bytes of memory obtained from the OS.             | bytes              |
| mem.general.lookups_number_delta | runtime.Lookups                  | Delta of Lookups from last collect interval                        | number             |
| mem.general.mallocs_number_delta | runtime.Mallocs                  | Delta of Mallocs from last collect interval                        | number             |
| mem.general.frees_number_delta   | runtime.Frees                    | Delta of Frees from last collect interval                          | number             |
| mem.general.lookups_number_total | runtime.Lookups                  | Lookups is the number of pointer lookups performed by the runtime. | number             |
| mem.general.mallocs_number_total | runtime.Mallocs                  | Mallocs is the cumulative count of heap objects allocated.         | number             |
| mem.general.frees_number_total   | runtime.Frees                    | Frees is the cumulative count of heap objects freed.               | number             |

| mem.heap.alloc_bytes             | runtime.ReadMemStats.Frees       | Number of frees issued to the system   | frees per second   |
| mem.heap.sys_bytes               | runtime.ReadMemStats.Mallocs     | Number of Mallocs issued to the system | mallocs per second |
| mem.heap.idle_bytes              | runtime.ReadMemStats.HeapIdle    | Memory on the heap not in use          | bytes              |
| mem.heap.inuse_bytes             | runtime.ReadMemStats.HeapInuse   | Memory on the heap in use              | bytes              |
| mem.heap.released_bytes          | runtime.ReadMemStats.HeapObjects | Total objects on the heap              | # Objects          |
| mmem.heap.objects_number         | runtime.ReadMemStats.Alloc       | Total bytes allocated                  | bytes              |
| mem.stack.inuse_bytes            | runtime.ReadMemStats.HeapSys     | Total bytes acquired from system       | bytes              |

| mem.stack.sys_bytes              | runtime.ReadMemStats.Frees       | Number of frees issued to the system   | frees per second   |
| mem.stack.mspan_inuse_bytes      | runtime.ReadMemStats.Mallocs     | Number of Mallocs issued to the system | mallocs per second |
| mem.stack.mspan_sys_bytes        | runtime.ReadMemStats.HeapIdle    | Memory on the heap not in use          | bytes              |
| mem.stack.mcache_inuse_bytes     | runtime.ReadMemStats.HeapInuse   | Memory on the heap in use              | bytes              |
| mem.stack.mcache_sys_bytes       | runtime.ReadMemStats.HeapObjects | Total objects on the heap              | # Objects          |
| mem.othersys_bytes               | runtime.ReadMemStats.Alloc       | Total bytes allocated                  | bytes              |

| gc.sys_bytes                     | runtime.ReadMemStats.Frees       | Number of frees issued to the system   | frees per second   |
| gc.next_bytes                    | runtime.ReadMemStats.Mallocs     | Number of Mallocs issued to the system | mallocs per second |
| gc.between_period_s              | runtime.ReadMemStats.HeapIdle    | Memory on the heap not in use          | bytes              |
| gc.time_from_last_gc_s           | runtime.ReadMemStats.HeapInuse   | Memory on the heap in use              | bytes              |
| gc.pause_ns_total_delta          | runtime.ReadMemStats.HeapObjects | Total objects on the heap              | # Objects          |
| gc.pause_ns_total                | runtime.ReadMemStats.Alloc       | Total bytes allocated                  | bytes              |
| gc.pause_ns                      | runtime.ReadMemStats.Frees       | Number of frees issued to the system   | frees per second   |
| gc.pause_last_ns                 | runtime.ReadMemStats.Mallocs     | Number of Mallocs issued to the system | mallocs per second |
| gc.number_delta                  | runtime.ReadMemStats.HeapIdle    | Memory on the heap not in use          | bytes              |
| gc.number_total                  | runtime.ReadMemStats.HeapInuse   | Memory on the heap in use              | bytes              |
| gc.cpu_fraction_total            | runtime.ReadMemStats.HeapObjects | Total objects on the heap              | # Objects          |
>>>>>>> edit README.md
