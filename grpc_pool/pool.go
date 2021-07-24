package grpc_pool

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/uopensail/ulib/commonconfig"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type grpcConn struct {
	conn    *grpc.ClientConn
	status  bool
	url     string
	timeout int
}

type Pool struct {
	Config  *commonconfig.GRPCClientConfig
	Clients []*grpcConn
	Count   int
}

func check(conns []*grpcConn) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		for i := 0; i < len(conns); i++ {
			if conns[i].status {
				health, err := grpc_health_v1.NewHealthClient(conns[i].conn).Check(context.Background(),
					&grpc_health_v1.HealthCheckRequest{Service: ""})
				if err != nil || health.Status != grpc_health_v1.HealthCheckResponse_SERVING {
					conns[i].status = false
				}
			} else {
				timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(conns[i].timeout))
				conn, err := grpc.DialContext(timeoutCtx,
					conns[i].url,
					grpc.WithDefaultServiceConfig(`{loadBalancingConfig:[{"round_robin":{}}]}`),
					grpc.WithInsecure(),
					grpc.WithBlock(),
				)
				if err != nil {
					conns[i].status = false
				} else {
					conns[i].conn, conns[i].status = conn, true
				}
				cancel()
			}

		}
	}
}

func NewPool(cfg *commonconfig.GRPCClientConfig) *Pool {
	if !strings.HasPrefix(cfg.Url, "grpc:///") {
		panic(fmt.Errorf("prefix error: %s", cfg.Url))
	}
	urls := strings.Split(cfg.Url[8:], ",")
	maxConn := 3
	if cfg.MaxConn > 0 {
		maxConn = cfg.MaxConn
	}

	timeout := 30
	if cfg.DialTimeout > 0 {
		timeout = cfg.DialTimeout
	}
	pool := &Pool{
		Config:  cfg,
		Clients: make([]*grpcConn, maxConn*len(urls)),
		Count:   maxConn * len(urls),
	}

	for i := 0; i < len(urls); i++ {
		for j := 0; j < maxConn; j++ {
			timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(timeout))
			conn, err := grpc.DialContext(timeoutCtx,
				urls[i],
				grpc.WithInsecure(),
				grpc.WithBlock(),
			)
			if err != nil {
				pool.Clients[i*maxConn+j].status = false
			} else {
				pool.Clients[i*maxConn+j].conn, pool.Clients[i*maxConn+j].status = conn, true
			}
			cancel()
		}
	}
	if cfg.HealthCheck {
		go check(pool.Clients)
	}
	return pool
}

func (pool *Pool) GetConn() *grpc.ClientConn {
	index := rand.Intn(pool.Count)
	tmpIndex := 0
	for j := 0; j < pool.Count; j++ {
		tmpIndex = (index + j) % pool.Count
		if pool.Clients[tmpIndex].status {
			return pool.Clients[tmpIndex].conn
		}
	}
	return nil
}
