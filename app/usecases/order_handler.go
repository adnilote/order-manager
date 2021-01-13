package usecases

import (
	"strings"
	"time"

	"github.com/adnilote/order-manager/app/business/entities"
	pb "github.com/adnilote/order-manager/app/business/entities/proto"
	"github.com/adnilote/order-manager/app/business/rules"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/sirupsen/logrus"
	confluent "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type OrderHandler struct {
	store         Store
	orderStreamer EventStreamer
}

func (handler *OrderHandler) HandleTotalPosition(order, oldOrder pb.Order) error {

	curTotalPosition, err := handler.store.GetTotalPosition(order.Account, order.OrderParameters.Instrument.Symbol)
	if err != nil {
		if err != entities.ErrNotFound {
			return err
		}
		curTotalPosition.Instrument = order.OrderParameters.Instrument
		curTotalPosition.Account = order.Account
		// position.Currency = order.OrderParameters.Currency TODO
	}

	totalPos, err := getTotalPosition(order, oldOrder, curTotalPosition)
	if err != nil {
		return err
	}

	err = handler.store.SaveTotalPosition(totalPos)
	if err != nil {
		return err
	}

	return nil
}

func (handler *OrderHandler) HandleCalculatedPosition(order, oldOrder pb.Order) error {

	strategy, _ := entities.ParseClientTag(order.ClientTag)
	curCalcPosition, err := handler.store.GetCalculatedPosition(
		order.Account,
		strategy,
		order.OrderParameters.Instrument.Id)
	if err != nil {
		if err != entities.ErrNotFound {
			return err
		}
		curCalcPosition.Instrument = order.OrderParameters.Instrument
		curCalcPosition.Account = order.Account
		curCalcPosition.Strategy, _ = entities.ParseClientTag(order.ClientTag)
		// position.Currency = order.OrderParameters.Currency TODO
	}

	calcPos, err := getCalculatedPosition(order, oldOrder, curCalcPosition)
	if err != nil {
		return err
	}

	err = handler.store.SaveCalculatedPosition(calcPos)
	if err != nil {
		return err
	}

	return nil
}

func (handler *OrderHandler) Handle(msg *confluent.Message) error {
	log := logrus.WithFields(logrus.Fields{
		"order": string(msg.Value),
		"msg":   msg.String(),
	})

	order := pb.Order{}
	err := jsonpb.Unmarshal(strings.NewReader(string(msg.Value)), &order)
	if err != nil {
		log.WithError(err).Error("unmarshal")
		return nil
	}

	active := rules.IsActive(order)
	duringDay := rules.IsIn24HoursFromNow(order.PlaceTime)
	if active || duringDay {
		err := handler.orderStreamer.SendToAll(order)
		if err != nil {
			log.WithError(err).Error("send to order streamer")
		}
	}

	oldOrder, err := handler.store.GetOrder(order.OrderId)
	if err != nil && err != entities.ErrNotFound {
		log.WithError(err).Error("get last order")
		return err
	}

	err = handler.HandleTotalPosition(order, oldOrder)
	if err != nil {
		log.WithError(err).Error("handle total pos")
		return err
	}

	err = handler.HandleCalculatedPosition(order, oldOrder)
	if err != nil {
		log.WithError(err).Error("handle calc pos")
		return err
	}

	switch {
	case !active && !duringDay:
		err = handler.store.DeleteOrder(order)
	case !active && duringDay:
		err = handler.store.SaveOrderWithExpiration(order, time.Duration(time.Hour*24))
	case active && !duringDay:
		err = handler.store.SaveOrder(order)
	case active && duringDay:
		err = handler.store.SaveOrder(order)
	}

	if err != nil {
		log.WithError(err).Error("save order to cache")
		return err
	}

	return nil
}
