package usecases

import (
	"context"
	"time"

	pb "github.com/adnilote/order-manager/app/business/entities/proto"
	"github.com/adnilote/order-manager/app/config"
	"github.com/adnilote/order-manager/app/kafka"
	"github.com/adnilote/order-manager/app/usecases/streamer"
)

type Store interface {
	SaveOrder(pb.Order) error
	SaveOrderWithExpiration(order pb.Order, expiration time.Duration) error
	GetOrder(string) (pb.Order, error)
	DeleteOrder(pb.Order) error
	GetTotalPosition(account string, contractID string) (pb.TotalPosition, error)
	SaveTotalPosition(pb.TotalPosition) error
	GetCalculatedPosition(account, strategy, contractID string) (pb.CalculatedPosition, error)
	SaveCalculatedPosition(pb.CalculatedPosition) error
}

type EventStreamer interface {
	SendToAll(data interface{}) error
	Register(wsuuid string) chan interface{}
	Unregister(wsuuid string)
}

type OrderManager struct {
	orderStreamer EventStreamer
	consumer      kafka.Consumer
}

func NewOrderManager(ctx context.Context, config config.Kafka, store Store) (*OrderManager, error) {
	orderStreamer := streamer.NewEventStreamer(5)

	handler := &OrderHandler{
		store:         store,
		orderStreamer: orderStreamer,
	}
	consumer, err := kafka.NewConsumer(config.OrderConsumer, config.Address, handler)
	if err != nil {
		return nil, err
	}

	err = consumer.Start(ctx)
	if err != nil {
		return nil, err
	}

	manager := &OrderManager{
		orderStreamer: orderStreamer,
	}
	return manager, nil
}

func (manager *OrderManager) Register(wsuuid string) chan interface{} {
	return manager.orderStreamer.Register(wsuuid)
}

func (manager *OrderManager) Unregister(wsuuid string) {
	manager.orderStreamer.Unregister(wsuuid)
}

func (manager *OrderManager) Close() {
	manager.consumer.Close()
}
