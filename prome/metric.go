package prome

import (
	"encoding/json"
	"math"
	"math/rand"
	"time"

	"go.uber.org/zap"

	"github.com/uopensail/ulib/zlog"
)

const (
	StatusOK int8 = iota
	StatusERR
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
	name       string
	counter    int
	costTime   int64
	status     int8
	sampleRate float32
}

func (mi *MetricsItem) SampleRate(rate float32) *MetricsItem {
	mi.sampleRate = rate
	return mi
}

func (mi *MetricsItem) MarkOk() *MetricsItem {
	mi.status = StatusOK
	return mi
}

func (mi *MetricsItem) MarkErr() *MetricsItem {
	mi.status = StatusERR
	return mi
}

func (mi *MetricsItem) SetCounter(counter int) *MetricsItem {
	mi.counter = counter
	return mi
}

func (mi *MetricsItem) End() {
	if mi.sampleRate >= 1 || rand.Float32() < mi.sampleRate {
		mi.costTime = time.Now().UnixNano() - mi.costTime
		select {
		case metricsInstance.channel <- mi:
		default:
			zlog.LOG.Warn("Metrics channel is full")
		}
	}
}

type MetricsInstance struct {
	metricsMap       map[string][StatusMax]*MetricsGather
	channel          chan *MetricsItem
	lastTickerTime   int64
	lastCollectTime  int64
	lastMetricsInfos []MetricsInfo
}

var metricsInstance = &MetricsInstance{
	make(map[string][StatusMax]*MetricsGather),
	make(chan *MetricsItem, MetricsChannelBuff),
	time.Now().UnixNano(),
	time.Now().UnixNano(),
	make([]MetricsInfo, 0, 0),
}

func init() {
	go metricsInstance.startLoop()
}

func NewStat(name string) *MetricsItem {
	return &MetricsItem{
		name:       name,
		costTime:   time.Now().UnixNano(),
		sampleRate: 1.0,
	}
}

func NewCounterStat(name string, counter int) *MetricsItem {
	return &MetricsItem{
		name:     name,
		costTime: time.Now().UnixNano(),
		counter:  counter,
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
	metricsInfos := make([]MetricsInfo, len(mInstance.metricsMap)*int(StatusMax))
	index := 0
	now := time.Now().UnixNano()
	collectInterval := float32((now - mInstance.lastCollectTime) / int64(time.Second))
	for k, v := range mInstance.metricsMap {
		for j := int8(0); j < StatusMax; j++ {
			gather := v[j]
			if gather != nil && gather.total > 0 {
				metricsInfos[index].Total = gather.total
				metricsInfos[index].AvgCounter = float32(gather.counter) / float32(gather.total)
				metricsInfos[index].Name = k
				metricsInfos[index].MaxCost = gather.maxCost / float64(time.Millisecond)
				metricsInfos[index].AvgCost = float64(gather.sumCost) / float64(gather.total) / float64(time.Millisecond)
				metricsInfos[index].QPS = float32(gather.total) / collectInterval

				//计算cost bucket
				leCosts := []uint32{3, 6, 10, 15, 20, 25, 35, 45, 59, 98, 149, 200, 999}
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

func (mInstance *MetricsInstance) AddItem(mi *MetricsItem) {

	gathers, ok := mInstance.metricsMap[mi.name]
	if ok == false {
		var gs [StatusMax]*MetricsGather
		mInstance.metricsMap[mi.name] = gs
		gathers = gs
	}
	gather := gathers[mi.status]
	if gather == nil {
		gather = &MetricsGather{}
		gathers[mi.status] = gather
	}
	gather.total++
	gather.sumCost += mi.costTime
	gather.maxCost = math.Max(float64(mi.costTime), gather.maxCost)
	gather.counter += mi.counter
	//cost ms
	costMs := mi.costTime / int64(time.Millisecond)
	bIndex := calcBucketIndex(uint32(costMs))
	gather.costBucket[bIndex]++
	metricsInstance.metricsMap[mi.name] = gathers
}

type MetricsGather struct {
	maxCost    float64
	sumCost    int64
	total      int
	counter    int
	costBucket [MetricsCostBucketSize + 1]uint32
}

type MetricsInfo struct {
	Name       string             `json:"name"`
	Status     string             `json:"status"`
	QPS        float32            `json:"qps"`
	Total      int                `json:"total"`
	AvgCost    float64            `json:"avg_cost"`
	MaxCost    float64            `json:"max_cost"`
	AvgCounter float32            `json:"avg_counter"`
	Counter    float32            `json:"counter"`
	CostBucket map[float64]uint64 `json:"-"`
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
