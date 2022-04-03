package serve

import (
	"net/http"
	"os"

	"github.com/dewadg/go-playground-api/internal/gql"
	"github.com/dewadg/go-playground-api/internal/rest"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			router := chi.NewRouter()
			router.Use(cors.Handler(cors.Options{
				AllowedOrigins: []string{"https://*", "http://*"},
				AllowedMethods: []string{"GET", "POST", "OPTIONS"},
			}))

			if err := gql.Register(router); err != nil {
				logrus.WithError(err).Fatal("failed to register gql handler")
			}

			if err := rest.Register(router); err != nil {
				logrus.WithError(err).Fatal("failed to register rest handler")
			}

			address := "127.0.0.1:8000"
			if os.Getenv("APP_ENV") == "production" {
				address = "0.0.0.0:8000"
			}

			logrus.Info("starting http server at ", address)
			if err := http.ListenAndServe(address, router); err != nil {
				logrus.WithError(err).Fatal("failed to start http server")
			}
		},
	}
}
