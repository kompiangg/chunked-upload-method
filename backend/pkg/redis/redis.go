package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Hostname string
	Username string
	Password string
	DB       int
}

func InitRedis(config RedisConfig) *redis.Client {
	var rdb *redis.Client
	sleepDuration := 1

	for {
		rdb = redis.NewClient(&redis.Options{
			Username: config.Username,
			Addr:     config.Hostname,
			Password: config.Password,
			DB:       config.DB,
		})

		ctx := context.Background()

		err := rdb.Ping(ctx).Err()
		if err == nil {
			break
		}

		log.Printf("[INFO] redis connection cant be established, trying to connect in %d second\n", sleepDuration)
		log.Println("[INFO] cause:", err)
		time.Sleep(time.Duration(sleepDuration * int(time.Second)))
		sleepDuration++
	}

	return rdb
}
