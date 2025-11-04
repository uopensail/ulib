package prome

import (
	"fmt"
	"net/http"

	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// serverLabelNames defines the label names used for server metrics
	serverLabelNames = []string{"name", "status"}
)

// ServerExporter handles the collection and export of server performance metrics
type ServerExporter struct {
	httpServer *http.Server

	maxCostTime *prometheus.Desc
	avgCounter  *prometheus.Desc
	p90CostTime *prometheus.Desc
	p95CostTime *prometheus.Desc
	p99CostTime *prometheus.Desc
	qps         *prometheus.Desc
	avgCostTime *prometheus.Desc
	costBucket  *prometheus.Desc
}

// NewServerExporter creates a new ServerExporter instance with the specified namespace
//
// @param namespace: The namespace for Prometheus metrics
// @return: Pointer to initialized ServerExporter
func NewServerExporter(namespace string) *ServerExporter {
	serverE := ServerExporter{
		maxCostTime: newServerDesc(namespace, "max_cost_time", "Max Cost Time"),
		avgCostTime: newServerDesc(namespace, "avg_cost_time", "Average Cost Time"),
		p90CostTime: newServerDesc(namespace, "p90_cost_time", "90th Percentile Of Cost Time"),
		p95CostTime: newServerDesc(namespace, "p95_cost_time", "95th Percentile Of Cost Time"),
		p99CostTime: newServerDesc(namespace, "p99_cost_time", "99th Percentile Of Cost Time"),
		avgCounter:  newServerDesc(namespace, "avg_counter", "Average Counter"),
		qps:         newServerDesc(namespace, "qps", "Queries Per Second"),
		costBucket:  newServerDesc(namespace, "cost", "Bucket Of Cost Time"),
	}
	return &serverE
}

// Describe sends the metric descriptors to the provided channel
// Implements prometheus.Collector interface
func (serverE *ServerExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- serverE.maxCostTime
	ch <- serverE.avgCostTime
	ch <- serverE.avgCounter
	ch <- serverE.qps
	ch <- serverE.costBucket
	// Fixed bug: Added missing p90, p95, p99 descriptors
	ch <- serverE.p90CostTime
	ch <- serverE.p95CostTime
	ch <- serverE.p99CostTime
}

// Exporter combines ServerExporter with additional exporter-specific metrics
type Exporter struct {
	ServerExporter
	port int

	up           prometheus.Gauge
	totalScrapes prometheus.Counter
}

// newServerDesc creates a new Prometheus descriptor with the given parameters
//
// @param namespace: Metric namespace
// @param metricName: Name of the metric
// @param docString: Help text describing the metric
// @return: Prometheus descriptor
func newServerDesc(namespace, metricName string, docString string) *prometheus.Desc {
	return prometheus.NewDesc(
		fmt.Sprintf("%s_%s", namespace, metricName),
		docString,
		serverLabelNames,
		nil,
	)
}

// NewExporter creates a new Exporter instance with the specified namespace
//
// @param namespace: The namespace for Prometheus metrics
// @return: Pointer to initialized Exporter
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

// Describe sends all metric descriptors to the provided channel
// Implements prometheus.Collector interface
func (exporter *Exporter) Describe(ch chan<- *prometheus.Desc) {
	exporter.ServerExporter.Describe(ch)
	ch <- exporter.up.Desc()
	ch <- exporter.totalScrapes.Desc()
}

// Collect gathers and sends all metrics to the provided channel
// Implements prometheus.Collector interface
func (exporter *Exporter) Collect(ch chan<- prometheus.Metric) {
	defer func() {
		if err := recover(); err != nil {
			zlog.LOG.Error(fmt.Sprintf("exporter collect error:%v", err))
		}
	}()

	exporter.scrape(ch)

	ch <- exporter.up
	ch <- exporter.totalScrapes

	zlog.LOG.Debug("prometheus collect done") // Changed to Debug to reduce log noise
}

// scrape performs the actual metric collection and sends metrics to the channel
func (exporter *Exporter) scrape(ch chan<- prometheus.Metric) {
	exporter.totalScrapes.Inc()

	minfo := GlobalmetricsIns.GetMetricsInfo()

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
				ch <- prometheus.MustNewConstMetric(exporter.p90CostTime, prometheus.GaugeValue,
					float64(mi.P90Cost), labelvs...)
				ch <- prometheus.MustNewConstMetric(exporter.p95CostTime, prometheus.GaugeValue,
					float64(mi.P95Cost), labelvs...)
				ch <- prometheus.MustNewConstMetric(exporter.p99CostTime, prometheus.GaugeValue,
					float64(mi.P99Cost), labelvs...)

				// Fixed bug: Added nil check for histogram to prevent panic
				if mi.CostBucket != nil {
					ch <- prometheus.MustNewConstHistogram(exporter.costBucket, uint64(mi.Total),
						mi.AvgCost*float64(mi.Total), mi.CostBucket, labelvs...)
				} else {
					zlog.LOG.Warn("CostBucket is nil for metric", zap.String("name", mi.Name))
				}
			}
		}
	}
}

// Start begins the Prometheus metrics HTTP server on the specified port
//
// @param port: The port number to listen on
// @return: error if the server fails to start
func (exporter *Exporter) Start(port int) error {
	// Fixed bug: Check if exporter is already registered
	if !prometheus.Unregister(exporter) {
		zlog.LOG.Debug("Exporter was not previously registered, proceeding with registration")
	}

	err := prometheus.Register(exporter)
	if err != nil {
		return fmt.Errorf("register prometheus fail: %w", err) // Enhanced error wrapping
	}

	portListenStr := fmt.Sprintf(":%d", port)
	exporter.httpServer = &http.Server{
		Addr:    portListenStr,
		Handler: promhttp.Handler(),
	}

	zlog.LOG.Info("Starting Prometheus metrics server", zap.Int("port", port))
	return exporter.httpServer.ListenAndServe()
}

// Close shuts down the HTTP server gracefully
func (exporter *Exporter) Close() error {
	if exporter.httpServer != nil {
		zlog.LOG.Info("Shutting down Prometheus metrics server")
		return exporter.httpServer.Close()
	}
	return nil
}
