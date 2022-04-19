package store

import (
	"context"

	"github.com/sirupsen/logrus"
)

type CreateItemRequest struct {
	Input []string
}

type UpdateItemRequest struct {
	Input []string
}

type CreateItemResponse struct {
	ShareID string
}

type UpdateItemResponse struct {
	ShareID string
}

// TODO: implement queue on this
func CreateItem(ctx context.Context, payload CreateItemRequest) (CreateItemResponse, error) {
	id, err := popID(ctx)
	if err != nil {
		return CreateItemResponse{}, err
	}

	_, err = primarySetter.Set(ctx, Item{
		ID:    id,
		Input: payload.Input,
	})
	if err != nil {
		return CreateItemResponse{}, err
	}

	return CreateItemResponse{
		ShareID: id,
	}, nil
}

type Item struct {
	ID    string
	Input []string
}

func FindItem(ctx context.Context, id string) (Item, error) {
	cachedItem, err := cacheGetter.Get(ctx, id)
	if err == nil {
		return cachedItem, nil
	} else {
		logrus.WithError(err).Error("store.FindItem: failed to fetch item cache")
	}

	item, err := primaryGetter.Get(ctx, id)
	if err != nil {
		return Item{}, err
	}

	_, err = cacheSetter.Set(ctx, item)
	if err != nil {
		logrus.WithError(err).Error("store.FindItem: failed to cache item")
	}

	return item, nil
}

func UpdateItem(ctx context.Context, shareID string, payload UpdateItemRequest) (Item, error) {
	item, err := primarySetter.Set(ctx, Item{
		ID:    shareID,
		Input: payload.Input,
	})

	if err = cacheClearer.Clear(ctx, shareID); err != nil {
		logrus.WithError(err).Error("error while removing item from redis after update")
	}

	return item, nil
}
