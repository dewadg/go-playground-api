package auth

import (
	"context"
	"errors"
	"os"

	"github.com/dewadg/go-playground-api/internal/pkg/adapter"
	"github.com/dewadg/go-playground-api/internal/pkg/entity"
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

func ValidateJWT(s string) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	_, err := jwt.ParseWithClaims(s, &claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	return claims, err
}

func Whoami(ctx context.Context) (entity.User, error) {
	userID, ok := ctx.Value(entity.KeyUserID).(int64)
	if !ok {
		return entity.User{}, errors.New("missing user on context")
	}

	return authRepo.findUserByID(ctx, userID)
}
