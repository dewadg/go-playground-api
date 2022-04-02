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
