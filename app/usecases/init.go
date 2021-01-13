package usecases

import (
	"context"

	"github.com/adnilote/order-manager/app/config"
)

type Usecases struct {
	OrderManager *OrderManager
}

func Init(ctx context.Context, store Store) (*Usecases, error) {
	orderManager, err := NewOrderManager(ctx, config.Config.Kafka, store)
	if err != nil {
		return nil, err
	}
	us := &Usecases{
		OrderManager: orderManager,
	}
	return us, nil
}

func (u Usecases) Close() {
	u.OrderManager.Close()
}
