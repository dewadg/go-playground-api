package grpc

import (
	"context"
	"fmt"

	"github.com/dewadg/go-playground-api/pkg/pb"
	"google.golang.org/grpc/peer"
)

type orchestratorService struct {
	pb.UnimplementedOrchestratorServiceServer

	clients map[string]pb.ExecutionServiceClient
}

func newOrchestratorService() *orchestratorService {
	return &orchestratorService{
		clients: make(map[string]pb.ExecutionServiceClient),
	}
}

func (svc *orchestratorService) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	p, _ := peer.FromContext(ctx)
	clientAddr := p.Addr.String()

	return &pb.RegisterResponse{
		Message: fmt.Sprintf("%s: registered", clientAddr),
	}, nil
}
