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

	query := `UPDATE items SET data = ?, updated_at = NOW(), assigned_at = NOW() WHERE id = ?`
	_, err = db.ExecContext(ctx, query, string(payloadBytes), id)
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
	db, err := adapter.GetMysqlDB()
	if err != nil {
		return Item{}, err
	}

	query := `SELECT data FROM items WHERE id = ? LIMIT 1`

	var dataBytes []byte
	if err = db.QueryRowContext(ctx, query, id).Scan(&dataBytes); err != nil {
		return Item{}, err
	}

	var item Item
	if err = json.Unmarshal(dataBytes, &item); err != nil {
		return Item{}, err
	}

	item.ID = id

	return item, nil
}
