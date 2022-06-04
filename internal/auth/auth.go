package auth

import (
	"context"
	"os"

	"github.com/dewadg/go-playground-api/internal/pkg/adapter"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var signingKey []byte
var authRepo *sqliteAuthRepository

func init() {
	signingKey = []byte(os.Getenv("JWT_SIGNING_KEY"))

	db, err := adapter.GetSqliteDB()
	if err != nil {
		logrus.WithError(err).Fatal("auth: failed to init sqlite db")
	}

	authRepo = newSqliteAuthRepository(db)
}

func GenerateAccessToken(ctx context.Context, email string) (string, error) {
	return authRepo.GenerateAccessToken(ctx, email)
}

func GenerateJWT(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(signingKey)
}
