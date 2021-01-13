package usecases

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	pb "github.com/adnilote/order-manager/app/business/entities/proto"
	"github.com/adnilote/order-manager/app/config"
	storelib "github.com/adnilote/order-manager/app/store"
	"github.com/adnilote/order-manager/app/usecases/mock"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const TimeFormat = "2006-01-02T15:04:05.000Z"

var store Store

func init() {
	configFileName := os.Getenv("TEST_CONFIG")
	if configFileName == "" {
		configFileName = "config-test.yaml"
	}

	err := config.LoadConfig(configFileName)
	if err != nil {
		log.Fatal("Load config failed: ", err)
	}

	store, err = storelib.NewStore(context.TODO())
	if err != nil {
		log.Fatal("Load store failed: ", err)
	}
}

// save positions to store
func TestOrderHandler(t *testing.T) {
	mc := minimock.NewController(t)

	var cases = []struct {
		name                 string
		order                string
		curPos               pb.TotalPosition
		err                  error
		resultPos            pb.TotalPosition
		resultLastUpdateTime string
	}{
		{
			name:  "order without trades",
			order: `{"currentModificationId":"3a7ea288-0d36-4a5f-931c-2d370738a063","placeTime":"2020-12-20T15:55:49.559Z","username":"aa@ya.ru","orderId":"3a7ea288-0d36-4a5f-931c-2d370738a063","orderState":{"lastUpdateTime":"2020-12-20T15:55:49.559Z","status":"rejected","reason":"Shorting is not allowed (position: 3, active: (Limit=0/0, Stop=0/0))","trades":[]},"orderParameters":{"instrument":{"symbol":"BAT.EXANTE"},"orderType":"market","side":"sell","duration":"good_till_cancel","quantity":8},"account":"ONJ0303.001"}`,
			resultPos: pb.TotalPosition{
				Account: "ONJ0303.001",
				Instrument: &pb.Instrument{
					Symbol: "BAT.EXANTE",
				},
			},
		},
		{
			name:  "order1",
			order: `{"currentModificationId":"994a8cfb-6229-4827-b790-8172f664d982","placeTime":"2020-12-20T15:55:16.384Z","username":"aa@ya.ru","orderId":"994a8cfb-6229-4827-b790-8172f664d982","orderState":{"lastUpdateTime":"2020-12-20T15:55:16.581Z","status":"filled","trades":[{"time":"2020-12-20T15:55:16.581Z","quantity":8,"price":0.242664}]},"orderParameters":{"instrument":{"symbol":"BAT.EXANTE"},"orderType":"market","side":"sell","duration":"good_till_cancel","quantity":8},"account":"ONJ0303.001"}`,
			resultPos: pb.TotalPosition{
				Account: "ONJ0303.001",
				Instrument: &pb.Instrument{
					Symbol: "BAT.EXANTE",
				},
				Quantity:     -8,
				AveragePrice: 0.242664,
			},
			resultLastUpdateTime: "2020-12-20T15:55:16.581Z",
		},
		{
			name:  "order2",
			order: `{"currentModificationId":"994a8cfb-6229-4827-b790-8172f664d982","placeTime":"2020-12-20T15:55:16.384Z","username":"aa@ya.ru","orderId":"994a8cfb-6229-4827-b790-8172f664d982","orderState":{"lastUpdateTime":"2020-12-20T15:55:16.581Z","status":"filled","trades":[{"time":"2020-12-20T15:55:16.581Z","quantity":8,"price":0.242664}]},"orderParameters":{"instrument":{"symbol":"BAT.EXANTE"},"orderType":"market","side":"sell","duration":"good_till_cancel","quantity":8},"account":"ONJ0303.001"}`,
			resultPos: pb.TotalPosition{
				Account: "ONJ0303.001",
				Instrument: &pb.Instrument{
					Symbol: "BAT.EXANTE",
				},
				Quantity:     -16,
				AveragePrice: 0.242664,
			},
			resultLastUpdateTime: "2020-12-20T15:55:16.581Z",
		},
		{
			name:  "order3",
			order: `{"currentModificationId":"994a8cfb-6229-4827-b790-8172f664d982","placeTime":"2020-12-20T15:55:16.384Z","username":"aa@ya.ru","orderId":"994a8cfb-6229-4827-b790-8172f664d982","orderState":{"lastUpdateTime":"2020-12-20T15:55:16.581Z","status":"filled","trades":[{"time":"2020-12-20T15:55:16.581Z","quantity":16,"price":0.242664}]},"orderParameters":{"instrument":{"symbol":"BAT.EXANTE"},"orderType":"market","side":"buy","duration":"good_till_cancel","quantity":8},"account":"ONJ0303.001"}`,
			resultPos: pb.TotalPosition{
				Account: "ONJ0303.001",
				Instrument: &pb.Instrument{
					Symbol: "BAT.EXANTE",
				},
				Quantity:     0,
				AveragePrice: 0.242664,
			},
			resultLastUpdateTime: "2020-12-20T15:55:16.581Z",
		},
	}

	handler := OrderHandler{
		store:         store,
		orderStreamer: mock.NewEventStreamerMock(mc),
	}
	for _, tcase := range cases {
		msg := &kafka.Message{
			Value: []byte(tcase.order),
		}

		err := handler.Handle(msg)
		require.Nil(t, err, "[%s]", tcase.name)

		totalPos, err := handler.store.GetTotalPosition(
			tcase.resultPos.Account, tcase.resultPos.Instrument.Symbol)
		require.Nil(t, err, "[%s]", tcase.name)

		if tcase.resultLastUpdateTime != "" {
			tcase.resultPos.LastUpdateTime, err = time.Parse(TimeFormat, tcase.resultLastUpdateTime)
		}
		require.Nil(t, err, "[%s]", tcase.name)
		require.Equal(t, tcase.resultPos, totalPos, "[%s]", tcase.name)

	}
}
