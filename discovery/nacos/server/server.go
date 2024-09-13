package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"go-zero-learn/discovery/nacos/server/internal/config"
	"google.golang.org/grpc"
	"log"
	"os"
	"sync"
	"time"

	"go-zero-learn/remote/unary"
)

var configFile = flag.String("f", "etc/nacos.yaml", "the config file")

type GreetServer struct {
	lock     sync.Mutex
	alive    bool
	downTime time.Time
}

func NewGreetServer() *GreetServer {
	return &GreetServer{
		alive: true,
	}
}

func (gs *GreetServer) Greet(ctx context.Context, req *unary.Request) (*unary.Response, error) {
	fmt.Println("=>", req)

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &unary.Response{
		Greet: "hello from " + hostname,
	}, nil
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		unary.RegisterGreeterServer(grpcServer, NewGreetServer())
	})
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		st := time.Now()
		resp, err = handler(ctx, req)
		log.Printf("method: %s time: %v\n", info.FullMethod, time.Since(st))
		return resp, err
	}

	server.AddUnaryInterceptors(interceptor)

	sc := []constant.ServerConfig{
		*constant.NewServerConfig("192.168.188.101", 8848),
	}

	cc := &constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	opts := nacos.NewNacosConfig("nacos.rpc", c.ListenOn, sc, cc)
	err := nacos.RegisterService(opts)
	if err != nil {
		fmt.Println("注册nacos.rpc失败", err)
	}

	server.Start()

}
