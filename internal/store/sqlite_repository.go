package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

type sqliteRepository struct {
	db *sql.DB
}

func newSqliteRepository(db *sql.DB) *sqliteRepository {
	return &sqliteRepository{db: db}
}

func (s *sqliteRepository) Set(ctx context.Context, payload Item) (Item, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return Item{}, err
	}

	query := `UPDATE items SET data = ?, assigned_at = ? WHERE id = ?`
	_, err = s.db.ExecContext(ctx, query, string(payloadBytes), time.Now().Format("2006-01-02 15:04:05"), payload.ID)
	if err != nil {
		return Item{}, err
	}

	return payload, nil
}

func (s *sqliteRepository) Get(ctx context.Context, id string) (Item, error) {
	query := `SELECT data FROM items WHERE id = ? LIMIT 1`

	var dataBytes []byte
	if err := s.db.QueryRowContext(ctx, query, id).Scan(&dataBytes); err != nil {
		return Item{}, err
	}

	var item Item
	if err := json.Unmarshal(dataBytes, &item); err != nil {
		return Item{}, err
	}

	item.ID = id

	return item, nil
}
