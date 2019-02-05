package collector

import (
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestCollectMemoryStatsDeltaCalculations(t *testing.T) {
	testCollector := NewCollector()
	m := &runtime.MemStats{}

	// fist collectMemoryStats run with data below
	m.Lookups = 1
	m.Mallocs = 1
	m.Frees = 1
	testCollector.collectMemoryStats(m)

	// second collectMemoryStats run with data below
	m.Lookups = 3
	m.Mallocs = 3
	m.Frees = 3
	testCollector.collectMemoryStats(m)

	AssertEqual(t, testCollector.memory.LookupsDelta, int64(2))
	AssertEqual(t, testCollector.memory.MallocsDelta, int64(2))
	AssertEqual(t, testCollector.memory.FreesDelta, int64(2))
}

func TestRunCollectGcStatsOnEmptyRuntimeMemStatsStruct(t *testing.T) {
	testCollector := NewCollector()
	m := &runtime.MemStats{}

	testCollector.collectGcStats(m)

	AssertEqual(t, testCollector.gc.GCSys, int64(0))
	AssertEqual(t, testCollector.gc.NextGC, int64(0))
	AssertEqual(t, testCollector.gc.BetweenGCPerdiod, int64(0))
	AssertEqual(t, testCollector.gc.LastGC, int64(0))
	AssertEqual(t, testCollector.gc.TimeFromLastGC, int64(0))
	AssertEqual(t, testCollector.gc.PauseTotalNsDelta, int64(0))
	AssertEqual(t, testCollector.gc.PauseTotalNs, int64(0))
	AssertEqual(t, testCollector.gc.PauseNs, int64(0))
	AssertEqual(t, testCollector.gc.LastPauseNs, int64(0))
	AssertEqual(t, testCollector.gc.NumGCDelta, int64(0))
	AssertEqual(t, testCollector.gc.NumGC, int64(0))
	AssertEqual(t, testCollector.gc.GCCPUFraction, float64(0))
}

func TestOnceRunCollectGcStatsOnFilledMemStatsStruct(t *testing.T) {
	testCollector := NewCollector()
	m := &runtime.MemStats{}
	m.GCSys = 1
	m.NextGC = 2
	m.LastGC = 1000000000
	m.PauseTotalNs = 1
	m.PauseNs = [256]uint64{1}
	m.NumGC = 1
	m.GCCPUFraction = 1

	testCollector.collectGcStats(m)

	AssertEqual(t, testCollector.gc.GCSys, int64(1))
	AssertEqual(t, testCollector.gc.NextGC, int64(2))
	AssertEqual(t, testCollector.gc.BetweenGCPerdiod, int64(0))
	AssertEqual(t, testCollector.gc.LastGC, int64(1000000000))
	AssertEqual(t, testCollector.gc.TimeFromLastGC, time.Now().Unix()-1)
	AssertEqual(t, testCollector.gc.PauseTotalNsDelta, int64(1))
	AssertEqual(t, testCollector.gc.PauseTotalNs, int64(1))
	AssertEqual(t, testCollector.gc.PauseNs, int64(1))
	AssertEqual(t, testCollector.gc.LastPauseNs, int64(1))
	AssertEqual(t, testCollector.gc.NumGCDelta, int64(1))
	AssertEqual(t, testCollector.gc.NumGC, int64(1))
	AssertEqual(t, testCollector.gc.GCCPUFraction, float64(1))
}

func TestDoubleRunCollectGcStatsOnFilledMemStatsStruct(t *testing.T) {
	testCollector := NewCollector()
	m := &runtime.MemStats{}

	// fist collectGcStats run with data below
	m.GCSys = 1
	m.NextGC = 2
	m.LastGC = 1000000000
	m.PauseTotalNs = 1
	m.PauseNs = [256]uint64{1}
	m.NumGC = 1
	m.GCCPUFraction = 1
	testCollector.collectGcStats(m)

	// second collectGcStats run with data below
	m.GCSys = 2
	m.NextGC = 3
	m.LastGC = 2000000000
	m.PauseTotalNs = 4
	m.PauseNs = [256]uint64{1, 3}
	m.NumGC = 2
	m.GCCPUFraction = 2
	testCollector.collectGcStats(m)

	AssertEqual(t, testCollector.gc.GCSys, int64(2))
	AssertEqual(t, testCollector.gc.NextGC, int64(3))
	AssertEqual(t, testCollector.gc.BetweenGCPerdiod, int64(1))
	AssertEqual(t, testCollector.gc.LastGC, int64(2000000000))
	AssertEqual(t, testCollector.gc.TimeFromLastGC, time.Now().Unix()-2)
	AssertEqual(t, testCollector.gc.PauseTotalNsDelta, int64(3))
	AssertEqual(t, testCollector.gc.PauseTotalNs, int64(4))
	AssertEqual(t, testCollector.gc.PauseNs, int64(3))
	AssertEqual(t, testCollector.gc.LastPauseNs, int64(3))
	AssertEqual(t, testCollector.gc.NumGCDelta, int64(1))
	AssertEqual(t, testCollector.gc.NumGC, int64(2))
	AssertEqual(t, testCollector.gc.GCCPUFraction, float64(2))
}

// AssertEqual checks if values are equal
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}
