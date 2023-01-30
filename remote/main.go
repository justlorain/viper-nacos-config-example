package main

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

var (
	nacosConfig NacosConfig
	mysqlConfig MySQLConfig
	wg          sync.WaitGroup
)

const configPath = "./remote/config.yaml"

func initViperConfig() {
	// note: use absolute path
	viper.SetConfigFile(configPath)
	// read nacosConfig
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("viper read nacosConfig failed: %v", err)
	}
	err = viper.Unmarshal(&nacosConfig)
	if err != nil {
		log.Fatalf("viper unmarshal nacosConfig failed: %v", err)
	}
	wg.Done()
}

func initNacosConfig() {
	// server nacosConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(nacosConfig.Host, uint64(nacosConfig.Port), constant.WithContextPath("/nacos")),
	}
	// client nacosConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
	)

	// client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Fatalf("create nacos nacosConfig client failed: %v", err)
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataId,
		Group:  nacosConfig.Group,
	})

	err = json.Unmarshal([]byte(content), &mysqlConfig)
	if err != nil {
		log.Fatalf("json unmatshal failed: %v", err)
	}

}

func main() {
	wg.Add(1)
	go initViperConfig()
	wg.Wait()
	go initNacosConfig()
	time.Sleep(time.Second * 90)
}
