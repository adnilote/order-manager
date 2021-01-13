package redis

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/adnilote/order-manager/app/business/entities"
	pb "github.com/adnilote/order-manager/app/business/entities/proto"
	redis "github.com/go-redis/redis/v8"
)

func (store *Client) GetOrderByOrderID(orderID string) (pb.Order, error) {
	pattern := fmt.Sprintf("%s:*/%s", store.OrderTableName, orderID)

	var cursor uint64
	var err error
	var keys []string
	for {
		keys, cursor, err = store.client.Scan(store.ctx, cursor, pattern, 100).Result()
		if err != nil {
			return pb.Order{}, err
		}
		if len(keys) != 0 || cursor == 0 {
			break
		}
	}

	if len(keys) == 0 {
		return pb.Order{}, entities.ErrNotFound
	}

	key := strings.TrimPrefix(keys[0], fmt.Sprintf("%s:", store.OrderTableName))
	order, err := store.GetOrderByKey(key)
	if err != nil {
		return pb.Order{}, err
	}
	return order, nil
}

func (store *Client) GetOrderByKey(key string) (pb.Order, error) {
	key = fmt.Sprintf("%s:%s", store.OrderTableName, key)
	data, err := store.client.Get(store.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return pb.Order{}, entities.ErrNotFound
		}
	}

	var order pb.Order
	err = json.Unmarshal([]byte(data), &order)
	if err != nil {
		return pb.Order{}, entities.ErrNotFound
	}
	return order, nil
}

func (store *Client) SaveOrder(order pb.Order, expiration time.Duration) error {
	key := fmt.Sprintf("%s:%s", store.OrderTableName, entities.GetOrderKey(order))

	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = store.client.Set(store.ctx, key, data, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (store *Client) DeleteOrder(order pb.Order) error {
	key := fmt.Sprintf("%s:%s", store.OrderTableName, entities.GetOrderKey(order))
	err := store.client.Do(store.ctx, "del", key).Err()
	if err != nil {
		return err
	}
	return nil
}
