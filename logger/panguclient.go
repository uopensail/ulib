package logger

import (
	"fmt"
	"sync"

	"github.com/uopensail/ulib/zlog"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

type panguClient struct {
	cli         LogServerClient
	serviceName string
	regaddr     []string
}

const (
	regPath = "/recommend/pangu/"
	scheme  = "pangu"
	service = "panguservice"
)

var once sync.Once
var pangucli *panguClient

func GetPanguClientInstance() *panguClient {
	return pangucli
}

func InitPanguClient(serviceName string, regaddr []string) *panguClient {
	once.Do(func() {
		if nil == pangucli {
			fmt.Printf("init pangu client instance\n")
			pangucli = new(panguClient)
			pangucli.serviceName = serviceName
			pangucli.regaddr = regaddr
			resolver.Register(&panguResolverBuilder{
				panguScheme:      scheme,
				panguServiceName: service,
				panguWatchPrefix: regPath + serviceName,
				panguzkHost:      regaddr,
			})
			pangucli.initPanguClient()
		}
	})

	return pangucli
}

func (c *panguClient) initPanguClient() *panguClient {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", scheme, service),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)

	if nil != err {
		zlog.SLOG.Error(err)
		panic(err)
	}

	c.cli = NewLogServerClient(conn)
	return c
}

func (c *panguClient) GetRealClient() LogServerClient {
	return pangucli.cli
}
