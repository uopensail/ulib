package prome

import (
	"bytes"
	"encoding/json"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/uopensail/ulib/zlog"
)

const (
	StatusUnset int8 = iota
	StatusOK
	StatusERR
	StatusMISS
	StatusMax
)

const (
	MetricsChannelBuff        = 30000
	MetricsCollectInterval    = 30
	MetricsGatherDBBuffSize   = 300
	MetricsCostBucketSize     = 100
	MetricsCostBucketMaxValue = 200
)

type MetricsItem struct {
	Name       string
	Counter    int
	CostTime   int64
	Status     int8
	sampleRate float32
}

type metricsItems []*MetricsItem

func (mi metricsItems) Len() int { return len(mi) }
func (mi metricsItems) Less(i, j int) bool {
	return mi[i].CostTime < mi[j].CostTime
}
func (mi metricsItems) Swap(i, j int) {
	mi[i], mi[j] = mi[j], mi[i]
}

func (mi metricsItems) GetP90() *MetricsItem {
	n := len(mi)
	index := int(float32(n) * 0.9)
	if index >= n {
		index = n - 1
	}
	return mi[index]
}

func (mi metricsItems) GetP95() *MetricsItem {
	n := len(mi)
	index := int(float32(n) * 0.95)
	if index >= n {
		index = n - 1
	}
	return mi[index]
}
func (mi metricsItems) GetP99() *MetricsItem {
	n := len(mi)
	index := int(float32(n) * 0.99)
	if index >= n {
		index = n - 1
	}
	return mi[index]
}

func (mi *MetricsItem) SampleRate(rate float32) *MetricsItem {
	mi.sampleRate = rate
	return mi
}

func (mi *MetricsItem) MarkOk() *MetricsItem {
	mi.Status = StatusOK
	return mi
}
func (mi *MetricsItem) MarkMiss() *MetricsItem {
	mi.Status = StatusMISS
	return mi
}

func (mi *MetricsItem) MarkErr() *MetricsItem {
	mi.Status = StatusERR
	return mi
}

func (mi *MetricsItem) SetCounter(counter int) *MetricsItem {
	mi.Counter = counter
	return mi
}

func (mi *MetricsItem) End() {
	if mi.sampleRate >= 1 || rand.Float32() < mi.sampleRate {
		mi.CostTime = time.Now().UnixNano() - mi.CostTime
		GlobalmetricsIns.Push(mi)
	}
}

type MetricsInstance struct {
	metricsMap       map[string][StatusMax]*MetricsGather
	channel          chan *MetricsItem
	lastTickerTime   int64
	lastCollectTime  int64
	lastMetricsInfos []MetricsInfo
}

var GlobalmetricsIns = &MetricsInstance{
	make(map[string][StatusMax]*MetricsGather),
	make(chan *MetricsItem, MetricsChannelBuff),
	time.Now().UnixNano(),
	time.Now().UnixNano(),
	make([]MetricsInfo, 0, 0),
}

func init() {
	go GlobalmetricsIns.startLoop()
}

func NewStat(name string) *MetricsItem {
	return &MetricsItem{
		Name:       name,
		Status:     StatusOK,
		CostTime:   time.Now().UnixNano(),
		sampleRate: 1.0,
	}
}

func NewCounterStat(name string, counter int) *MetricsItem {
	return &MetricsItem{
		Name:     name,
		Status:   StatusOK,
		CostTime: time.Now().UnixNano(),
		Counter:  counter,
	}
}

func (mInstance *MetricsInstance) GetMetricsInfo() []MetricsInfo {
	mInstance.lastCollectTime = time.Now().UnixNano()
	return mInstance.lastMetricsInfos
}

func (mInstance *MetricsInstance) startLoop() {
	ticker := time.NewTicker(time.Second * time.Duration(MetricsCollectInterval))

	for {
		select {
		case cItem := <-mInstance.channel:
			mInstance.AddItem(cItem)
		case <-ticker.C:
			mInstance.tickerCollectInfos()
		}
	}

}

func calcBucketIndex(ms uint32) int {
	var bIndex int
	if ms <= 50 {
		bIndex = int(ms) - 1
		if bIndex < 0 {
			bIndex = 0
		}
	} else if ms > MetricsCostBucketMaxValue {
		bIndex = MetricsCostBucketSize
	} else {
		bIndex = int((ms-51)/3 + 50)
	}
	return bIndex
}

