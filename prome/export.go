package prome

import (
	"fmt"
	"net/http"

	"github.com/uopensail/ulib/zlog"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	serverLabelNames = []string{"name", "status"}
)

type ServerExporter struct {
	httpServer *http.Server

	maxCostTime *prometheus.Desc
	avgCounter  *prometheus.Desc
	qps         *prometheus.Desc
	avgCostTime *prometheus.Desc
	costBucket  *prometheus.Desc
}

func NewServerExporter(namespace string) *ServerExporter {
	serverE := ServerExporter{
		maxCostTime: newServerDesc(namespace, "max_cost_time", "最大耗时"),
		avgCostTime: newServerDesc(namespace, "avg_cost_time", "平均耗时"),
		avgCounter:  newServerDesc(namespace, "avg_counter", "平均计数"),
		qps:         newServerDesc(namespace, "qps", "qps"),
		costBucket:  newServerDesc(namespace, "cost", "耗时分桶"),
	}
	return &serverE
}

func (serverE *ServerExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- serverE.maxCostTime
	ch <- serverE.avgCostTime
	ch <- serverE.avgCounter
	ch <- serverE.qps
	ch <- serverE.costBucket

}

type Exporter struct {
	ServerExporter
	port int

	up           prometheus.Gauge
	totalScrapes prometheus.Counter
}

func newServerDesc(namespace, metricName string, docString string) *prometheus.Desc {
	return prometheus.NewDesc(fmt.Sprintf("%s_%s", namespace, metricName),
		docString, serverLabelNames, nil)
}

func NewExporter(namespace string) *Exporter {

	export := Exporter{
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "up",
			Help: "Was the last scrape of server successful.",
		}),
		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "exporter_total_scrapes",
			Help: "Current total server scrapes.",
		}),
	}
	export.ServerExporter = *NewServerExporter(namespace)
	return &export
}

// Implements prometheus.Collector.
func (exporter *Exporter) Describe(ch chan<- *prometheus.Desc) {
	exporter.ServerExporter.Describe(ch)
	ch <- exporter.up.Desc()
	ch <- exporter.totalScrapes.Desc()
}

// Implements prometheus.Collector.
func (exporter *Exporter) Collect(ch chan<- prometheus.Metric) {
	defer func() {
		if err := recover(); err != nil {
			zlog.LOG.Error(fmt.Sprintf("exporter collect error:", err))
		}
	}()

	exporter.scrape(ch)

	ch <- exporter.up
	ch <- exporter.totalScrapes

	zlog.LOG.Info("prometheus collect done")
}

func (exporter *Exporter) scrape(ch chan<- prometheus.Metric) {
	exporter.totalScrapes.Inc()

	minfo := metricsInstance.GetMetricsInfo()

	if len(minfo) == 0 {
		exporter.up.Set(0)
	} else {
		exporter.up.Set(1)

		for _, mi := range minfo {
			labelvs := []string{mi.Name, mi.Status}
			ch <- prometheus.MustNewConstMetric(exporter.qps, prometheus.GaugeValue,
				float64(mi.QPS), labelvs...)
			ch <- prometheus.MustNewConstMetric(exporter.avgCounter, prometheus.GaugeValue,
				float64(mi.AvgCounter), labelvs...)
			if mi.Status == "OK" {
				ch <- prometheus.MustNewConstMetric(exporter.avgCostTime, prometheus.GaugeValue,
					float64(mi.AvgCost), labelvs...)
				ch <- prometheus.MustNewConstMetric(exporter.maxCostTime, prometheus.GaugeValue,
					float64(mi.MaxCost), labelvs...)
				ch <- prometheus.MustNewConstHistogram(exporter.costBucket, uint64(mi.Total),
					mi.AvgCost*float64(mi.Total), mi.CostBucket, labelvs...)
			}
		}
	}

}

func (exporter *Exporter) Start(port int) error {
	err := prometheus.Register(exporter)
	if err != nil {
		return fmt.Errorf("register prometheus fail:%s", err)
	}
	portListenStr := fmt.Sprintf(":%d", port)
	exporter.httpServer = &http.Server{Addr: portListenStr, Handler: promhttp.Handler()}
	return exporter.httpServer.ListenAndServe()
}

func (exporter *Exporter) Close() {
	exporter.httpServer.Close()
}
