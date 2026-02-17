package redis

import (
	"fmt"
	"sync"
	"time"

	"github.com/yaoqiuzheng98/common/config"
	"github.com/yaoqiuzheng98/common/etcd"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

var _client = &Client{
	rwMutex: &sync.RWMutex{},
	once:    &sync.Once{},
}

func GetClient() *Client {
	_client.once.Do(func() {
		_client.init()
	})
	_client.rwMutex.RLock()
	defer _client.rwMutex.RUnlock()
	return _client
}

type Client struct {
	rds     *redis.Redis
	rwMutex *sync.RWMutex
	once    *sync.Once
}

func (receiver *Client) GetRedis() *redis.Redis {
	return receiver.rds
}

func connectRedis(c Config) *redis.Redis {
	conf := redis.RedisConf{
		Host:        fmt.Sprintf("%s:%d", c.Host, c.Port),
		Type:        "node",
		Pass:        c.Password,
		Tls:         false,
		NonBlock:    false,
		PingTimeout: 3 * time.Second,
	}
	rds := redis.MustNewRedis(conf)
	return rds
}

func (receiver *Client) init() {
	cfg := config.GetConfig()
	// 创建 etcd subscriber
	ss := subscriber.MustNewEtcdSubscriber(subscriber.EtcdConf{
		Hosts: []string{cfg.EtcdHttpAddr},                  // etcd 地址
		Key:   etcd.GetMiddlewareKey(etcd.MiddlewareRedis), // 配置key
	})

	// 创建 configurator
	cc := configurator.MustNewConfigCenter[Config](configurator.Config{
		Type: "json", // 配置值类型：json,yaml,toml
	}, ss)

	// 获取配置
	// 注意: 配置如果发生变更，调用的结果永远获取到最新的配置
	c, err := cc.GetConfig()
	if err != nil {
		panic(err)
	}
	rds := connectRedis(c)
	_client.rwMutex.Lock()
	defer _client.rwMutex.Unlock()
	_client.rds = rds

	//如果想监听配置变化，可以添加 listener
	cc.AddListener(func() {
		c, err := cc.GetConfig()
		if err != nil {
			panic(err)
		}
		rds := connectRedis(c)
		_client.rwMutex.Lock()
		defer _client.rwMutex.Unlock()
		_client.rds = rds
	})
}
