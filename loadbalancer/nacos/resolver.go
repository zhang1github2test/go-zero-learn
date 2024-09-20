package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go-zero-learn/loadbalancer"
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
	loadbalancer.Register(&NacosResolver{})
}

type NacosResolver struct {
	Starting bool
	client   naming_client.INamingClient
}

func (r *NacosResolver) Scheme() string {
	return "nacos"
}

func (r *NacosResolver) Start() bool {
	if r.client != nil {
		return true
	}
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
		return false
	}
	r.client = client
	return true
}

func (r *NacosResolver) Next(serviceName string) string {
	groupName := viper.GetString("nacos.groupName")
	clusterName := viper.GetString("nacos.clusterName")

	insance, err := r.client.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		GroupName:   groupName,
		Clusters:    []string{clusterName},
	})
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%v:%v", insance.Ip, insance.Port)
}

func (r *NacosResolver) Running() bool {
	return r.client != nil
}
