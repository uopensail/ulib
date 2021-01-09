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
	httpServer    *http.Server
	maxCostTime   *prometheus.CounterVec
	avgCount      *prometheus.CounterVec
	totalRequst   *prometheus.CounterVec
	qps           *prometheus.CounterVec
	minCostTime   *prometheus.CounterVec
	avgCostTime   *prometheus.CounterVec
	avgCostBucket *prometheus.HistogramVec
}

func NewServerExporter(namespace string) *ServerExporter {
	serverE := ServerExporter{
		maxCostTime: newServerMetric(namespace, "max_cost_time", "最大耗时"),
		minCostTime: newServerMetric(namespace, "min_cost_time", "最小耗时"),
		avgCostTime: newServerMetric(namespace, "avg_cost_time", "平均耗时"),
		avgCount:    newServerMetric(namespace, "avg_conut", "平均计数"),
		totalRequst: newServerMetric(namespace, "total_request", "请求的总数量"),
		qps:         newServerMetric(namespace, "qps", "qps"),
		avgCostBucket: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "cost_bucket",
			Help:      "耗时分桶",
		}, serverLabelNames),
	}
	return &serverE
}

func (serverE *ServerExporter) Reset() {
	serverE.maxCostTime.Reset()
	serverE.minCostTime.Reset()
	serverE.avgCostTime.Reset()
	serverE.avgCount.Reset()
	serverE.totalRequst.Reset()
	serverE.qps.Reset()
}

func (serverE *ServerExporter) Describe(ch chan<- *prometheus.Desc) {
	serverE.maxCostTime.Describe(ch)
	serverE.minCostTime.Describe(ch)
	serverE.avgCostTime.Describe(ch)
	serverE.avgCount.Describe(ch)
	serverE.totalRequst.Describe(ch)
	serverE.qps.Describe(ch)

}

func (serverE *ServerExporter) Collect(ch chan<- prometheus.Metric) {
	serverE.maxCostTime.Collect(ch)
	serverE.minCostTime.Collect(ch)
	serverE.avgCostTime.Collect(ch)
	serverE.avgCount.Collect(ch)
	serverE.totalRequst.Collect(ch)
	serverE.qps.Collect(ch)

}

type Exporter struct {
	ServerExporter
	port int

	up           prometheus.Gauge
	totalScrapes prometheus.Counter
}

func newServerMetric(namespace, metricName string, docString string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      metricName,
			Help:      docString,
		},
		serverLabelNames)
}

func NewExporter(namespace string) *Exporter {

	export := Exporter{
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "up",
			Help:      "Was the last scrape of server successful.",
		}),
		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "exporter_total_scrapes",
			Help:      "Current total server scrapes.",
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

	exporter.ServerExporter.Reset()

	exporter.scrape()

	ch <- exporter.up
	ch <- exporter.totalScrapes
	exporter.ServerExporter.Collect(ch)

	zlog.LOG.Info("prometheus collect done")
}

func (exporter *Exporter) scrape() {
	exporter.totalScrapes.Inc()

	minfo := metricsInstance.GetMetricsInfo()

	if len(minfo) == 0 {
		exporter.up.Set(0)
	} else {
		exporter.up.Set(1)
		for _, mi := range minfo {
			maplabels := make(map[string]string, 2)
			maplabels["name"] = mi.Name
			maplabels["status"] = mi.Status
			labels := prometheus.Labels(maplabels)
			exporter.avgCount.With(labels).Add(float64(mi.AvgCounter))
			exporter.totalRequst.With(labels).Add(float64(mi.Total))
			exporter.qps.With(labels).Add(float64(mi.QPS))
			exporter.maxCostTime.With(labels).Add(float64(mi.MaxCost))
			exporter.avgCostTime.With(labels).Add(float64(mi.AvgCost))
			//exporter.avgCostBucket.With(labels)
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
