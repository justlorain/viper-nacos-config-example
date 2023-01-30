package main

type NacosConfig struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	DataId string `json:"dataId"`
	Group  string `json:"group"`
}

type MySQLConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}
