package httputil

import (
	"context"
	"net/http"
	"strings"

	"github.com/dewadg/go-playground-api/internal/pkg/entity"
	"github.com/dgrijalva/jwt-go"
)

type JWTDecoder func(s string) (jwt.MapClaims, error)

func AuthMiddleware(jwtDec JWTDecoder) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken := r.Header.Get("Authorization")
			if len(accessToken) == 0 {
				next.ServeHTTP(w, r)
				return
			}
			splitAccessToken := strings.Split(accessToken, " ")
			if len(splitAccessToken) < 2 {
				http.Error(w, "invalid access token", http.StatusUnauthorized)
				return
			}

			claims, err := jwtDec(splitAccessToken[1])
			if err != nil {
				http.Error(w, "invalid access token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), entity.KeyUserID, int64(claims["user_id"].(float64)))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
