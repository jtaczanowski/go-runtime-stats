package collector

import (
	"runtime"
	"sync"
	"time"
)

type cpu struct {
	// CPU
	NumCPU       int64 `json:"cpu.count"`
	NumGoroutine int64 `json:"cpu.goroutines_number"`
	CgoCallDelta int64 `json:"cpu.cgo_calls_number_delta"`
	NumCgoCall   int64 `json:"cpu.cgo_calls_number_total"`
}

type memory struct {
	// General
	Alloc        int64 `json:"mem.general.alloc_bytes"`
	TotalAlloc   int64 `json:"mem.general.total_bytes"`
	Sys          int64 `json:"mem.general.sys_bytes"`
	LookupsDelta int64 `json:"mem.general.lookups_number_delta"`
	MallocsDelta int64 `json:"mem.general.mallocs_number_delta"`
	FreesDelta   int64 `json:"mem.general.frees_number_delta"`
	Lookups      int64 `json:"mem.general.lookups_number_total"`
	Mallocs      int64 `json:"mem.general.mallocs_number_total"`
	Frees        int64 `json:"mem.general.frees_number_total"`

	// Heap
	HeapAlloc    int64 `json:"mem.heap.alloc_bytes"`
	HeapSys      int64 `json:"mem.heap.sys_bytes"`
	HeapIdle     int64 `json:"mem.heap.idle_bytes"`
	HeapInuse    int64 `json:"mem.heap.inuse_bytes"`
	HeapReleased int64 `json:"mem.heap.released_bytes"`
	HeapObjects  int64 `json:"mem.heap.objects_number"`

	// Stack
	StackInuse  int64 `json:"mem.stack.inuse_bytes"`
	StackSys    int64 `json:"mem.stack.sys_bytes"`
	MSpanInuse  int64 `json:"mem.stack.mspan_inuse_bytes"`
	MSpanSys    int64 `json:"mem.stack.mspan_sys_bytes"`
	MCacheInuse int64 `json:"mem.stack.mcache_inuse_bytes"`
	MCacheSys   int64 `json:"mem.stack.mcache_sys_bytes"`

	OtherSys int64 `json:"mem.othersys_bytes"`
}

type gc struct {
	// GC
	GCSys             int64   `json:"gc.sys_bytes"`
	NextGC            int64   `json:"gc.next_bytes"`
	BetweenGCPerdiod  int64   `json:"gc.between_period_s"`
	LastGC            int64   `json:"-"`
	TimeFromLastGC    int64   `json:"gc.time_from_last_gc_s"`
	PauseTotalNsDelta int64   `json:"gc.pause_ns_total_delta"`
	PauseTotalNs      int64   `json:"gc.pause_ns_total"`
	PauseNs           int64   `json:"gc.pause_ns"`
	LastPauseNs       int64   `json:"gc.pause_last_ns"`
	NumGCDelta        int64   `json:"gc.number_delta"`
	NumGC             int64   `json:"gc.number_total"`
	GCCPUFraction     float64 `json:"gc.cpu_fraction_total"`
}

type Collector struct {
	*cpu    `json:"cpu"`
	*memory `json:"memory"`
	*gc     `json:"gc"`
	Mutex   *sync.Mutex `json:"-"`
}

func NewCollector() *Collector {
	return &Collector{&cpu{}, &memory{}, &gc{}, &sync.Mutex{}}
}

func (c *Collector) CollectStats() {
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)

	c.Mutex.Lock()
	c.collectCPUStats()
	c.collectMemoryStats(m)
	c.collectGcStats(m)
	c.Mutex.Unlock()

}

func (c *Collector) collectCPUStats() {
	c.cpu.NumCPU = int64(runtime.NumCPU())
	c.cpu.NumGoroutine = int64(runtime.NumGoroutine())
	c.cpu.NumCgoCall = int64(runtime.NumCgoCall())
	c.cpu.CgoCallDelta = int64(runtime.NumCgoCall()) - c.cpu.NumCgoCall
	c.cpu.NumCgoCall = int64(runtime.NumCgoCall())
}

func (c *Collector) collectMemoryStats(m *runtime.MemStats) {

	// General
	c.memory.Alloc = int64(m.Alloc)
	c.memory.TotalAlloc = int64(m.TotalAlloc)
	c.memory.Sys = int64(m.Sys)
	c.memory.LookupsDelta = int64(m.Lookups) - c.memory.Lookups
	c.memory.MallocsDelta = int64(m.Mallocs) - c.memory.Mallocs
	c.memory.FreesDelta = int64(m.Frees) - c.memory.Frees
	c.memory.Lookups = int64(m.Lookups)
	c.memory.Mallocs = int64(m.Mallocs)
	c.memory.Frees = int64(m.Frees)

	// Heap
	c.memory.HeapAlloc = int64(m.HeapAlloc)
	c.memory.HeapSys = int64(m.HeapSys)
	c.memory.HeapIdle = int64(m.HeapIdle)
	c.memory.HeapInuse = int64(m.HeapInuse)
	c.memory.HeapReleased = int64(m.HeapReleased)
	c.memory.HeapObjects = int64(m.HeapObjects)

	// Stack
	c.memory.StackInuse = int64(m.StackInuse)
	c.memory.StackSys = int64(m.StackSys)
	c.memory.MSpanInuse = int64(m.MSpanInuse)
	c.memory.MSpanSys = int64(m.MSpanSys)
	c.memory.MCacheInuse = int64(m.MCacheInuse)
	c.memory.MCacheSys = int64(m.MCacheSys)

	c.memory.OtherSys = int64(m.OtherSys)

}

func (c *Collector) collectGcStats(m *runtime.MemStats) {
	c.gc.GCSys = int64(m.GCSys)
	c.gc.NextGC = int64(m.NextGC)
	if int64(m.LastGC) > c.gc.LastGC {
		c.gc.PauseNs = int64(m.PauseNs[(m.NumGC+255)%256])
		c.gc.LastPauseNs = int64(m.PauseNs[(m.NumGC+255)%256])
	}
	if int64(m.LastGC) != 0 {
		// time in second from last gc cycle
		c.gc.TimeFromLastGC = time.Now().Unix() - (int64(m.LastGC) / 1000000000)
	}
	if c.gc.LastGC != 0 && int64(m.LastGC) > c.gc.LastGC {
		// c.gc.BetweenGCPerdiod - calculated time period between last and previous GC
		c.gc.BetweenGCPerdiod = int64(m.LastGC) - c.gc.LastGC
		// convert to nano seconds to seconds
		c.gc.BetweenGCPerdiod = c.gc.BetweenGCPerdiod / 1000000000
	}
	// if there was no GC cycle during last goruntimestats interval (default 60s) set BetweenGCPerdiod and c.gc.PauseNs to 0
	if int64(m.LastGC) == c.gc.LastGC {
		c.gc.BetweenGCPerdiod = 0
		c.gc.PauseNs = 0
	}
	c.gc.PauseTotalNsDelta = int64(m.PauseTotalNs) - c.gc.PauseTotalNs
	c.gc.PauseTotalNs = int64(m.PauseTotalNs)
	c.gc.LastGC = int64(m.LastGC)
	c.gc.NumGCDelta = int64(m.NumGC) - c.gc.NumGC
	c.gc.NumGC = int64(m.NumGC)
	c.gc.GCCPUFraction = float64(m.GCCPUFraction)
}
