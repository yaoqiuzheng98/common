package redis

import (
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	result := GetClient().GetRedis().Ping()
	log.Println(result)
}
