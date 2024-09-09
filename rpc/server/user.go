package main

import (
	"flag"
	"fmt"

	"go-zero-learn/rpc/server/internal/config"
	productServer "go-zero-learn/rpc/server/internal/server/product"
	userServer "go-zero-learn/rpc/server/internal/server/user"
	"go-zero-learn/rpc/server/internal/svc"
	"go-zero-learn/rpc/server/pb/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		user.RegisterProductServer(grpcServer, productServer.NewProductServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
