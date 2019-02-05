package api

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/jtaczanowski/go-runtime-stats/collector"
)

func TestApi(t *testing.T) {
	testCollector := collector.NewCollector()
	testAPI := NewHttpServer(9999, testCollector)

	req := httptest.NewRequest("GET", "http://localhost:9999", nil)
	w := httptest.NewRecorder()
	testAPI.apiHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	excepted := `{
  "cpu": {
    "cpu.count": 0,
    "cpu.goroutines_number": 0,
    "cpu.cgo_calls_number_delta": 0,
    "cpu.cgo_calls_number_total": 0
  },
  "memory": {
    "mem.general.alloc_bytes": 0,
    "mem.general.total_bytes": 0,
    "mem.general.sys_bytes": 0,
    "mem.general.lookups_number_delta": 0,
    "mem.general.mallocs_number_delta": 0,
    "mem.general.frees_number_delta": 0,
    "mem.general.lookups_number_total": 0,
    "mem.general.mallocs_number_total": 0,
    "mem.general.frees_number_total": 0,
    "mem.heap.alloc_bytes": 0,
    "mem.heap.sys_bytes": 0,
    "mem.heap.idle_bytes": 0,
    "mem.heap.inuse_bytes": 0,
    "mem.heap.released_bytes": 0,
    "mem.heap.objects_number": 0,
    "mem.stack.inuse_bytes": 0,
    "mem.stack.sys_bytes": 0,
    "mem.stack.mspan_inuse_bytes": 0,
    "mem.stack.mspan_sys_bytes": 0,
    "mem.stack.mcache_inuse_bytes": 0,
    "mem.stack.mcache_sys_bytes": 0,
    "mem.othersys_bytes": 0
  },
  "gc": {
    "gc.sys_bytes": 0,
    "gc.next_bytes": 0,
    "gc.between_period_s": 0,
    "gc.time_from_last_gc_s": 0,
    "gc.pause_ns_total_delta": 0,
    "gc.pause_ns_total": 0,
    "gc.pause_ns": 0,
    "gc.pause_last_ns": 0,
    "gc.number_delta": 0,
    "gc.number_total": 0,
    "gc.cpu_fraction_total": 0
  }
}`

	if string(body) != excepted {
		t.Errorf("Body from api different then excepted")
	}
	if resp.StatusCode != 200 {
		t.Errorf("Status code: %v from api different then excepted", resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type: %v from api different then excepted", resp.Header.Get("Content-Type"))
	}
}
