package main

import (
	"bytes"
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
	config Config
	wg     sync.WaitGroup
)

const (
	configPath = "./config.yaml"
	configType = "yaml"
)

func initViperConfig() {
	viper.SetConfigFile(configPath)
	// read config
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("viper read config failed: %v", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("viper unmarshal config failed: %v", err)
	}
	viper.WatchConfig()
	wg.Done()
}

func initNacosConfig() {
	// server config
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(config.Nacos.Host, uint64(config.Nacos.Port), constant.WithContextPath("/nacos")),
	}
	// client config
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
		log.Fatalf("create nacos config client failed: %v", err)
	}

	jsonBytes, err := json.Marshal(config)
	if err != nil {
		log.Fatalf("marshal json config failed: %v", err)
	}
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  config.Nacos.DataId,
		Group:   config.Nacos.Group,
		Content: string(jsonBytes),
	})
	if err != nil {
		log.Fatalf("nacos publish config failed: %v", err)
	}

	err = client.ListenConfig(vo.ConfigParam{
		DataId: config.Nacos.DataId,
		Group:  config.Nacos.Group,
		OnChange: func(namespace, group, dataId, data string) {
			viper.SetConfigType(configType)
			err := viper.ReadConfig(bytes.NewBufferString(data))
			if err != nil {
				log.Fatal("viper read config failed: ", err)
			}
			err = viper.WriteConfigAs(configPath)
			log.Println("apply new config to local")
			if err != nil {
				log.Fatal("viper write config failed: ", err)
			}
		},
	})
	if err != nil {
		log.Fatalf("nacos listen config failed: %v", err)
	}
}

func main() {
	wg.Add(1)
	go initViperConfig()
	wg.Wait()
	go initNacosConfig()
	time.Sleep(time.Second * 90)
}
