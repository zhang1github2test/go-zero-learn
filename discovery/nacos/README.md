### 1、基于nacos的rpc服务注册

安装依赖包：

```go
go get  github.com/zeromicro/zero-contrib/zrpc/registry/nacos
```

核心代码

```go
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

```

完成注册后，在nacos上可以看到对应的服务。

![image-20240913174955442](README.assets/image-20240913174955442.png)

### 2、基于nacos的rpc服务发现

客户端代码：

```go
package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
	"github.com/zeromicro/zero-examples/rpc/remote/unary"
)

const timeFormat = "15:04:05"

var config = flag.String("f", "etc/consul.yaml", "config file")

func main() {
	flag.Parse()

	var c zrpc.RpcClientConf
	conf.MustLoad(*config, &c)
	client := zrpc.MustNewClient(c)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			conn := client.Conn()
			greet := unary.NewGreeterClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			resp, err := greet.Greet(ctx, &unary.Request{
				Name: "kevin",
			})
			if err != nil {
				fmt.Printf("%s X %s\n", time.Now().Format(timeFormat), err.Error())
			} else {
				fmt.Printf("%s => %s\n", time.Now().Format(timeFormat), resp.Greet)
			}
			cancel()
		}
	}
}

```

客户端的改动点：

#### 1、引入nacos的rpc协议解析包

```go
_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
```



```yaml
Target: nacos://192.168.188.101:8848/nacos.rpc?timeout=5000s
```

target的格式是什么样的？

```txt
nacos://[user:passwd]@host/service?param=value
```

其中param可以配置的值有namespaceid、timeout、appName、notLoadCacheAtStart、updateCacheWhenEmpty等属性可以配置。

#### 2、在服务的配置项里面添加Target属性

```go
Target: nacos://192.168.188.101:8848/nacos.rpc?timeout=5000s
```



