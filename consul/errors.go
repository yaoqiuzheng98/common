// consul/errors.go
package consul

import "errors"

var (
	// 配置错误
	ErrInvalidConfig = errors.New("invalid consul config")
	ErrInvalidScheme = errors.New("invalid scheme, must be http or https")

	// 连接错误
	ErrConnectionFailed  = errors.New("failed to connect to consul")
	ErrHealthCheckFailed = errors.New("consul health check failed")

	// 服务注册错误
	ErrRegistrationFailed   = errors.New("service registration failed")
	ErrDeregistrationFailed = errors.New("service deregistration failed")

	// 服务发现错误
	ErrServiceNotFound    = errors.New("service not found")
	ErrNoHealthyInstance  = errors.New("no healthy instance available")
	ErrInvalidServiceName = errors.New("invalid service name")
)
