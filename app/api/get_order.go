package api

import (
	pb "github.com/adnilote/order-manager/app/business/entities/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOrderUpdate(request *pb.GetOrderUpdateRequest,
	stream pb.OrderManager_GetOrderUpdateServer) error {

	uuid := uuid.Must(uuid.NewUUID()).String()
	orderCh := s.orderManager.Register(uuid)
	defer s.orderManager.Unregister(uuid)

	for order := range orderCh {
		tmpOrder := order.(pb.Order)
		err := stream.Send(&pb.GetOrderUpdateResponse{
			Order: &tmpOrder,
		})
		if err != nil {
			return status.Errorf(codes.Canceled, "error sending msg: %s", err)
		}
	}
	return nil
}
