package logger

import (
	"context"
	"time"

	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/utils"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

const defaultChannelSize = 1000
const defaultBufferSize = 1000

type LogSDKConfig struct {
	commonconfig.RegisterDiscoveryConfig `json:"discovery" toml:"discovery"`
	ChannelSize                          int `json:"channel_size" toml:"channel_size"`
	BufferSize                           int `json:"buffer_size" toml:"buffer_size"`
	FlushInterval                        int `json:"flush_interval" toml:"flush_interval"`
}

type SDK struct {
	conn    *grpc.ClientConn
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
	conn, _ := utils.NewKratosGrpcConn(cfg.RegisterDiscoveryConfig)

	globalLoggerSDK = &SDK{
		conn:    conn,
		channel: make(chan *Log, channelSize),
	}

	flushInterval := 3
	if cfg.FlushInterval > 0 {
		flushInterval = cfg.FlushInterval
	}
	go globalLoggerSDK.Flush(bufferSize, flushInterval)
}

func Push(log *Log) {

	if globalLoggerSDK == nil {
		return
	}
	select {
	case globalLoggerSDK.channel <- log:
		break
	default:
		zlog.LOG.Warn("logsdk.full")
		break
	}
}

func (sdk *SDK) write(logs []*Log) {
	stat := prome.NewStat("Logger.SDK.write")
	defer stat.End()

	if sdk.conn == nil {
		stat.MarkErr()
		return
	}
	req := &BatchRequest{
		Logs: logs,
	}
	_, err := NewLogServerClient(sdk.conn).Batch(context.Background(), req)
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
