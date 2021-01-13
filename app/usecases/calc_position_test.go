package usecases

import (
	"testing"
	"time"

	pb "github.com/adnilote/order-manager/app/business/entities/proto"
	"github.com/stretchr/testify/require"
)

var (
	EmptyOrder = pb.Order{}

	OrderBuy1 = pb.Order{
		OrderParameters: &pb.OrderParameters{
			Side: pb.Side(0), // buy
		},
		OrderState: &pb.OrderState{
			Trades: []pb.Trade{
				{
					Position: 1,
					Quantity: 10,
					Price:    120.1,
					Time:     time.Now(),
				},
			},
		},
	}

	OrderBuy2 = pb.Order{
		OrderParameters: &pb.OrderParameters{
			Side: pb.Side(0),
		},
		OrderState: &pb.OrderState{
			Trades: []pb.Trade{
				{
					Position: 1,
					Quantity: 10,
					Price:    120.1,
					Time:     OrderBuy1.OrderState.Trades[0].Time,
				},
				{
					Position: 2,
					Quantity: 15,
					Price:    120.7,
					Time:     OrderBuy1.OrderState.Trades[0].Time.Add(time.Minute),
				},
			},
		},
	}

	OrderSell1 = pb.Order{
		OrderParameters: &pb.OrderParameters{
			Side: pb.Side(1),
		},
		OrderState: &pb.OrderState{
			Trades: []pb.Trade{
				{
					Position: 1,
					Quantity: 5,
					Price:    15.1,
					Time:     time.Now(),
				},
			},
		},
	}

	OrderSell2 = pb.Order{
		OrderParameters: &pb.OrderParameters{
			Side: pb.Side(1),
		},
		OrderState: &pb.OrderState{
			Trades: []pb.Trade{
				{
					Position: 1,
					Quantity: 5,
					Price:    15.1,
					Time:     OrderBuy1.OrderState.Trades[0].Time,
				},
				{
					Position: 2,
					Quantity: 15,
					Price:    14.9,
					Time:     OrderBuy1.OrderState.Trades[0].Time.Add(time.Minute),
				},
			},
		},
	}
)

