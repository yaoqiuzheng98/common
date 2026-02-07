package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// ==================== 字符串操作 ====================

// Set 设置键值
func (c *Client) Set(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.Set(ctx, key, value, expiration).Err()
}

// Get 获取值
func (c *Client) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.Get(ctx, key).Result()
}

// GetEx 获取值并设置过期时间
func (c *Client) GetEx(key string, expiration time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.GetEx(ctx, key, expiration).Result()
}

// SetNX 设置键值（仅当键不存在时）
func (c *Client) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.SetNX(ctx, key, value, expiration).Result()
}

// MGet 批量获取
func (c *Client) MGet(keys ...string) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.MGet(ctx, keys...).Result()
}

// MSet 批量设置
func (c *Client) MSet(pairs ...interface{}) error {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.MSet(ctx, pairs...).Err()
}

// Incr 自增
func (c *Client) Incr(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.Incr(ctx, key).Result()
}

// IncrBy 增加指定值
func (c *Client) IncrBy(key string, value int64) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.IncrBy(ctx, key, value).Result()
}

// Decr 自减
func (c *Client) Decr(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.Decr(ctx, key).Result()
}

// ==================== 键操作 ====================

// Exists 检查键是否存在
func (c *Client) Exists(keys ...string) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.Exists(ctx, keys...).Result()
}

// Del 删除键
func (c *Client) Del(keys ...string) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.Del(ctx, keys...).Result()
}

// Expire 设置过期时间
func (c *Client) Expire(key string, expiration time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.Expire(ctx, key, expiration).Result()
}

// TTL 获取剩余过期时间
func (c *Client) TTL(key string) (time.Duration, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.TTL(ctx, key).Result()
}

// Rename 重命名键
func (c *Client) Rename(key, newKey string) error {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.Rename(ctx, key, newKey).Err()
}

// Keys 查找键（生产环境慎用）
func (c *Client) Keys(pattern string) ([]string, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.Keys(ctx, pattern).Result()
}

// Scan 扫描键（推荐使用）
func (c *Client) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.Scan(ctx, cursor, match, count).Result()
}

// ==================== 哈希操作 ====================

// HSet 设置哈希字段
func (c *Client) HSet(key string, values ...interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.HSet(ctx, key, values...).Result()
}

// HGet 获取哈希字段
func (c *Client) HGet(key, field string) (string, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.HGet(ctx, key, field).Result()
}

// HGetAll 获取所有哈希字段
func (c *Client) HGetAll(key string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.HGetAll(ctx, key).Result()
}

// HMGet 批量获取哈希字段
func (c *Client) HMGet(key string, fields ...string) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.HMGet(ctx, key, fields...).Result()
}

// HDel 删除哈希字段
func (c *Client) HDel(key string, fields ...string) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.HDel(ctx, key, fields...).Result()
}

// HExists 检查哈希字段是否存在
func (c *Client) HExists(key, field string) (bool, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.HExists(ctx, key, field).Result()
}

// HIncrBy 哈希字段增加
func (c *Client) HIncrBy(key, field string, incr int64) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.HIncrBy(ctx, key, field, incr).Result()
}

// ==================== 列表操作 ====================

// LPush 从左侧插入
func (c *Client) LPush(key string, values ...interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.LPush(ctx, key, values...).Result()
}

// RPush 从右侧插入
func (c *Client) RPush(key string, values ...interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.RPush(ctx, key, values...).Result()
}

// LPop 从左侧弹出
func (c *Client) LPop(key string) (string, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.LPop(ctx, key).Result()
}

// RPop 从右侧弹出
func (c *Client) RPop(key string) (string, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.RPop(ctx, key).Result()
}

// LRange 获取列表范围
func (c *Client) LRange(key string, start, stop int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.LRange(ctx, key, start, stop).Result()
}

// LLen 获取列表长度
func (c *Client) LLen(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.LLen(ctx, key).Result()
}

// ==================== 集合操作 ====================

// SAdd 添加集合成员
func (c *Client) SAdd(key string, members ...interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.SAdd(ctx, key, members...).Result()
}

// SMembers 获取所有集合成员
func (c *Client) SMembers(key string) ([]string, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.SMembers(ctx, key).Result()
}

// SIsMember 检查是否是集合成员
func (c *Client) SIsMember(key string, member interface{}) (bool, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.SIsMember(ctx, key, member).Result()
}

// SRem 移除集合成员
func (c *Client) SRem(key string, members ...interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.SRem(ctx, key, members...).Result()
}

// SCard 获取集合成员数量
func (c *Client) SCard(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.SCard(ctx, key).Result()
}

// ==================== 有序集合操作 ====================

// ZAdd 添加有序集合成员
func (c *Client) ZAdd(key string, members ...*redis.Z) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.ZAdd(ctx, key, members...).Result()
}

// ZRange 获取有序集合范围
func (c *Client) ZRange(key string, start, stop int64) ([]string, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.ZRange(ctx, key, start, stop).Result()
}

// ZRangeWithScores 获取有序集合范围（带分数）
func (c *Client) ZRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.ZRangeWithScores(ctx, key, start, stop).Result()
}

// ZRem 移除有序集合成员
func (c *Client) ZRem(key string, members ...interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.ZRem(ctx, key, members...).Result()
}

// ZScore 获取成员分数
func (c *Client) ZScore(key, member string) (float64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.ZScore(ctx, key, member).Result()
}

// ZCard 获取有序集合成员数量
func (c *Client) ZCard(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.ReadTimeout)
	defer cancel()

	return c.client.ZCard(ctx, key).Result()
}

// ==================== JSON 操作（辅助方法）====================

// SetJSON 设置 JSON 对象
func (c *Client) SetJSON(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	return c.Set(key, data, expiration)
}

// GetJSON 获取 JSON 对象
func (c *Client) GetJSON(key string, dest interface{}) error {
	data, err := c.Get(key)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(data), dest); err != nil {
		return fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return nil
}

// ==================== 分布式锁 ====================

// Lock 获取分布式锁
func (c *Client) Lock(key string, expiration time.Duration) (bool, error) {
	return c.SetNX(key, "locked", expiration)
}

// Unlock 释放分布式锁
func (c *Client) Unlock(key string) error {
	_, err := c.Del(key)
	return err
}

// ==================== 管道操作 ====================

// Pipeline 执行管道操作
func (c *Client) Pipeline(fn func(redis.Pipeliner) error) error {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	pipe := c.client.Pipeline()

	if err := fn(pipe); err != nil {
		return err
	}

	_, err := pipe.Exec(ctx)
	return err
}

// ==================== 事务操作 ====================

// Transaction 执行事务
func (c *Client) Transaction(fn func(*redis.Tx) error, keys ...string) error {
	ctx, cancel := context.WithTimeout(c.ctx, c.config.WriteTimeout)
	defer cancel()

	return c.client.Watch(ctx, fn, keys...)
}
