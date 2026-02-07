package redis

import (
	"fmt"
	"time"
)

// BuildKey 构建键名
func BuildKey(parts ...string) string {
	key := ""
	for i, part := range parts {
		if i > 0 {
			key += ":"
		}
		key += part
	}
	return key
}

// BuildCacheKey 构建缓存键
func BuildCacheKey(prefix, key string) string {
	return fmt.Sprintf("cache:%s:%s", prefix, key)
}

// BuildLockKey 构建锁键
func BuildLockKey(resource string) string {
	return fmt.Sprintf("lock:%s", resource)
}

// ParseDuration 解析时间字符串
func ParseDuration(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}

// FormatDuration 格式化时间
func FormatDuration(d time.Duration) string {
	return d.String()
}
