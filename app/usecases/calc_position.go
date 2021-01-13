package usecases

import (
	"errors"

	pb "github.com/adnilote/order-manager/app/business/entities/proto"
	"github.com/sirupsen/logrus"
)

var (
	ErrInvalidOrder = errors.New("invalid order")
)

func getAveragePrice(curAveragePrice, tradeAveragePrice, curQuantity, tradeQuantity float64) (float64, error) {
	switch {
	// increases position from zero
	case curQuantity == 0:
		return tradeAveragePrice, nil

	// increases position
	case sgn(curQuantity*tradeQuantity) > 0:
		return (curQuantity*curAveragePrice + tradeQuantity*tradeAveragePrice) / (tradeQuantity + curQuantity), nil

	// decreases position
	case sgn(curQuantity*tradeQuantity) < 0 && abs(curQuantity) >= abs(tradeQuantity):
		return curAveragePrice, nil

	// reverses position
	case sgn(curQuantity*tradeQuantity) < 0 && abs(curQuantity) < abs(tradeQuantity):
		return tradeAveragePrice, nil

	default:
		return curAveragePrice, errors.New("avr price: unexpected situation")
	}
}

func getTotalPosition(order, oldOrder pb.Order, curPos pb.TotalPosition) (pb.TotalPosition, error) {
	var err error
	if order.OrderParameters == nil {
		return pb.TotalPosition{}, ErrInvalidOrder
	}

	updates := GetUpdate(oldOrder, order)

	position := curPos

	for _, trade := range updates {

		quantity := getQuantity(trade.Quantity, order.OrderParameters.Side.String())

		position.AveragePrice, err = getAveragePrice(position.AveragePrice, trade.Price,
			position.Quantity, quantity)
		if err != nil {
			logrus.WithFields(logrus.Fields{"order": order, "oldOrder": oldOrder, "curPos": curPos}).
				WithError(err).Error()
		}

		position.Quantity += quantity
		position.LastUpdateTime = trade.Time
	}

	return position, nil
}

func getCalculatedPosition(order, oldOrder pb.Order, curPos pb.CalculatedPosition) (pb.CalculatedPosition, error) {
	var err error
	if order.OrderParameters == nil {
		return pb.CalculatedPosition{}, ErrInvalidOrder
	}

	updates := GetUpdate(oldOrder, order)

	position := curPos

	for _, trade := range updates {

		quantity := getQuantity(trade.Quantity, order.OrderParameters.Side.String())

		position.AveragePrice, err = getAveragePrice(position.AveragePrice, trade.Price,
			position.Quantity, quantity)
		if err != nil {
			logrus.WithFields(logrus.Fields{"order": order, "oldOrder": oldOrder, "curPos": curPos}).
				WithError(err).Error()
		}

		position.Quantity += quantity
		position.LastUpdateTime = trade.Time
	}

	return position, nil
}
