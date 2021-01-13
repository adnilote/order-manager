package rules

import (
	pb "github.com/adnilote/order-manager/app/business/entities/proto"
)

// IsActive true if placing = 0; pending = 1; working = 2;
func IsActive(order pb.Order) bool {
	state := order.GetOrderState()
	if state != nil {
		return state.GetStatus().String() == pb.Status_name[0] ||
			state.GetStatus().String() == pb.Status_name[1] ||
			state.GetStatus().String() == pb.Status_name[2]
	}
	return false
}
