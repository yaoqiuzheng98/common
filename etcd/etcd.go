package etcd

import (
	"os"

	"github.com/yaoqiuzheng98/common/environment"
)

func GetHttpAddress() string {
	if addr := os.Getenv("ETCD_HTTP_ADDR"); addr != "" {
		return addr
	}
	if environment.GetEnvironment() == environment.Development {
		return "127.0.0.1:2379"
	}
	panic("ETCD_HTTP_ADDR env variable not set")
}
