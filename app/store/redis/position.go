package redis

import (
	"encoding/json"
	"fmt"

	"github.com/adnilote/order-manager/app/business/entities"
	pb "github.com/adnilote/order-manager/app/business/entities/proto"
	redis "github.com/go-redis/redis/v8"
)

func (store *Client) GetTotalPosition(key string) (pb.TotalPosition, error) {
	data, err := store.client.Do(store.ctx,
		"HGET", store.TotalPositionTableName,
		key).
		Text()
	if err != nil {
		if err == redis.Nil {
			return pb.TotalPosition{}, entities.ErrNotFound
		}
	}

	var position pb.TotalPosition
	err = json.Unmarshal([]byte(data), &position)
	if err != nil {
		return pb.TotalPosition{}, entities.ErrNotFound
	}
	return position, nil
}

func (store *Client) GetAllTotalPosition() ([]pb.TotalPosition, error) {
	data, err := store.client.Do(store.ctx,
		"HGETALL", store.TotalPositionTableName).
		Result()
	if err != nil {
		if err == redis.Nil {
			return nil, entities.ErrNotFound
		}
	}

	resp := data.([]string)
	var result []pb.TotalPosition
	var position pb.TotalPosition
	for i := range resp {
		err = json.Unmarshal([]byte(resp[i]), &position)
		if err != nil {
			return nil, err
		}
		result = append(result, position)
	}

	return result, nil
}

func (store *Client) SaveTotalPosition(position pb.TotalPosition) error {
	data, err := json.Marshal(position)
	if err != nil {
		return err
	}

	err = store.client.Do(store.ctx,
		"HSET", store.TotalPositionTableName,
		entities.GetTotalPositionKey(position), data).
		Err()
	if err != nil {
		return err
	}
	return nil
}

func (store *Client) DeleteTotalPosition(position pb.TotalPosition) error {
	key := entities.GetTotalPositionKey(position)
	res, err := store.client.Do(store.ctx, "HDEL",
		store.TotalPositionTableName, key).Result()
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (store *Client) SaveCalculatedPosition(position pb.CalculatedPosition) error {
	data, err := json.Marshal(position)
	if err != nil {
		return err
	}

	err = store.client.Do(store.ctx,
		"HSET", store.CalculatedPositionTableName,
		entities.GetCalculatedPositionKey(position), data).
		Err()
	if err != nil {
		return err
	}
	return nil
}

func (store *Client) GetCalculatedPosition(key string) (pb.CalculatedPosition, error) {
	data, err := store.client.Do(store.ctx,
		"HGET", store.CalculatedPositionTableName,
		key).
		Text()
	if err != nil {
		if err == redis.Nil {
			return pb.CalculatedPosition{}, entities.ErrNotFound
		}
	}

	var position pb.CalculatedPosition
	err = json.Unmarshal([]byte(data), &position)
	if err != nil {
		return pb.CalculatedPosition{}, entities.ErrNotFound
	}
	return position, nil
}
