package logger

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"google.golang.org/grpc/resolver"
)

type panguResolverBuilder struct {
	panguScheme      string
	panguServiceName string
	panguWatchPrefix string
	panguzkHost      []string
	addrs            []string
}

type panguResolver struct {
	target        resolver.Target
	cc            resolver.ClientConn
	addrsStore    map[string][]string
	isInitialized bool
	zkCli         *zk.Conn
	zkHost        []string
	prefix        string
}

type registUnit struct {
	Host string `json:"host"`
}

func (fr *panguResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &panguResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			fr.panguScheme: fr.addrs,
		},
		zkHost: fr.panguzkHost,
		prefix: fr.panguWatchPrefix,
	}
	r.initClient()

	return r, nil
}

func (fr *panguResolverBuilder) Scheme() string { return fr.panguScheme }

func (r *panguResolver) initClient() {
	fmt.Printf("init panguResolver\n")
	conn, _, err := zk.Connect(r.zkHost, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}

	r.zkCli = conn
	r.start()
	go r.watch()
}

func (r *panguResolver) start() {
	childs, _, err := r.zkCli.Children(r.prefix)
	if nil != err {
		fmt.Printf("err:%v prefix:%v\n", err.Error(), r.prefix)
		panic("not get panguaddrs")
	}

	addrs := make([]resolver.Address, 0, 10)
	for _, child := range childs {
		addrs = append(addrs, resolver.Address{Addr: child})
	}
	fmt.Printf("pangu addrs:%v\n", addrs)
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (r *panguResolver) watch() {

	for {
		childs, _, watcher, err := r.zkCli.ChildrenW(r.prefix)
		if nil != err {
			fmt.Printf("err:%v prefix:%v\n", err.Error(), r.prefix)
			continue
		}

		if len(childs) > 0 {
			addrs := make([]resolver.Address, 0, 10)
			for _, child := range childs {
				addrs = append(addrs, resolver.Address{Addr: child})
			}
			r.cc.UpdateState(resolver.State{Addresses: addrs})
		} else {
			addrs := make([]resolver.Address, 0, 10)
			r.cc.UpdateState(resolver.State{Addresses: addrs})
		}

		event := <-watcher
		if event.Err != nil {
			fmt.Printf("watch err:%v\n", event)
		} else {
			fmt.Printf("event:%v\n", event)
		}
	}
}

func (r *panguResolver) ResolveNow(o resolver.ResolveNowOptions) {
	fmt.Printf("ResolveNow called\n")
}

func (r *panguResolver) Close() {
	fmt.Printf("Close called\n")
}
