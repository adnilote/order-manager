package entities

import (
	"encoding/json"
	"fmt"

	pb "github.com/adnilote/order-manager/app/business/entities/proto"
)

func GetOrderKey(o pb.Order) string {
	strategy, _ := ParseClientTag(o.ClientTag)
	return fmt.Sprintf(
		"%s/%s/%s/%s",
		o.Account, strategy, o.OrderParameters.Instrument.Symbol, o.OrderId,
	)
}

type ClientTag struct {
	Strategy string `json:"strategy"`
	Event    string `json:"event"`
}

func ParseClientTag(clientTag string) (strategy string, event string) {
	var tag ClientTag
	err := json.Unmarshal([]byte(clientTag), &tag)
	if err != nil {
		strategy = clientTag
		return
	}
	return tag.Strategy, tag.Event
}
