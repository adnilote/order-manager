package entities

import (
	"fmt"

	pb "github.com/adnilote/order-manager/app/business/entities/proto"
)

func GetTotalPositionKey(p pb.TotalPosition) string {
	return fmt.Sprintf(
		"%s/%s",
		p.Account, p.Instrument.Symbol,
	)
}

func GetCalculatedPositionKey(p pb.CalculatedPosition) string {
	return fmt.Sprintf(
		"%s/%s/%s",
		p.Account, p.Strategy, p.Instrument.Symbol,
	)
}
