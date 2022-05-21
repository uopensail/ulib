package logger

import (
	"strconv"
	"testing"

	"github.com/uopensail/ulib/commonconfig"
)

func TestSDK(t *testing.T) {

	cfg := LogSDKConfig{
		RegisterDiscoveryConfig: commonconfig.RegisterDiscoveryConfig{
			EtcdConfig: commonconfig.EtcdConfig{
				Endpoints: []string{"127.0.0.1:2181"},
				Name:      "pangu_server",
			},
		},
	}
	Init(cfg)
	for i := 0; i < 100000; i++ {
		log := &Log{
			ProductId: "test",
			UserId:    "user_" + strconv.Itoa(i),
		}
		Push(log)
	}
	select {}
}