func (mInstance *MetricsInstance) tickerCollectInfos() {
	if mInstance.metricsMap == nil || len(mInstance.metricsMap) <= 0 {
		return
	}
	leCosts := []uint32{3, 6, 10, 15, 20, 25, 35, 45, 59, 98, 149, 200, 999}
	metricsInfos := make([]MetricsInfo, len(mInstance.metricsMap)*int(StatusMax))
	index := 0
	now := time.Now().UnixNano()
	collectInterval := float32((now - mInstance.lastCollectTime) / int64(time.Second))
	for k, v := range mInstance.metricsMap {
		for j := int8(0); j < StatusMax; j++ {
			gather := v[j]
			if gather != nil && gather.total > 0 {
				sort.Sort(gather.items)
				metricsInfos[index].Total = gather.total
				metricsInfos[index].AvgCounter = float32(gather.counter) / float32(gather.total)
				metricsInfos[index].Name = k
				metricsInfos[index].MaxCost = gather.maxCost / float64(time.Millisecond)
				metricsInfos[index].AvgCost = float64(gather.sumCost) / float64(gather.total) / float64(time.Millisecond)
				metricsInfos[index].QPS = float32(gather.total) / collectInterval
				metricsInfos[index].P90Cost = float64(gather.items.GetP90().CostTime) / float64(time.Millisecond)
				metricsInfos[index].P95Cost = float64(gather.items.GetP95().CostTime) / float64(time.Millisecond)
				metricsInfos[index].P99Cost = float64(gather.items.GetP99().CostTime) / float64(time.Millisecond)
				// calculate cost bucket
				costBucket := make(map[float64]uint64, len(leCosts))
				cbi := 0
				bucketSum := uint64(0)
				for bi := 0; bi < len(leCosts); bi++ {
					leCost := leCosts[bi]
					endBucket := calcBucketIndex(leCost)
					for ; cbi <= endBucket; cbi++ {
						bucketSum += uint64(gather.costBucket[cbi])
					}
					costBucket[float64(leCost)] = bucketSum
				}
				metricsInfos[index].CostBucket = costBucket
				if j == StatusOK {
					metricsInfos[index].Status = "OK"
				} else if j == StatusMISS {
					metricsInfos[index].Status = "MISS"
				} else {
					metricsInfos[index].Status = "ERR"
				}
				index++
			}
		}
	}

	infos := metricsInfos[:index]
	if now-mInstance.lastCollectTime > 3*MetricsCollectInterval*int64(time.Second) {
		go printStat(infos)
	}
	mInstance.lastMetricsInfos = infos
	mInstance.metricsMap = make(map[string][StatusMax]*MetricsGather, MetricsGatherDBBuffSize)
	mInstance.lastTickerTime = now

}
func (mInstance *MetricsInstance) Push(mi *MetricsItem) {
	select {
	case mInstance.channel <- mi:
	default:
		// if channel is full, pop front, and try to push again
		<-mInstance.channel
		select {
		case mInstance.channel <- mi:
		default:
			zlog.LOG.Warn("Metrics channel is full")
		}
	}
}
func (mInstance *MetricsInstance) AddItem(mi *MetricsItem) {
	gathers, ok := mInstance.metricsMap[mi.Name]
	if !ok {
		var gs [StatusMax]*MetricsGather
		mInstance.metricsMap[mi.Name] = gs
		gathers = gs
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
	//cost ms
	costMs := mi.CostTime / int64(time.Millisecond)
	bIndex := calcBucketIndex(uint32(costMs))
	gather.costBucket[bIndex]++
	GlobalmetricsIns.metricsMap[mi.Name] = gathers
}

type MetricsGather struct {
	maxCost    float64
	sumCost    int64
	total      int
	counter    int
	items      metricsItems
	costBucket [MetricsCostBucketSize + 1]uint32
}

func newMetricsGather() *MetricsGather {
	return &MetricsGather{
		items: make([]*MetricsItem, 0, 1000),
	}
}

func (gather *MetricsGather) push(item *MetricsItem) {
	gather.items = append(gather.items, item)
}

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

type MapFI map[float64]uint64

func (mi MapFI) MarshalJSON() ([]byte, error) {
	bf := bytes.Buffer{}

	for k, v := range mi {
		bf.WriteString(strconv.FormatInt(int64(k), 10))
		bf.WriteByte(':')
		bf.WriteString(strconv.FormatUint(v, 10))
		bf.WriteByte(',')
	}
	ret := bf.String()
	return json.Marshal(ret[:len(ret)-1])
}

func (info *MetricsInfo) String() string {
	data, _ := json.Marshal(info)
	return string(data)
}

func printStat(metricsInfo []MetricsInfo) {
	zlog.LOG.Info("prome info")
	for i := 0; i < len(metricsInfo); i++ {
		zlog.LOG.Info("prome: ", zap.String("stat", metricsInfo[i].String()))
	}
}
