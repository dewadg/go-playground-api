package grpc

import (
	"context"
	"fmt"
	"sync"

	"github.com/dewadg/go-playground-api/pkg/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"
)

type orchestratorService struct {
	pb.UnimplementedOrchestratorServiceServer

	clientsLock sync.Mutex
	clients     map[string]pb.ExecutionServiceClient
}

func newOrchestratorService() *orchestratorService {
	return &orchestratorService{
		clients: make(map[string]pb.ExecutionServiceClient),
	}
}

func (svc *orchestratorService) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	p, _ := peer.FromContext(ctx)
	clientAddr := p.Addr.String()

	conn, err := grpc.DialContext(
		ctx,
		clientAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	clientID := uuid.New().String() + ":" + request.GetId()
	client := pb.NewExecutionServiceClient(conn)

	svc.clientsLock.Lock()
	svc.clients[clientID] = client
	svc.clientsLock.Unlock()

	return &pb.RegisterResponse{
		Message: fmt.Sprintf("%s: registered", clientID),
	}, nil
}
