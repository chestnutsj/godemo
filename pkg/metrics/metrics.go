package metrics

import (
	"github.com/chestnutsj/godemo/pkg/tools"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var startTime prometheus.Gauge

func init() {
	startTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "demo",
		Subsystem: tools.AppName(),
		Name:      "start_time",
		Help:      "this app start time",
	})
	prometheus.MustRegister(startTime)
}

func StartMetrics(port string) {
	address := ":" + port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	log.Println("Using Metrics port:", listener.Addr().(*net.TCPAddr).Port)

	http.Handle("/metrics", promhttp.Handler())
	log.Println("localhost:" + strconv.Itoa(listener.Addr().(*net.TCPAddr).Port) + "/metrics")
	startTime.SetToCurrentTime()
	panic(http.Serve(listener, nil))
}
