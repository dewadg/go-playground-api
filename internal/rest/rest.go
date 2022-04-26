package rest

import (
	"context"
	"errors"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/dewadg/go-playground-api/internal/gql"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context) error {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
	}))

	if err := gql.Register(router); err != nil {
		return errors.New("failed to register gql handler")
	}

	webDir := os.Getenv("WEB_DIR")
	fs := http.FileServer(http.Dir(webDir))

	router.Handle("/*", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if strings.Contains(request.URL.Path, ".") || request.URL.Path == "/" {
			fs.ServeHTTP(writer, request)
			return
		}

		http.ServeFile(writer, request, path.Join(webDir, "/index.html"))
	}))

	address := "127.0.0.1:8000"
	if os.Getenv("APP_ENV") == "production" {
		address = "0.0.0.0:8000"
	}

	server := http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		<-ctx.Done()

		logrus.Info("shutting down http server")
		err := server.Shutdown(context.Background())
		if err != nil {
			logrus.WithError(err).Error("error while shutting down http server")
		} else {
			logrus.Info("http server shut down")
		}
	}()

	logrus.WithField("address", address).Info("http server started")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
