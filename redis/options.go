package redis

import "context"

// Option 配置选项
type Option func(*Client)

// WithLogger 设置日志器
func WithLogger(logger Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithContext 设置上下文
func WithContext(ctx context.Context) Option {
	return func(c *Client) {
		c.ctx = ctx
	}
}
