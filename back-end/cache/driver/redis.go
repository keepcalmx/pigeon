package driver

import (
	"fmt"
	"log"
	"os"

	"github.com/RediSearch/redisearch-go/redisearch"
	"gopkg.in/yaml.v3"
)

var redisConf map[any]any
var redisAddr string

func init() {
	yf, err := os.ReadFile("./config/redis.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yf, &redisConf)
	if err != nil {
		log.Fatal(err)
	}

	redisAddr = fmt.Sprintf("%v:%v", redisConf["host"], redisConf["port"])
}

func SearchEngine(name string) *redisearch.Client {
	client := redisearch.NewClient(redisAddr, name)
	return client
}
