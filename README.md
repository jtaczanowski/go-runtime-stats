# go-runtime-stats [![Build Status](https://travis-ci.org/jtaczanowski/go-runtime-stats.png?branch=master)](https://travis-ci.org/jtaczanowski/go-runtime-stats)

go-runtime-stats - Golang package providing collecting go runtime stats and sending it to graphite server.

Example usage:

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
