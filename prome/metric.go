package prome

import (
	"bytes"
	"encoding/json"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/uopensail/ulib/zlog"
)

// Constants for metric status and configuration
const (
	StatusUnset int8 = iota
	StatusOK
	StatusERR
	StatusMISS
	StatusMax
)

// Configuration constants for metrics collection
const (
	MetricsChannelBuff        = 30000
	MetricsCollectInterval    = 30
	MetricsGatherDBBuffSize   = 300
	MetricsCostBucketSize     = 100
	MetricsCostBucketMaxValue = 200
)

// MetricsItem represents a single metric measurement with timing and status information
type MetricsItem struct {
	Name       string
	Counter    int
	CostTime   int64
	Status     int8
	sampleRate float32
}

// metricsItems is a collection of MetricsItem that implements sort.Interface
type metricsItems []*MetricsItem

// Len returns the number of items in the collection
func (mi metricsItems) Len() int { return len(mi) }

// Less compares two items by their CostTime for sorting
func (mi metricsItems) Less(i, j int) bool {
	return mi[i].CostTime < mi[j].CostTime
}

// Swap exchanges two items in the collection
func (mi metricsItems) Swap(i, j int) {
	mi[i], mi[j] = mi[j], mi[i]
}

// GetP90 returns the 90th percentile item from the sorted collection
func (mi metricsItems) GetP90() *MetricsItem {
	if len(mi) == 0 {
		return nil
	}
	n := len(mi)
	index := int(float32(n) * 0.9)
	if index >= n {
		index = n - 1
	}
	return mi[index]
}

// GetP95 returns the 95th percentile item from the sorted collection
func (mi metricsItems) GetP95() *MetricsItem {
	if len(mi) == 0 {
		return nil
	}
	n := len(mi)
	index := int(float32(n) * 0.95)
	if index >= n {
		index = n - 1
	}
	return mi[index]
}

// GetP99 returns the 99th percentile item from the sorted collection
func (mi metricsItems) GetP99() *MetricsItem {
	if len(mi) == 0 {
		return nil
	}
	n := len(mi)
	index := int(float32(n) * 0.99)
	if index >= n {
		index = n - 1
	}
	return mi[index]
}

// SampleRate sets the sampling rate for the metric item
func (mi *MetricsItem) SampleRate(rate float32) *MetricsItem {
	mi.sampleRate = rate
	return mi
}

// MarkOk sets the status to OK
func (mi *MetricsItem) MarkOk() *MetricsItem {
	mi.Status = StatusOK
	return mi
}

// MarkMiss sets the status to MISS
func (mi *MetricsItem) MarkMiss() *MetricsItem {
	mi.Status = StatusMISS
	return mi
}

// MarkErr sets the status to ERR
func (mi *MetricsItem) MarkErr() *MetricsItem {
	mi.Status = StatusERR
	return mi
}

// SetCounter sets the counter value for the metric item
func (mi *MetricsItem) SetCounter(counter int) *MetricsItem {
	mi.Counter = counter
	return mi
}

// End finalizes the metric measurement and pushes it to the global metrics instance
func (mi *MetricsItem) End() {
	if mi.sampleRate >= 1 || rand.Float32() < mi.sampleRate {
		mi.CostTime = time.Now().UnixNano() - mi.CostTime
		GlobalmetricsIns.Push(mi)
	}
}

// MetricsInstance manages the collection and aggregation of metrics
type MetricsInstance struct {
	mu               sync.RWMutex
	metricsMap       map[string][StatusMax]*MetricsGather
	channel          chan *MetricsItem
	lastTickerTime   int64
	lastCollectTime  int64
	lastMetricsInfos []MetricsInfo
}

// GlobalmetricsIns is the global instance for metrics collection
var GlobalmetricsIns = &MetricsInstance{
	metricsMap: make(map[string][StatusMax]*MetricsGather),
	channel:    make(chan *MetricsItem, MetricsChannelBuff),
}

func init() {
	go GlobalmetricsIns.startLoop()
}

// NewStat creates a new MetricsItem with default sampling rate
func NewStat(name string) *MetricsItem {
	return &MetricsItem{
		Name:       name,
		Status:     StatusOK,
		CostTime:   time.Now().UnixNano(),
		sampleRate: 1.0,
	}
}

// NewCounterStat creates a new MetricsItem with a specific counter value
func NewCounterStat(name string, counter int) *MetricsItem {
	return &MetricsItem{
		Name:     name,
		Status:   StatusOK,
		CostTime: time.Now().UnixNano(),
		Counter:  counter,
	}
}

