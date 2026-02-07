package redis

import (
	"errors"

	"github.com/go-redis/redis/v8"
)

// 常见错误
var (
	ErrNil         = redis.Nil
	ErrKeyNotFound = errors.New("key not found")
)

// IsNil 判断是否是 Nil 错误
func IsNil(err error) bool {
	return errors.Is(err, redis.Nil)
}

// IsNotFound 判断是否是键不存在错误
func IsNotFound(err error) bool {
	return IsNil(err)
}
