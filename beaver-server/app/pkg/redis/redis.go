package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const defaultPoolSize = 100

type Client = redis.Client
type Options = redis.Options

func NewClient(options *redis.Options) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
		PoolSize: 100,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
