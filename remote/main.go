package main

import (
	"bytes"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"log"
)

var config Config

const configPath = "./config.yaml"

func initViperConfig() {
	viper.SetConfigFile(configPath)
	// read nacosConfig
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("viper read nacosConfig failed: %v", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("viper unmarshal nacosConfig failed: %v", err)
	}
}

func initNacosConfig() {
	// server nacosConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(config.Nacos.Host, uint64(config.Nacos.Port), constant.WithContextPath("/nacos")),
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
		DataId: config.Nacos.DataId,
		Group:  config.Nacos.Group,
	})

	err = viper.ReadConfig(bytes.NewBufferString(content))
	if err != nil {
		log.Fatal("viper read config failed: ", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("viper unmarshal config failed: %v", err)
	}

	fmt.Println(config)
}

func main() {
	initViperConfig()
	initNacosConfig()
}
