package redis

import (
	"strconv"
	"testing"

	"github.com/yaoqiuzheng98/common/consul"
)

func TestRedis(t *testing.T) {
	client := consul.GetClient()
	host, err := client.GetRedisHost()
	if err != nil {
		t.Fatal(err)
	}
	port, err := client.GetRedisPort()
	if err != nil {
		t.Fatal(err)
	}
	db, err := client.GetRedisDB()
	if err != nil {
		t.Fatal(err)
	}
	password, err := client.GetRedisPassword()
	if err != nil {
		t.Fatal(err)
	}

	config := DefaultConfig()
	config.Host = host
	portNumber, err := strconv.Atoi(port)
	if err != nil {
		t.Fatal(err)
		return
	}
	dbNumber, err := strconv.Atoi(db)
	if err != nil {
		t.Fatal(err)
		return
	}
	config.Port = portNumber
	config.DB = dbNumber
	config.Password = password
	c, err := NewClient(config)
	if err != nil {
		t.Fatal(err)
		return
	}
	err = c.Ping()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ok")
}
