package redis

import (
	"fmt"
	"time"
)

// Redis 配置
type Config struct {
	// 基础配置
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	DB       int    `json:"db" yaml:"db" mapstructure:"db"`

	// 连接池配置
	PoolSize     int `json:"pool_size" yaml:"pool_size" mapstructure:"pool_size"`
	MinIdleConns int `json:"min_idle_conns" yaml:"min_idle_conns" mapstructure:"min_idle_conns"`
	MaxRetries   int `json:"max_retries" yaml:"max_retries" mapstructure:"max_retries"`

	// 超时配置
	DialTimeout  time.Duration `json:"dial_timeout" yaml:"dial_timeout" mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout" yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout" mapstructure:"write_timeout"`
	PoolTimeout  time.Duration `json:"pool_timeout" yaml:"pool_timeout" mapstructure:"pool_timeout"`

	// 空闲连接配置
	IdleTimeout        time.Duration `json:"idle_timeout" yaml:"idle_timeout" mapstructure:"idle_timeout"`
	IdleCheckFrequency time.Duration `json:"idle_check_frequency" yaml:"idle_check_frequency" mapstructure:"idle_check_frequency"`

	// 集群配置
	ClusterMode  bool     `json:"cluster_mode" yaml:"cluster_mode" mapstructure:"cluster_mode"`
	ClusterAddrs []string `json:"cluster_addrs" yaml:"cluster_addrs" mapstructure:"cluster_addrs"`

	// 哨兵配置
	SentinelMode       bool     `json:"sentinel_mode" yaml:"sentinel_mode" mapstructure:"sentinel_mode"`
	SentinelAddrs      []string `json:"sentinel_addrs" yaml:"sentinel_addrs" mapstructure:"sentinel_addrs"`
	SentinelMasterName string   `json:"sentinel_master_name" yaml:"sentinel_master_name" mapstructure:"sentinel_master_name"`
}

// 默认配置
func DefaultConfig() *Config {
	return &Config{
		Host:               "localhost",
		Port:               6379,
		Password:           "",
		DB:                 0,
		PoolSize:           10,
		MinIdleConns:       5,
		MaxRetries:         3,
		DialTimeout:        5 * time.Second,
		ReadTimeout:        3 * time.Second,
		WriteTimeout:       3 * time.Second,
		PoolTimeout:        4 * time.Second,
		IdleTimeout:        5 * time.Minute,
		IdleCheckFrequency: 1 * time.Minute,
		ClusterMode:        false,
		SentinelMode:       false,
	}
}

// 获取地址
func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// 验证配置
func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("redis host is required")
	}

	if c.Port <= 0 || c.Port > 65535 {
		return fmt.Errorf("invalid redis port: %d", c.Port)
	}

	if c.DB < 0 {
		return fmt.Errorf("invalid redis db: %d", c.DB)
	}

	if c.PoolSize <= 0 {
		c.PoolSize = 10
	}

	if c.MinIdleConns < 0 {
		c.MinIdleConns = 0
	}

	if c.ClusterMode && len(c.ClusterAddrs) == 0 {
		return fmt.Errorf("cluster mode requires cluster addresses")
	}

	if c.SentinelMode && len(c.SentinelAddrs) == 0 {
		return fmt.Errorf("sentinel mode requires sentinel addresses")
	}

	if c.SentinelMode && c.SentinelMasterName == "" {
		return fmt.Errorf("sentinel mode requires master name")
	}

	return nil
}
