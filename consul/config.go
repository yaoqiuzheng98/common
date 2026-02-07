// consul/config.go
package consul

import (
	"os"
	"time"
)

// Config Consul 配置
type Config struct {
	Address    string        // Consul 地址，如 "localhost:8500"
	Scheme     string        // http 或 https
	Datacenter string        // 数据中心
	Token      string        // ACL Token
	Timeout    time.Duration // 超时时间
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Address:    "localhost:8500",
		Scheme:     "http",
		Datacenter: "dc1",
		Token:      "",
		Timeout:    10 * time.Second,
	}
}

// ConfigFromEnv 从环境变量加载配置
func ConfigFromEnv() *Config {
	cfg := DefaultConfig()

	if addr := os.Getenv("CONSUL_HTTP_ADDR"); addr != "" {
		cfg.Address = addr
	}
	if scheme := os.Getenv("CONSUL_HTTP_SCHEME"); scheme != "" {
		cfg.Scheme = scheme
	}
	if dc := os.Getenv("CONSUL_DATACENTER"); dc != "" {
		cfg.Datacenter = dc
	}
	if token := os.Getenv("CONSUL_HTTP_TOKEN"); token != "" {
		cfg.Token = token
	}
	if timeout := os.Getenv("CONSUL_TIMEOUT"); timeout != "" {
		if d, err := time.ParseDuration(timeout); err == nil {
			cfg.Timeout = d
		}
	}

	return cfg
}

// Validate 验证配置
func (c *Config) Validate() error {
	if c.Address == "" {
		return ErrInvalidConfig
	}
	if c.Scheme != "http" && c.Scheme != "https" {
		return ErrInvalidScheme
	}
	return nil
}
