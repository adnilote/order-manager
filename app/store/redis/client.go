package redis

import (
	"context"

	"github.com/adnilote/order-manager/app/config"
	redis "github.com/go-redis/redis/v8"
)

type Client struct {
	ctx                         context.Context
	client                      *redis.Client
	config                      config.Redis
	OrderTableName              string
	TotalPositionTableName      string
	CalculatedPositionTableName string
}

func NewClient(ctx context.Context, config config.Redis) (*Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.Database,
	})

	redis := Client{
		ctx:                         ctx,
		client:                      redisClient,
		config:                      config,
		OrderTableName:              "orders",
		TotalPositionTableName:      "total-positions",
		CalculatedPositionTableName: "calculated-positions",
	}

	err := redis.Ping()
	if err != nil {
		return nil, err
	}
	return &redis, nil
}

func (store *Client) Ping() error {
	_, err := store.client.Ping(store.ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
