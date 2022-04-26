package serve

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/dewadg/go-playground-api/internal/grpc"
	"github.com/dewadg/go-playground-api/internal/rest"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, stop := context.WithCancel(context.Background())

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt, os.Kill)

			go func() {
				if err := rest.Run(ctx); err != nil {
					logrus.WithError(err).Fatal("failed to start http server")
				}
			}()

			go func() {
				if err := grpc.Run(ctx); err != nil {
					logrus.WithError(err).Fatal("failed to start grpc server")
				}
			}()

			<-sigChan
			stop()

			if os.Getenv("APP_ENV") == "production" {
				logrus.Info("gracefully shutting down for 5 seconds...")
				time.Sleep(5 * time.Second)
				logrus.Info("shutdown completed")
			}
		},
	}
}
