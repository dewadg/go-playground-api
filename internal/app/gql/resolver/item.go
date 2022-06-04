package resolver

import (
	"context"

	"github.com/dewadg/go-playground-api/internal/store"
	"github.com/graph-gophers/graphql-go"
)

type CreateItemRequest struct {
	Input []string
}

type CreateItemResponse struct {
	ShareID graphql.ID
}

func (r *resolver) CreateItem(ctx context.Context, args struct{ Payload CreateItemRequest }) (CreateItemResponse, error) {
	result, err := store.CreateItem(ctx, store.CreateItemRequest{
		Input: args.Payload.Input,
	})
	if err != nil {
		return CreateItemResponse{}, err
	}

	return CreateItemResponse{
		ShareID: graphql.ID(result.ShareID),
	}, nil
}

type Item struct {
	ID    graphql.ID
	Input []string
}

func (r *resolver) Item(ctx context.Context, args struct{ ID graphql.ID }) (Item, error) {
	item, err := store.FindItem(ctx, string(args.ID))
	if err != nil {
		return Item{}, err
	}

	return Item{
		ID:    graphql.ID(item.ID),
		Input: item.Input,
	}, nil
}

type UpdateItemRequest struct {
	Input []string
}

type UpdateItemResponse struct {
	ShareID graphql.ID
}

func (r *resolver) UpdateItem(ctx context.Context, args struct {
	ShareID string
	Payload CreateItemRequest
}) (Item, error) {
	result, err := store.UpdateItem(ctx, args.ShareID, store.UpdateItemRequest{
		Input: args.Payload.Input,
	})
	if err != nil {
		return Item{}, err
	}

	return Item{
		ID:    graphql.ID(result.ID),
		Input: result.Input,
	}, nil
}
