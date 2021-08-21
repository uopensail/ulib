package grpc_util

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"

	"google.golang.org/grpc/health/grpc_health_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&LocalResolverBuilder{})
}

type localResolver struct {
	cc        resolver.ClientConn
	healthCli map[string]*grpc.ClientConn
}

func (*localResolver) ResolveNow(options resolver.ResolveNowOptions) {}
func (*localResolver) Close()                                        {}

func (r *localResolver) watcher(urls []string) {
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	for {
		<-ticker.C
		r.update(urls)
	}
}
func (r *localResolver) update(urls []string) {
	for i := 0; i < len(urls); i++ {
		url := urls[i]
		if c, ok := r.healthCli[url]; !ok || c == nil {
			r.healthCli[url] = newHealthClient(url)
		}
	}
	address := make([]resolver.Address, 0, len(urls))
	for i := 0; i < len(urls); i++ {
		url := urls[i]
		c := r.healthCli[url]
		if healthCheck(url, c) == nil {
			address = append(address, resolver.Address{Addr: url})
		} else {
			if c != nil {
				c.Close()
			}
			r.healthCli[url] = nil
		}
	}
	r.cc.UpdateState(resolver.State{
		Addresses: address,
	})
}
func newHealthClient(url string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	c, err := grpc.DialContext(ctx, url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil
	}
	return c
}
func healthCheck(url string, c *grpc.ClientConn) error {
	if c == nil {
		return errors.New("connection nil")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(200)*time.Millisecond)
	defer cancel()
	cli := grpc_health_v1.NewHealthClient(c)
	response, err := cli.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		return err
	}

	if response.Status != grpc_health_v1.HealthCheckResponse_SERVING {
		return errors.New(response.Status.String())
	}
	zlog.LOG.Info("grpc.check.success", zap.String("url", url))
	return nil
}

type LocalResolverBuilder struct{}

func (*LocalResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver,
	error) {
	urlConfig := strings.TrimPrefix(target.Endpoint, "grpc://")
	urls := strings.Split(strings.Trim(urlConfig, " "), ",")
	r := localResolver{
		cc:        cc,
		healthCli: make(map[string]*grpc.ClientConn, len(urls)),
	}
	r.update(urls)
	go r.watcher(urls)
	return &r, nil
}
func (*LocalResolverBuilder) Scheme() string {
	return "grpc"
}
