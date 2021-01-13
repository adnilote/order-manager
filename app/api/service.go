package api

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/adnilote/order-manager/app/business/entities/proto"
	"github.com/adnilote/order-manager/app/config"
	"github.com/adnilote/order-manager/app/usecases"
	"google.golang.org/grpc"
)

type Service struct {
	ctx        context.Context
	grpcServer *grpc.Server
}
type Server struct {
	orderManager interface {
		Register(wsuuid string) chan interface{}
		Unregister(wsuuid string)
	}
}

func NewService(ctx context.Context, us *usecases.Usecases) *Service {

	srvc := Service{
		ctx: ctx,
	}

	var opts []grpc.ServerOption
	srvc.grpcServer = grpc.NewServer(opts...)

	pb.RegisterOrderManagerServer(srvc.grpcServer, &Server{
		orderManager: us.OrderManager,
	})

	return &srvc
}

func (srvc *Service) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Config.HTTP.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = srvc.grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

func (srvc *Service) Close() {
	srvc.grpcServer.GracefulStop()
}
