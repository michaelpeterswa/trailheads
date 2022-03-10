package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func InitRedis(ctx context.Context, url string, port int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", url, port),
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
