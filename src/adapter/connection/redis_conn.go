package connection

import (
	"context"
	"fmt"
	"github.com/rasyidmm/superindo-test/internal/config"
	dbConfig "github.com/rasyidmm/superindo-test/internal/config/redis"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ClientRedis *redis.Client
)

func init() {
	cfg := config.GetConfig()
	ClientRedis, err = connection(cfg.Redis.Superindo)
	if err != nil {
		panic(err)
	}
}

func connection(config dbConfig.Redis) (*redis.Client, error) {
	var (
		client *redis.Client
		err    error
	)
	currentWaitTime := 2
	trialCount := 0
	for client == nil && trialCount < 5 {
		client = redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%s:%s", config.Host, config.Port), Username: config.Username, Password: config.Password})

		if _, err = client.Ping(context.Background()).Result(); err != nil {
			fmt.Println("unable connecting to Redis.")
			fmt.Println("retrying in", currentWaitTime, "seconds...")
			time.Sleep(time.Duration(currentWaitTime) * time.Second)
			currentWaitTime = currentWaitTime * 2
			client = nil
		}
		trialCount++
	}
	if _, err = client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return client, nil
}
