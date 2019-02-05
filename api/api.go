package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jtaczanowski/go-runtime-stats/collector"
)

type HttpServer struct {
	port      int
	collector *collector.Collector
}

func NewHttpServer(port int, collector *collector.Collector) *HttpServer {
	return &HttpServer{
		port:      port,
		collector: collector,
	}
}

func (h *HttpServer) Start() error {
	var httpServerError = make(chan error)
	http.HandleFunc("/", h.apiHandler) // set router
	go func() {
		httpServerError <- http.ListenAndServe(":"+strconv.Itoa(h.port), nil)
	}()
	select {
	case err := <-httpServerError:
		fmt.Printf("Goruntime stats http server could not be started: %s", err.Error())
		return err
	default:
		fmt.Printf("Goruntime stats http server started http://0.0.0.0:%s\n", strconv.Itoa(h.port))
	}
	return nil
}

func (h *HttpServer) apiHandler(w http.ResponseWriter, r *http.Request) {
	h.collector.Mutex.Lock()
	var prettyJSON bytes.Buffer
	uglyJSON, _ := json.Marshal(h.collector)
	json.Indent(&prettyJSON, uglyJSON, "", "  ")
	h.collector.Mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.Write(prettyJSON.Bytes())
}