// GetMetricsInfo returns the latest collected metrics information
func (mInstance *MetricsInstance) GetMetricsInfo() []MetricsInfo {
	mInstance.mu.RLock()
	defer mInstance.mu.RUnlock()

	mInstance.lastCollectTime = time.Now().UnixNano()
	return mInstance.lastMetricsInfos
}

// startLoop begins the main processing loop for metrics collection
func (mInstance *MetricsInstance) startLoop() {
	ticker := time.NewTicker(time.Second * time.Duration(MetricsCollectInterval))
	defer ticker.Stop()

	for {
		select {
		case cItem := <-mInstance.channel:
			mInstance.AddItem(cItem)
		case <-ticker.C:
			mInstance.tickerCollectInfos()
		}
	}
}

// calcBucketIndex calculates the appropriate bucket index for a given cost in milliseconds
func calcBucketIndex(ms uint32) int {
	if ms <= 0 {
		return 0
	}
	if ms <= 50 {
		return int(ms) - 1
	}
	if ms > MetricsCostBucketMaxValue {
		return MetricsCostBucketSize
	}
	return int((ms-51)/3 + 50)
}

// tickerCollectInfos collects and processes metrics at regular intervals
func (mInstance *MetricsInstance) tickerCollectInfos() {
	mInstance.mu.Lock()
	defer mInstance.mu.Unlock()

	if len(mInstance.metricsMap) == 0 {
		return
	}

	leCosts := []uint32{3, 6, 10, 15, 20, 25, 35, 45, 59, 98, 149, 200, 999}
	metricsInfos := make([]MetricsInfo, 0, len(mInstance.metricsMap)*int(StatusMax))
	now := time.Now().UnixNano()
	collectInterval := float32((now - mInstance.lastTickerTime) / int64(time.Second))

	for name, statusGathers := range mInstance.metricsMap {
		for status, gather := range statusGathers {
			if gather != nil && gather.total > 0 {
				info := mInstance.collectGatherInfo(name, int8(status), gather, collectInterval, leCosts)
				if info != nil {
					metricsInfos = append(metricsInfos, *info)
				}
			}
		}
	}

	if now-mInstance.lastCollectTime > 3*MetricsCollectInterval*int64(time.Second) {
		go printStat(metricsInfos)
	}

	mInstance.lastMetricsInfos = metricsInfos
	mInstance.metricsMap = make(map[string][StatusMax]*MetricsGather, MetricsGatherDBBuffSize)
	mInstance.lastTickerTime = now
}

// collectGatherInfo collects information from a single MetricsGather
func (mInstance *MetricsInstance) collectGatherInfo(name string, status int8, gather *MetricsGather, collectInterval float32, leCosts []uint32) *MetricsInfo {
	if gather.total == 0 {
		return nil
	}

	// Sort items for percentile calculations
	sort.Sort(gather.items)

	info := &MetricsInfo{
		Total:      gather.total,
		AvgCounter: float32(gather.counter) / float32(gather.total),
		Name:       name,
		MaxCost:    gather.maxCost / float64(time.Millisecond),
		AvgCost:    float64(gather.sumCost) / float64(gather.total) / float64(time.Millisecond),
		QPS:        float32(gather.total) / collectInterval,
	}

	// Calculate percentiles with nil checks
	if p90 := gather.items.GetP90(); p90 != nil {
		info.P90Cost = float64(p90.CostTime) / float64(time.Millisecond)
	}
	if p95 := gather.items.GetP95(); p95 != nil {
		info.P95Cost = float64(p95.CostTime) / float64(time.Millisecond)
	}
	if p99 := gather.items.GetP99(); p99 != nil {
		info.P99Cost = float64(p99.CostTime) / float64(time.Millisecond)
	}

	// Calculate cost bucket distribution
	info.CostBucket = make(MapFI, len(leCosts))
	bucketSum := uint64(0)
	cbi := 0

	for _, leCost := range leCosts {
		endBucket := calcBucketIndex(leCost)
		for ; cbi <= endBucket && cbi < len(gather.costBucket); cbi++ {
			bucketSum += uint64(gather.costBucket[cbi])
		}
		info.CostBucket[float64(leCost)] = bucketSum
	}

	// Set status string
	switch status {
	case StatusOK:
		info.Status = "OK"
	case StatusMISS:
		info.Status = "MISS"
	default:
		info.Status = "ERR"
	}

	return info
}

