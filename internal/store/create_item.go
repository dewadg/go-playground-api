package store

import (
	"context"
	"encoding/json"

	"github.com/dewadg/go-playground-api/internal/adapter"
)

type CreateItemRequest struct {
	Input []string
}

type CreateItemResponse struct {
	ShareID string
}

// TODO: implement queue on this
func CreateItem(ctx context.Context, payload CreateItemRequest) (CreateItemResponse, error) {
	db, err := adapter.GetMysqlDB()
	if err != nil {
		return CreateItemResponse{}, err
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return CreateItemResponse{}, err
	}

	id, err := popID(ctx)
	if err != nil {
		return CreateItemResponse{}, err
	}

	query := `UPDATE items SET data = ?, updated_at = NOW() WHERE id = ?`
	_, err = db.ExecContext(ctx, query, string(payloadBytes), id)
	if err != nil {
		return CreateItemResponse{}, err
	}

	return CreateItemResponse{
		ShareID: id,
	}, nil
}
