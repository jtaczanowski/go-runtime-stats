package main

import (
	"time"

	"github.com/jtaczanowski/go-runtime-stats"
)

func main() {
	goruntimestats.Start(goruntimestats.Config{
		GraphiteHost:     "127.0.0.1",
		GraphitePort:     2003,
		GraphiteProtocol: "udp",
		GraphitePrefix:   "metrics.prefix",
		Interval:         time.Second * 60,
		HTTPOn:           true,
		HTTPPort:         9999,
	})

	// insted of empty select (used for block) put your code
	select {}
}