func TestGetTotalAndCalculatedPositions(t *testing.T) {
	t.Parallel()

	var cases = []struct {
		name      string
		lastOrder pb.Order
		order     pb.Order
		curPos    pb.TotalPosition
		err       error
		result    pb.TotalPosition
	}{
		{
			name:      "empty order",
			lastOrder: pb.Order{},
			order:     pb.Order{},
			curPos:    pb.TotalPosition{},
			err:       ErrInvalidOrder,
			result:    pb.TotalPosition{},
		},
		{
			name:      "empty trades",
			lastOrder: pb.Order{},
			order: pb.Order{
				OrderParameters: &pb.OrderParameters{},
			},
			curPos: pb.TotalPosition{
				Quantity:       10,
				AveragePrice:   120.1,
				LastUpdateTime: OrderBuy1.OrderState.Trades[0].Time,
			},
			result: pb.TotalPosition{
				Quantity:       10,
				AveragePrice:   120.1,
				LastUpdateTime: OrderBuy1.OrderState.Trades[0].Time,
			},
		},
		{
			name:      "empty trades2",
			lastOrder: OrderSell1,
			order: pb.Order{
				OrderParameters: &pb.OrderParameters{},
			},
			curPos: pb.TotalPosition{
				Quantity:       10,
				AveragePrice:   120.1,
				LastUpdateTime: OrderBuy1.OrderState.Trades[0].Time,
			},
			result: pb.TotalPosition{
				Quantity:       10,
				AveragePrice:   120.1,
				LastUpdateTime: OrderBuy1.OrderState.Trades[0].Time,
			},
		},
		{
			name:      "increase position from zero",
			lastOrder: pb.Order{},
			order:     OrderBuy1,
			curPos:    pb.TotalPosition{},
			result: pb.TotalPosition{
				Quantity:       10,
				AveragePrice:   120.1,
				LastUpdateTime: OrderBuy1.OrderState.Trades[0].Time,
			},
		},
		{
			name:      "decrease position from zero",
			lastOrder: pb.Order{},
			order:     OrderSell1,
			curPos:    pb.TotalPosition{},
			result: pb.TotalPosition{
				Quantity:       -5,
				AveragePrice:   15.1,
				LastUpdateTime: OrderSell1.OrderState.Trades[0].Time,
			},
		},
		{
			name:      "increase long position",
			lastOrder: OrderBuy1,
			order:     OrderBuy2,
			curPos: pb.TotalPosition{
				Quantity:     10,
				AveragePrice: 120.1,
			},

			result: pb.TotalPosition{
				Quantity:       25,
				AveragePrice:   120.46,
				LastUpdateTime: OrderBuy2.OrderState.Trades[1].Time,
			},
		},
		{
			name:      "increase short position",
			lastOrder: OrderSell1,
			order:     OrderSell2,
			curPos: pb.TotalPosition{
				Quantity:     -5,
				AveragePrice: 15.1,
			},

			result: pb.TotalPosition{
				Quantity:       -20,
				AveragePrice:   14.95,
				LastUpdateTime: OrderSell2.OrderState.Trades[1].Time,
			},
		},
		{
			name:      "decrease long position",
			lastOrder: EmptyOrder,
			order:     OrderSell1,
			curPos: pb.TotalPosition{
				Quantity:     10,
				AveragePrice: 120.1,
			},

			result: pb.TotalPosition{
				Quantity:       5,
				AveragePrice:   120.1,
				LastUpdateTime: OrderSell1.OrderState.Trades[0].Time,
			},
		},
		{
			name:      "decrease short position",
			lastOrder: EmptyOrder,
			order:     OrderBuy1,
			curPos: pb.TotalPosition{
				Quantity:     -15,
				AveragePrice: 15.11,
			},
			result: pb.TotalPosition{
				Quantity:       -5,
				AveragePrice:   15.11,
				LastUpdateTime: OrderBuy1.OrderState.Trades[0].Time,
			},
		},
		{
			name:      "close long position",
			lastOrder: EmptyOrder,
			order:     OrderSell1,
			curPos: pb.TotalPosition{
				Quantity:     5,
				AveragePrice: 151.1,
			},
			result: pb.TotalPosition{
				Quantity:       0,
				AveragePrice:   151.1,
				LastUpdateTime: OrderSell1.OrderState.Trades[0].Time,
			},
		},
		{
			name:      "close short position",
			lastOrder: EmptyOrder,
			order:     OrderBuy1,
			curPos: pb.TotalPosition{
				Quantity:     -10,
				AveragePrice: 12.5,
			},
			result: pb.TotalPosition{
				Quantity:       0,
				AveragePrice:   12.5,
				LastUpdateTime: OrderBuy1.OrderState.Trades[0].Time,
			},
		},
		{
			name:      "flattern short position",
			lastOrder: EmptyOrder,
			order:     OrderBuy1,
			curPos: pb.TotalPosition{
				Quantity:     -5,
				AveragePrice: 12.5,
			},
			result: pb.TotalPosition{
				Quantity:       5,
				AveragePrice:   120.1,
				LastUpdateTime: OrderBuy1.OrderState.Trades[0].Time,
			},
		},
		{
			name:      "flattern long position",
			lastOrder: EmptyOrder,
			order:     OrderSell1,
			curPos: pb.TotalPosition{
				Quantity:     3,
				AveragePrice: 12.5,
			},
			result: pb.TotalPosition{
				Quantity:       -2,
				AveragePrice:   15.1,
				LastUpdateTime: OrderSell1.OrderState.Trades[0].Time,
			},
		},
	}

	for _, tcase := range cases {
		totalPos, err := getTotalPosition(tcase.order, tcase.lastOrder, tcase.curPos)
		require.Equal(t, tcase.err, err, "[%s]", tcase.name)
		require.Equal(t, tcase.result, totalPos, "[%s]", tcase.name)

		curCalcPos := pb.CalculatedPosition{
			Quantity:       tcase.curPos.Quantity,
			AveragePrice:   tcase.curPos.AveragePrice,
			LastUpdateTime: tcase.curPos.LastUpdateTime,
		}
		calcPos, err := getCalculatedPosition(tcase.order, tcase.lastOrder, curCalcPos)
		require.Equal(t, tcase.err, err, "[%s]", tcase.name)
		require.Equal(t, totalPos.Quantity, calcPos.Quantity, "[%s]", tcase.name)
		require.Equal(t, totalPos.AveragePrice, calcPos.AveragePrice, "[%s]", tcase.name)
		require.Equal(t, totalPos.LastUpdateTime, calcPos.LastUpdateTime, "[%s]", tcase.name)
	}
}
