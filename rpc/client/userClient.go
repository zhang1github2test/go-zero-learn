package main

import (
	"context"
	"flag"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-learn/rpc/client/config"
	"go-zero-learn/rpc/client/user"
	"log"
)

var configFile = flag.String("f", "etc/user_client.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	conn := zrpc.MustNewClient(c.RpcClientConf)
	client := user.NewUser(conn)
	resp, err := client.Create(context.Background(), &user.UserReq{
		Id:   uuid.NewString(),
		Name: "zhang",
		Age:  16,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}
