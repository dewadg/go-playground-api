package rest

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func Register(router chi.Router) error {
	fs := http.FileServer(http.Dir(os.Getenv("WEB_DIR")))

	router.Handle("/*", fs)

	return nil
}
