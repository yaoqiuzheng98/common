package etcd

import (
	"fmt"

	"github.com/yaoqiuzheng98/common/environment"
	"github.com/yaoqiuzheng98/common/service"
)

func GetMiddlewareKey(middleware Middleware) string {
	return fmt.Sprintf("/config/%s/%s/%s", environment.GetEnvironment(), service.GetService(), middleware)
}
