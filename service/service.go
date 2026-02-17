package service

import "os"

func GetService() string {
	service := os.Getenv("Service")
	if service == "" {
		panic("Service 环境变量未设置")
	}
	return service
}
