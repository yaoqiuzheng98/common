package consul

import (
	"log"
	"testing"
)

func TestLoad(t *testing.T) {
	client, err := NewClient(DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}
	host, err := client.GetRedisHost()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(host)
}
