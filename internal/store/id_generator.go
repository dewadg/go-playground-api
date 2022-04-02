package store

import (
	"context"
	"database/sql"
	"math/rand"
	"time"

	"github.com/dewadg/go-playground-api/internal/adapter"
)

var allowedChars = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func generateIDs(numOfIDs int, length int) []string {
	reservedIDs := make([]string, 0)
	reservedIDsTracker := make(map[string]bool)

	rand.Seed(time.Now().UnixNano())
	generateID := func() string {
		id := make([]rune, length)
		for i := range id {
			id[i] = allowedChars[rand.Intn(len(allowedChars))]
		}

		return string(id)
	}

	for i := 0; i < numOfIDs; i++ {
		id := generateID()
		if _, exists := reservedIDsTracker[id]; exists {
			continue
		}

		reservedIDsTracker[id] = true
		reservedIDs = append(reservedIDs, id)
	}

	return reservedIDs
}

func SeedIDs(ctx context.Context) error {
	db, err := adapter.GetMysqlDB()
	if err != nil {
		return err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
			INSERT INTO
				items
			SET
				id = ?,
				created_at = NOW(),
				updated_at = NOW()
		`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	ids := generateIDs(1000, 5)
	for _, id := range ids {
		_, err = stmt.ExecContext(ctx, id)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func popID(ctx context.Context) (string, error) {
	db, err := adapter.GetMysqlDB()
	if err != nil {
		return "", err
	}

	query := `SELECT id FROM items ORDER BY updated_at ASC LIMIT 1`

	var id string
	if err = db.QueryRowContext(ctx, query).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
