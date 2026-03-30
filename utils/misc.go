package utils

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	etcd "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/zlog"
	etcdclient "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

// NewKratosGrpcConn creates an insecure gRPC client connection to a service
// discovered via etcd using the Kratos registry. rdConf supplies the etcd
// endpoints and the target service name.
func NewKratosGrpcConn(rdConf commonconfig.RegisterDiscoveryConfig) (*grpc.ClientConn, error) {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: rdConf.EtcdConfig.Endpoints,
	})
	if err != nil {
		zlog.LOG.Error("failed to create etcd client", zap.Error(err))
		return nil, err
	}
	r := etcd.New(client)
	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", rdConf.EtcdConfig.Name)),
		kgrpc.WithDiscovery(r),
	)
	if err != nil {
		zlog.LOG.Error("failed to dial gRPC endpoint", zap.Error(err))
		return nil, err
	}
	return conn, nil
}

// NewKratosHttpConn creates an HTTP client connection to a service discovered
// via etcd using the Kratos registry. rdConf supplies the etcd endpoints and
// the target service name.
func NewKratosHttpConn(rdConf commonconfig.RegisterDiscoveryConfig) (*khttp.Client, error) {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: rdConf.EtcdConfig.Endpoints,
	})
	if err != nil {
		zlog.LOG.Error("failed to create etcd client", zap.Error(err))
		return nil, err
	}
	r := etcd.New(client)
	conn, err := khttp.NewClient(
		context.Background(),
		khttp.WithEndpoint(fmt.Sprintf("discovery:///%s", rdConf.EtcdConfig.Name)),
		khttp.WithDiscovery(r),
		khttp.WithBlock(),
	)
	if err != nil {
		zlog.LOG.Error("failed to connect to HTTP endpoint", zap.Error(err))
		return nil, err
	}
	return conn, nil
}

// Reference is an atomic reference counter with an optional cleanup handler.
// Call Retain before handing out a reference, Release when done.
// Free (or LazyFree) blocks until the count reaches zero, then calls CloseHandler.
type Reference struct {
	CloseHandler   func()
	referenceCount int32
}

// Retain increments the reference count.
func (ref *Reference) Retain() {
	atomic.AddInt32(&ref.referenceCount, 1)
}

// Release decrements the reference count.
func (ref *Reference) Release() {
	atomic.AddInt32(&ref.referenceCount, -1)
}

// Free blocks until the reference count reaches zero, then invokes CloseHandler.
func (ref *Reference) Free() {
	for atomic.LoadInt32(&ref.referenceCount) > 0 {
		time.Sleep(time.Millisecond * 10)
	}
	if ref.CloseHandler != nil {
		ref.CloseHandler()
	}
}

// LazyFree waits lazySecond seconds (default 3 if ≤ 0), then behaves like Free.
// The wait and cleanup run in a new goroutine so the caller is not blocked.
func (ref *Reference) LazyFree(lazySecond int) {
	go func() {
		if lazySecond <= 0 {
			lazySecond = 3
		}
		time.Sleep(time.Second * time.Duration(lazySecond))
		for atomic.LoadInt32(&ref.referenceCount) > 0 {
			time.Sleep(time.Millisecond * 10)
		}
		if ref.CloseHandler != nil {
			ref.CloseHandler()
		}
	}()
}
