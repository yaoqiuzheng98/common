package redis

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yaoqiuzheng98/common/consul"
)

var _once = &sync.Once{}
var _client *Client

func GetClient() *Client {
	return _client
}

func Init() {
	_once.Do(func() {
		consul.Init()
		client := consul.GetClient()
		host, err := client.GetRedisHost()
		if err != nil {
			panic(err)
		}
		port, err := client.GetRedisPort()
		if err != nil {
			panic(err)
		}
		db, err := client.GetRedisDB()
		if err != nil {
			panic(err)
		}
		password, err := client.GetRedisPassword()
		if err != nil {
			panic(err)
		}

		config := DefaultConfig()
		config.Host = host
		portNumber, err := strconv.Atoi(port)
		if err != nil {
			panic(err)
		}
		dbNumber, err := strconv.Atoi(db)
		if err != nil {
			panic(err)
		}
		config.Port = portNumber
		config.DB = dbNumber
		config.Password = password
		c, err := NewClient(config)
		if err != nil {
			panic(err)
		}
		_client = c
		log.Println("redis 初始化完毕")
	})

}

// Redis 客户端
type Client struct {
	client     redis.UniversalClient
	config     *Config
	logger     Logger
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// 创建 Redis 客户端
func NewClient(config *Config, opts ...Option) (*Client, error) {
	// 验证配置
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	// 创建客户端实例
	client := &Client{
		config: config,
		logger: &defaultLogger{},
	}

	// 应用选项
	for _, opt := range opts {
		opt(client)
	}

	// 创建上下文
	client.ctx, client.cancelFunc = context.WithCancel(context.Background())

	// 根据模式创建不同的客户端
	var err error
	if config.ClusterMode {
		client.client, err = client.newClusterClient()
	} else if config.SentinelMode {
		client.client, err = client.newSentinelClient()
	} else {
		client.client, err = client.newStandaloneClient()
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create redis client: %w", err)
	}

	// 测试连接
	if err := client.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	client.logger.Info("Redis client connected successfully")

	return client, nil
}

// 创建单机客户端
func (c *Client) newStandaloneClient() (redis.UniversalClient, error) {
	return redis.NewClient(&redis.Options{
		Addr:               c.config.Addr(),
		Password:           c.config.Password,
		DB:                 c.config.DB,
		PoolSize:           c.config.PoolSize,
		MinIdleConns:       c.config.MinIdleConns,
		MaxRetries:         c.config.MaxRetries,
		DialTimeout:        c.config.DialTimeout,
		ReadTimeout:        c.config.ReadTimeout,
		WriteTimeout:       c.config.WriteTimeout,
		PoolTimeout:        c.config.PoolTimeout,
		IdleTimeout:        c.config.IdleTimeout,
		IdleCheckFrequency: c.config.IdleCheckFrequency,
	}), nil
}

// 创建集群客户端
func (c *Client) newClusterClient() (redis.UniversalClient, error) {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              c.config.ClusterAddrs,
		Password:           c.config.Password,
		PoolSize:           c.config.PoolSize,
		MinIdleConns:       c.config.MinIdleConns,
		MaxRetries:         c.config.MaxRetries,
		DialTimeout:        c.config.DialTimeout,
		ReadTimeout:        c.config.ReadTimeout,
		WriteTimeout:       c.config.WriteTimeout,
		PoolTimeout:        c.config.PoolTimeout,
		IdleTimeout:        c.config.IdleTimeout,
		IdleCheckFrequency: c.config.IdleCheckFrequency,
	}), nil
}

// 创建哨兵客户端
func (c *Client) newSentinelClient() (redis.UniversalClient, error) {
	return redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:         c.config.SentinelMasterName,
		SentinelAddrs:      c.config.SentinelAddrs,
		Password:           c.config.Password,
		DB:                 c.config.DB,
		PoolSize:           c.config.PoolSize,
		MinIdleConns:       c.config.MinIdleConns,
		MaxRetries:         c.config.MaxRetries,
		DialTimeout:        c.config.DialTimeout,
		ReadTimeout:        c.config.ReadTimeout,
		WriteTimeout:       c.config.WriteTimeout,
		PoolTimeout:        c.config.PoolTimeout,
		IdleTimeout:        c.config.IdleTimeout,
		IdleCheckFrequency: c.config.IdleCheckFrequency,
	}), nil
}

// 获取原始客户端
func (c *Client) GetClient() redis.UniversalClient {
	return c.client
}

// 获取上下文
func (c *Client) Context() context.Context {
	return c.ctx
}

// Ping 测试连接
func (c *Client) Ping() error {
	ctx, cancel := context.WithTimeout(c.ctx, 3*time.Second)
	defer cancel()

	return c.client.Ping(ctx).Err()
}

// 关闭客户端
func (c *Client) Close() error {
	c.logger.Info("Closing Redis client...")

	if c.cancelFunc != nil {
		c.cancelFunc()
	}

	if c.client != nil {
		return c.client.Close()
	}

	return nil
}

// 获取连接池统计信息
func (c *Client) PoolStats() *redis.PoolStats {
	return c.client.PoolStats()
}

// 打印连接池统计信息
func (c *Client) PrintPoolStats() {
	stats := c.PoolStats()
	c.logger.Info(fmt.Sprintf("Redis Pool Stats: Hits=%d, Misses=%d, Timeouts=%d, TotalConns=%d, IdleConns=%d, StaleConns=%d",
		stats.Hits, stats.Misses, stats.Timeouts, stats.TotalConns, stats.IdleConns, stats.StaleConns))
}
