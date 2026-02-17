package config

import (
	"github.com/yaoqiuzheng98/common/environment"
	"github.com/yaoqiuzheng98/common/etcd"
	"github.com/yaoqiuzheng98/common/service"
)

type Config struct {
	EtcdHttpAddr string
	Service      service.Service
	Env          environment.Environment
}

var _config = Config{
	EtcdHttpAddr: etcd.GetHttpAddress(),
	Service:      service.GetService(),
	Env:          environment.GetEnvironment(),
}

func GetConfig() Config {
	return _config
}
