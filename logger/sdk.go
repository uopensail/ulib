package logger

import (
	"context"
	"time"

	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/grpc_util"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
)

const defaultChannelSize = 1000
const defaultBufferSize = 1000

type LogSDKConfig struct {
	commonconfig.GRPCClientConfig `json:"grpc" toml:"grpc"`
	ChannelSize                   int `json:"channel_size" toml:"channel_size"`
	BufferSize                    int `json:"buffer_size" toml:"buffer_size"`
	FlushInterval                 int `json:"flush_interval" toml:"flush_interval"`
}

type SDK struct {
	pool    *grpc_util.Pool
	channel chan *Log
}

func Init(cfg LogSDKConfig) {
	channelSize := defaultChannelSize
	if cfg.ChannelSize > 0 {
		channelSize = cfg.ChannelSize
	}
	bufferSize := defaultBufferSize
	if cfg.BufferSize > 0 {
		bufferSize = cfg.BufferSize
	}

	globalLoggerSDK = &SDK{
		pool:    grpc_util.NewPool(&cfg.GRPCClientConfig),
		channel: make(chan *Log, channelSize),
	}

	flushInterval := 3
	if cfg.FlushInterval > 0 {
		flushInterval = cfg.FlushInterval
	}
	go globalLoggerSDK.Flush(bufferSize, flushInterval)
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

func (sdk *SDK) Flush(bufferSize int, flushInterval int) {
	ticker := time.NewTicker(time.Duration(flushInterval) * time.Second)
	defer ticker.Stop()
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
