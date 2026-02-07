package redis

import "log"

// Logger 日志接口
type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
	Warn(msg string)
}

// 默认日志实现
type defaultLogger struct{}

func (l *defaultLogger) Info(msg string) {
	log.Printf("[Redis INFO] %s", msg)
}

func (l *defaultLogger) Error(msg string) {
	log.Printf("[Redis ERROR] %s", msg)
}

func (l *defaultLogger) Debug(msg string) {
	log.Printf("[Redis DEBUG] %s", msg)
}

func (l *defaultLogger) Warn(msg string) {
	log.Printf("[Redis WARN] %s", msg)
}
