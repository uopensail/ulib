package logger

import (
	"github.com/uopensail/ulib/commonconfig"
	"strconv"
	"testing"
)

func TestSDK(t *testing.T) {
	cfg := commonconfig.GRPCClientConfig{
		Url:            "grpc:///127.0.0.1:9527",
		DialTimeout:    100,
		RequestTimeout: 100,
		InitConn:       2,
		MaxConn:        2,
		HealthCheck:    true,
	}
	Init(&cfg)
	for i := 0; i < 100000; i++ {
		log := &Log{
			ProductId: "test",
			UserId:    "user_" + strconv.Itoa(i),
			Event:     "test",
		}
		Push(log)
	}
	select {}
}
