package main

import (
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func init() {
	// 设置配置文件名和路径
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("当前工作目录:", dir)
	}
	viper.AddConfigPath(dir) // 当前目录
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件出错:", err)
		return
	}
}
func main() {
	var listenPort = flag.Uint64("listenPort", 8080, "the config file")
	flag.Parse()
	http.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("hello"))
	})
	// 获取环境变量
	host := viper.GetString("nacos.host")
	port := viper.GetUint64("nacos.port")
	namespaceId := viper.GetString("nacos.nameSpaceid")
	timeoutMs := viper.GetUint64("nacos.timeOutms")
	logDir := viper.GetString("nacos.logdir")
	cacheDir := viper.GetString("nacos.cacheDir")
	logLevel := viper.GetString("nacos.logLevel")

	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(host, port, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(namespaceId),
		constant.WithTimeoutMs(timeoutMs),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(logDir),
		constant.WithCacheDir(cacheDir),
		constant.WithLogLevel(logLevel),
	)

	// create naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		fmt.Println("erro")
	}
	client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "127.0.0.1",
		Port:        *listenPort,
		ServiceName: "hello-service",
		//GroupName:   "group-a",
		//ClusterName: "cluster-a",
		Weight:    10,
		Enable:    true,
		Healthy:   true,
		Ephemeral: true,
		Metadata:  map[string]string{"idc": "shanghai"},
	})

	err = http.ListenAndServe(fmt.Sprintf(":%v", *listenPort), nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
