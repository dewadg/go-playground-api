package store

import (
	"context"
	"os"

	"github.com/dewadg/go-playground-api/internal/pkg/adapter"
	"github.com/sirupsen/logrus"
)

type itemGetter interface {
	Get(ctx context.Context, id string) (Item, error)
}

type itemSetter interface {
	Set(ctx context.Context, payload Item) (Item, error)
}

type itemClearer interface {
	Clear(ctx context.Context, id string) error
}

var (
	primaryGetter itemGetter
	cacheGetter   itemGetter

	primarySetter itemSetter
	cacheSetter   itemSetter

	cacheClearer itemClearer
)

func init() {
	if os.Getenv("CLI_MODE") == "true" {
		logrus.Info("store: skip initializing dependencies in cli mode")

		return
	}

	db, err := adapter.GetSqliteDB()
	if err != nil {
		logrus.WithError(err).Fatal("store: failed to init sqlite db")
	}

	redisClient, err := adapter.GetRedisClient()
	if err != nil {
		logrus.WithError(err).Fatal("store: failed to init redis client")
	}

	sqliteRepo := newSqliteRepository(db)

	primaryGetter = sqliteRepo
	primarySetter = sqliteRepo

	redisRepo := newRedisRepository(redisClient)

	cacheSetter = redisRepo
	cacheGetter = redisRepo
	cacheClearer = redisRepo
}
