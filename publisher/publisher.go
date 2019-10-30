package publisher

import (
	"encoding/json"
	"io"
	"log"

	"github.com/jtaczanowski/go-graphite-client"
	"github.com/jtaczanowski/go-runtime-stats/collector"
)

type Publisher struct {
	collector *collector.Collector
	graphite  *graphite.Client
}

// NewPublisher - creates new Publisher struct
func NewPublisher(collector *collector.Collector, graphite *graphite.Client) *Publisher {
	return &Publisher{
		collector: collector,
		graphite:  graphite,
	}
}

// PublishToGraphite - prepares data from Collector struct to graphite metric format and passing it to graphite.SentData
func (p *Publisher) PublishToGraphite() {
	err := p.graphite.SendData(p.prepareDataToSend())
	if err != nil {
		log.Printf("Can not sent data to graphite: %s", err.Error())
	}
}

// prepareDataToSend - prepares data from Collector struct to graphite metric format
func (p *Publisher) prepareDataToSend() []map[string]float64 {
	var collectorMap map[string]map[string]int64
	dataToGraphite := make([]map[string]float64, 0)

	r, w := io.Pipe()
	go func() {
		p.collector.Mutex.Lock()
		json.NewEncoder(w).Encode(p.collector)
		w.Close()
		p.collector.Mutex.Unlock()
	}()
	json.NewDecoder(r).Decode(&collectorMap)

	for metricCategory := range collectorMap {
		for metricName, metricValue := range collectorMap[metricCategory] {
			dataToGraphite = append(dataToGraphite, map[string]float64{metricName: float64(metricValue)})
		}
	}
	return dataToGraphite
}
