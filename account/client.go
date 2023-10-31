package v1

import (
	"context"
	"fmt"
	"os"
	"strings"
	sync "sync"
	"time"

	proto2 "xy3-proto"
	"xy3-proto/pkg/conf/env"
	"xy3-proto/pkg/log"

	"github.com/go-kratos/kratos/contrib/registry/discovery/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// AppID .
const ServiceName = "account"

var (
	defautClients    []AccountClient
	maxDefaultClient int64
	count            int64
	mu               sync.Mutex
)

func DefaultClient() AccountClient {
	mu.Lock()
	defer mu.Unlock()
	if defautClients == nil {
		maxDefaultClient = env.CrossGRPCClientNum
		defautClients = make([]AccountClient, maxDefaultClient)

	}
	count++
	if count%1000 == 0 {
		log.Info("new client count:%v", count)
	}
	clientIndex := count % maxDefaultClient
	if defautClients[clientIndex] != nil {
		return defautClients[clientIndex]
	}
	cli, _, err := newClient()
	if err != nil {
		log.Error("new client err:%v", err)
		return nil
	}
	defautClients[clientIndex] = cli
	return defautClients[clientIndex]
}

func newClient() (AccountClient, func(), error) {
	if os.Getenv("DISCOVERY") != "" {
		nodes := strings.Split(env.DiscoveryNodes, ",")
		log.Debug("Discovery Nodes %v", nodes)
		r := discovery.New(&discovery.Config{
			Nodes:  nodes,
			Env:    env.DeployEnv,
			Region: env.Region,
			Zone:   env.Zone,
			Host:   env.Hostname,
		})
		conn, err := grpc.DialInsecure(
			context.Background(),
			grpc.WithEndpoint(fmt.Sprintf("discovery:///%s.%s", env.CrossNamespace, ServiceName)),
			grpc.WithDiscovery(r),
			grpc.WithTimeout(time.Second*5),
		)
		if err != nil {
			log.Error("NewAccount Grpc DialInsecure err %v", err)
			return nil, func() {}, err
		}
		return NewAccountClient(conn), func() {
			conn.Close()
			r.Close()
		}, nil
	} else {
		clientSet, err := proto2.GetClientSet()
		if err != nil {
			log.Error("NewClient GetClientSet err:%v\n", err)
			return nil, nil, err
		}

		r := proto2.NewRegistry(clientSet)
		conn, err := grpc.DialInsecure(
			context.Background(),
			grpc.WithEndpoint(fmt.Sprintf("%s.%s.svc.cluster.local:9000", ServiceName, env.CrossNamespace)), // 可用
			grpc.WithDiscovery(r),
		)
		if err != nil {
			log.Error("NewClient Grpc DialInsecure err %v", err)
			return nil, nil, err
		}
		return NewAccountClient(conn), func() {
			conn.Close()
		}, nil
	}
}
