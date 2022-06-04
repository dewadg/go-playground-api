package rest

import (
	"github.com/dewadg/go-playground-api/internal/auth"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/go-chi/chi"
)

func Register(router chi.Router) error {
	webDir := os.Getenv("WEB_DIR")
	fs := http.FileServer(http.Dir(webDir))

	router.Handle("/*", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if strings.Contains(request.URL.Path, ".") || request.URL.Path == "/" {
			fs.ServeHTTP(writer, request)
			return
		}

		http.ServeFile(writer, request, path.Join(webDir, "/index.html"))
	}))

	router.Get("/v1/auth/google/login", handleGoogleLogin())
	router.Get("/v1/auth/google/callback", handleGoogleCallback(auth.GenerateAccessToken))

	return nil
}
