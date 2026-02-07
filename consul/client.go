// consul/client.go
package consul

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/yaoqiuzheng98/common/environment"
)

var _once = &sync.Once{}

var _client *Client

func GetClient() *Client {
	if _client == nil {
		panic("consul client not initialized")
	}
	return _client
}

func Init() {
	_once.Do(func() {
		client, err := NewClient(ConfigFromEnv())
		if err != nil {
			panic(err)
		}
		_client = client
		log.Println("consul client initialized")
	})
}

// Client Consul 客户端封装
type Client struct {
	client *consulapi.Client
	config *Config
	env    environment.Environment
}

// NewClient 创建 Consul 客户端
func NewClient(cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	// 验证配置
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	// 创建 Consul 配置
	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = cfg.Address
	consulConfig.Scheme = cfg.Scheme
	consulConfig.Datacenter = cfg.Datacenter
	consulConfig.Token = cfg.Token
	if consulConfig.HttpClient == nil {
		consulConfig.HttpClient = &http.Client{}
	}
	consulConfig.HttpClient.Timeout = cfg.Timeout

	// 创建客户端
	client, err := consulapi.NewClient(consulConfig)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}

	c := &Client{
		client: client,
		config: cfg,
		env:    environment.GetEnvironment(),
	}

	// 健康检查
	if err := c.Health(); err != nil {
		return nil, err
	}

	return c, nil
}

// GetClient 获取原始 Consul 客户端（用于高级操作）
func (c *Client) GetClient() *consulapi.Client {
	return c.client
}

// GetConfig 获取配置
func (c *Client) GetConfig() *Config {
	return c.config
}

// Health 健康检查
func (c *Client) Health() error {
	_, err := c.client.Catalog().Datacenters()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrHealthCheckFailed, err)
	}
	return nil
}

// Close 关闭客户端（预留接口）
func (c *Client) Close() error {
	// Consul API 客户端不需要显式关闭
	return nil
}

func (c *Client) GetRedisHost() (string, error) {
	key := fmt.Sprintf("config/dentistry/%s/redis/host", c.env)
	return c.getValue(key)
}

func (c *Client) GetRedisPort() (string, error) {
	key := fmt.Sprintf("config/dentistry/%s/redis/port", c.env)
	return c.getValue(key)
}

func (c *Client) GetRedisDB() (string, error) {
	key := fmt.Sprintf("config/dentistry/%s/redis/db", c.env)
	return c.getValue(key)
}

func (c *Client) GetRedisPassword() (string, error) {
	key := fmt.Sprintf("config/dentistry/%s/redis/password", c.env)
	return c.getValue(key)
}

func (c *Client) getValue(key string) (string, error) {
	kv, _, err := c.client.KV().Get(key, nil)
	if err != nil {
		return "", err
	}
	return string(kv.Value), nil
}
