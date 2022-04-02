package store

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/dewadg/go-playground-api/internal/adapter"
)

func BenchmarkFindItem(b *testing.B) {
	db, err := adapter.GetMysqlDB()
	if err != nil {
		b.Fatal(err)
	}

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	data, _ := json.Marshal(CreateItemRequest{
		Input: []string{
			"Test",
		},
	})

	_, err = db.ExecContext(context.Background(), `
		INSERT INTO
			items
		SET
			id = ?,
			data = ?,
			created_at = NOW(),
			updated_at = NOW(),
			assigned_at = NOW()
	`, id, string(data))
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		_, _ = FindItem(context.Background(), id)
	}
}
