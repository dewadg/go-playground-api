package main

import (
	"net/http"
	"os"

	"github.com/dewadg/go-playground-api/internal/gql"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

func main() {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	if err := gql.Register(router); err != nil {
		logrus.WithError(err).Fatal("failed to register gql handler")
	}

	address := "127.0.0.1:8000"
	if os.Getenv("APP_ENV") == "production" {
		address = "0.0.0.0:8000"
	}

	logrus.Info("starting http server at ", address)
	if err := http.ListenAndServe(address, router); err != nil {
		logrus.WithError(err).Fatal("failed to start http server")
	}
}
