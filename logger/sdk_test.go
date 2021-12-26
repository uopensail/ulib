package logger

import (
	"strconv"
	"testing"

	"github.com/uopensail/ulib/commonconfig"
)

func TestSDK(t *testing.T) {

	cfg := LogSDKConfig{
		PanguServer: commonconfig.PanguConfig{
			ZKHosts:     []string{"127.0.0.1:2181"},
			ServiceName: "pangu_server",
		},
	}
	Init(cfg)
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
