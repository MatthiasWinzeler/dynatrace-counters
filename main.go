package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var counter = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "dynatrace_test_counter",
	},
	[]string{},
)

func init() {
	// deregister default metrics to reduce noise
	prometheus.Unregister(collectors.NewGoCollector())
}

func main() {
	go func() {
		for _ = range time.Tick(time.Minute) {
			counter.WithLabelValues().Inc()
			log.Println("Inc")
		}
	}()

	log.Println("Starting")

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}