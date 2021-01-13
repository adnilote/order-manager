package usecases

import (
	pb "github.com/adnilote/order-manager/app/business/entities/proto"
)

func abs(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}

func sgn(a float64) float64 {
	if a < 0 {
		return -1
	}
	if a > 0 {
		return 1
	}
	return 0
}

func getQuantity(absQuantity float64, side string) float64 {
	if side == "sell" {
		return -absQuantity

	}
	return absQuantity
}

func GetUpdate(oldOrder pb.Order, newOrder pb.Order) []pb.Trade {
	var lenOld, lenNew int
	if oldOrder.GetOrderState() == nil {
		lenOld = 0
	} else {
		lenOld = len(oldOrder.OrderState.Trades)
	}
	if newOrder.GetOrderState() == nil {
		lenNew = 0
	} else {
		lenNew = len(newOrder.OrderState.Trades)
	}

	if lenNew > lenOld {
		return newOrder.OrderState.Trades[lenOld:]
	}
	return []pb.Trade{}
}
