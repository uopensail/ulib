package logger

import (
	"context"
	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/grpc_pool"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
	"time"
)

const channelSize = 1000
const bufferSize = 1000

type SDK struct {
	pool    *grpc_pool.Pool
	channel chan *Log
}

func Init(cfg *commonconfig.GRPCClientConfig) {
	globalLoggerSDK = &SDK{
		pool:    grpc_pool.NewPool(cfg),
		channel: make(chan *Log, channelSize),
	}
	go globalLoggerSDK.Flush()
}

func Push(log *Log) {
	select {
	case globalLoggerSDK.channel <- log:
		break
	default:
		prome.NewStat("Logger.SDK.Buffer.Full").End()
		break
	}
}

func (sdk *SDK) write(logs []*Log) {
	stat := prome.NewStat("Logger.SDK.write")
	defer stat.End()
	conn := sdk.pool.GetConn()
	if conn == nil {
		stat.MarkErr()
	}
	req := &BatchRequest{
		Logs: logs,
	}
	_, err := NewLogServerClient(conn).Batch(context.Background(), req)
	if err != nil {
		zlog.LOG.Error("Logger.SDK.write error", zap.Error(err))
		stat.MarkErr()
	}
	stat.SetCounter(len(logs))
}

func (sdk *SDK) Flush() {
	ticker := time.NewTicker(5 * time.Second)
	index := 0
	buffer := make([]*Log, bufferSize)
	var log *Log
	for {
		select {
		case <-ticker.C:
			if index > 0 {
				sdk.write(buffer[:index])
				index = 0
			}
			break
		case log = <-sdk.channel:
			if index == bufferSize {
				sdk.write(buffer)
				index = 0
			}
			buffer[index] = log
			index++
			break
		}
	}
}

var globalLoggerSDK *SDK
