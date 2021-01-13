package store

import (
	"context"
	"fmt"
	"time"

	pb "github.com/adnilote/order-manager/app/business/entities/proto"
	"github.com/adnilote/order-manager/app/config"
	"github.com/adnilote/order-manager/app/store/redis"
)

type Store struct {
	redis *redis.Client
}

func NewStore(ctx context.Context) (*Store, error) {
	var store Store
	var err error

	store.redis, err = redis.NewClient(ctx, config.Config.Redis)
	if err != nil {
		return nil, err
	}

	return &store, nil
}

func (store *Store) GetOrder(id string) (pb.Order, error) {
	return store.redis.GetOrderByOrderID(id)
}

func (store *Store) SaveOrder(order pb.Order) error {
	return store.redis.SaveOrder(order, time.Duration(0))
}

func (store *Store) SaveOrderWithExpiration(order pb.Order, expiration time.Duration) error {
	return store.redis.SaveOrder(order, expiration)
}

func (store *Store) DeleteOrder(order pb.Order) error {
	return store.redis.DeleteOrder(order)
}

func (store *Store) GetTotalPosition(account string, symbol string) (pb.TotalPosition, error) {
	key := fmt.Sprintf("%s/%s", account, symbol)
	return store.redis.GetTotalPosition(key)
}

func (store *Store) SaveTotalPosition(position pb.TotalPosition) error {
	return store.redis.SaveTotalPosition(position)
}

func (store *Store) GetCalculatedPosition(account, strategy, contractID string) (pb.CalculatedPosition, error) {
	key := fmt.Sprintf("%s/%s/%s", account, strategy, contractID)
	return store.redis.GetCalculatedPosition(key)
}

func (store *Store) SaveCalculatedPosition(position pb.CalculatedPosition) error {
	return store.redis.SaveCalculatedPosition(position)
}
