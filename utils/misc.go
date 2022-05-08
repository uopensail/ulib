package utils

import (
	"context"
	"fmt"

	etcd "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/uopensail/ulib/commonconfig"
	"github.com/uopensail/ulib/zlog"
	etcdclient "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

func NewKratosGrpcConn(rdConf commonconfig.RegisterDiscoveryConfig) (*grpc.ClientConn, error) {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: rdConf.EtcdConfig.Endpoints,
	})
	if err != nil {
		zlog.LOG.Fatal("etcd error", zap.Error(err))
		return nil, err
	}
	r := etcd.New(client)
	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint(fmt.Sprintf("discovery:///%s", rdConf.EtcdConfig.Name)),
		kgrpc.WithDiscovery(r),
	)
	if err != nil {
		zlog.LOG.Fatal("etcd error", zap.Error(err))
		return nil, err
	}
	return conn, nil
}

func NewKratosHttpConn(rdConf commonconfig.RegisterDiscoveryConfig) (*khttp.Client, error) {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: rdConf.EtcdConfig.Endpoints,
	})
	if err != nil {
		zlog.LOG.Fatal("etcd error", zap.Error(err))
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
		zlog.LOG.Fatal("etcd error", zap.Error(err))
		return nil, err
	}
	return conn, nil
}