// Push adds a metric item to the processing channel
func (mInstance *MetricsInstance) Push(mi *MetricsItem) {
	select {
	case mInstance.channel <- mi:
	default:
		// If channel is full, pop one item and try again
		select {
		case <-mInstance.channel:
			zlog.LOG.Warn("Metrics channel full, dropped one item")
		default:
		}
		select {
		case mInstance.channel <- mi:
		default:
			zlog.LOG.Warn("Metrics channel is full, dropped metric item")
		}
	}
}

// AddItem processes a single metric item and adds it to the appropriate gather
func (mInstance *MetricsInstance) AddItem(mi *MetricsItem) {
	mInstance.mu.Lock()
	defer mInstance.mu.Unlock()

	gathers, ok := mInstance.metricsMap[mi.Name]
	if !ok {
		gathers = [StatusMax]*MetricsGather{}
		mInstance.metricsMap[mi.Name] = gathers
	}

	gather := gathers[mi.Status]
	if gather == nil {
		gather = newMetricsGather()
		gathers[mi.Status] = gather
	}

	gather.push(mi)
	gather.total++
	gather.sumCost += mi.CostTime
	gather.maxCost = math.Max(float64(mi.CostTime), gather.maxCost)
	gather.counter += mi.Counter

	// Calculate bucket index and increment counter
	costMs := mi.CostTime / int64(time.Millisecond)
	if costMs < 0 {
		costMs = 0
	}
	bIndex := calcBucketIndex(uint32(costMs))
	if bIndex >= 0 && bIndex < len(gather.costBucket) {
		gather.costBucket[bIndex]++
	}

	mInstance.metricsMap[mi.Name] = gathers
}

// MetricsGather aggregates multiple metric items of the same type
type MetricsGather struct {
	maxCost    float64
	sumCost    int64
	total      int
	counter    int
	items      metricsItems
	costBucket [MetricsCostBucketSize + 1]uint32
}

// newMetricsGather creates a new MetricsGather instance
func newMetricsGather() *MetricsGather {
	return &MetricsGather{
		items: make([]*MetricsItem, 0, 1000),
	}
}

// push adds a metric item to the gather
func (gather *MetricsGather) push(item *MetricsItem) {
	gather.items = append(gather.items, item)
}

// MetricsInfo contains aggregated metric information for export
type MetricsInfo struct {
	Name       string  `json:"name"`
	Status     string  `json:"status"`
	QPS        float32 `json:"qps"`
	Total      int     `json:"total"`
	AvgCost    float64 `json:"avg_cost"`
	P90Cost    float64 `json:"p90_cost"`
	P95Cost    float64 `json:"p95_cost"`
	P99Cost    float64 `json:"p99_cost"`
	MaxCost    float64 `json:"max_cost"`
	AvgCounter float32 `json:"avg_counter"`
	Counter    float32 `json:"counter"`
	CostBucket MapFI   `json:"cost_bucket"`
}

// MapFI is a map of float64 to uint64 used for cost bucket distribution
type MapFI map[float64]uint64

// MarshalJSON implements custom JSON marshaling for MapFI
func (mi MapFI) MarshalJSON() ([]byte, error) {
	if len(mi) == 0 {
		return []byte("{}"), nil
	}

	bf := bytes.Buffer{}
	bf.WriteByte('{')
	first := true

	// Sort keys for consistent output
	keys := make([]float64, 0, len(mi))
	for k := range mi {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	for _, k := range keys {
		v := mi[k]
		if !first {
			bf.WriteByte(',')
		}
		bf.WriteString(strconv.FormatFloat(k, 'f', -1, 64))
		bf.WriteByte(':')
		bf.WriteString(strconv.FormatUint(v, 10))
		first = false
	}
	bf.WriteByte('}')

	return bf.Bytes(), nil
}

// String returns the JSON string representation of MetricsInfo
func (info *MetricsInfo) String() string {
	data, err := json.Marshal(info)
	if err != nil {
		return "{}"
	}
	return string(data)
}

// printStat logs the collected metrics information
func printStat(metricsInfo []MetricsInfo) {
	zlog.LOG.Info("Metrics collection summary", zap.Int("count", len(metricsInfo)))
	for i, info := range metricsInfo {
		zlog.LOG.Info("Metric stat",
			zap.Int("index", i),
			zap.String("name", info.Name),
			zap.String("status", info.Status),
			zap.Float32("qps", info.QPS),
			zap.Float64("avg_cost", info.AvgCost),
		)
	}
}
