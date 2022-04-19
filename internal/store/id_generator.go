package store

import (
	"context"
	"database/sql"
	"math/rand"
	"strings"
	"time"

	"github.com/dewadg/go-playground-api/internal/adapter"
	"github.com/sirupsen/logrus"
)

var allowedChars = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func generateIDs(ctx context.Context, numOfIDs int, length int) chan string {
	reservedIDChan := make(chan string)
	reservedIDsTracker := make(map[string]bool)

	generateID := func() string {
		id := make([]rune, length)
		for i := range id {
			id[i] = allowedChars[rand.Intn(len(allowedChars))]
		}

		return string(id)
	}

	go func() {
		rand.Seed(time.Now().UnixNano())

		for i := 0; i < numOfIDs; i++ {
			select {
			case <-ctx.Done():
				return
			default:
			}

			id := generateID()
			if _, exists := reservedIDsTracker[id]; exists {
				i--
				continue
			}

			reservedIDsTracker[id] = true
			reservedIDChan <- id
		}

		close(reservedIDChan)
	}()

	return reservedIDChan
}

func SeedIDs(ctx context.Context, numOfIDs int, length int) error {
	db, err := adapter.GetSqliteDB()
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
				items(id, created_at)
			VALUES
				(?, ?)
		`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	idChan := generateIDs(ctx, numOfIDs, length)
	i := 1
	now := time.Now().UTC().Format("2006-01-02 15:04:05")
	for id := range idChan {
		_, err = stmt.ExecContext(ctx, id, now)
		if err != nil {
			if strings.HasPrefix(err.Error(), "Error 1062") {
				continue
			}
			return err
		}

		logrus.Infof("%d id inserted", i)
		i++
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func popID(ctx context.Context) (string, error) {
	db, err := adapter.GetSqliteDB()
	if err != nil {
		return "", err
	}

	query := `SELECT id FROM items WHERE assigned_at IS NULL ORDER BY assigned_at ASC LIMIT 1`

	var id string
	if err = db.QueryRowContext(ctx, query).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
