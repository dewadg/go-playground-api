package grpc

import (
	"context"
	"errors"
	"net"
	"os"

	"github.com/dewadg/go-playground-api/pkg/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Run(ctx context.Context) error {
	address := "127.0.0.1:9000"
	if os.Getenv("APP_ENV") == "production" {
		address = "0.0.0.0:9000"
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterOrchestratorServiceServer(server, newOrchestratorService())

	go func() {
		<-ctx.Done()

		logrus.Info("shutting down grpc server")
		server.GracefulStop()
		logrus.Info("grpc server shut down")
	}()

	logrus.WithField("address", address).Info("grpc server started")
	if err = server.Serve(listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		return err
	}

	return nil
}
