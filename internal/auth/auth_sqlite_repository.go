package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dewadg/go-playground-api/internal/pkg/entity"
)

type sqliteAuthRepository struct {
	db *sql.DB
}

func newSqliteAuthRepository(db *sql.DB) *sqliteAuthRepository {
	return &sqliteAuthRepository{
		db: db,
	}
}

func (r *sqliteAuthRepository) GenerateAccessToken(ctx context.Context, email string) (string, error) {
	user, err := r.findUserByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return "", err
		}

		if err = r.createUser(ctx, email); err != nil {
			return "", err
		}

		user, err = r.findUserByEmail(ctx, email)
	}

	accessToken, err := GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	query := `
		INSERT INTO
			access_tokens(token, created_at, updated_at)
		VALUES
			(?, ?, ?)
	`

	now := time.Now().UTC()
	_, err = r.db.ExecContext(ctx, query, accessToken, now, now)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (r *sqliteAuthRepository) findUserByEmail(ctx context.Context, email string) (entity.User, error) {
	query := `
		SELECT
			id,
			email,
			created_at,
			updated_at
		FROM users
		WHERE
			email = ?
			AND deleted_at IS NULL
		LIMIT 1`

	var user entity.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}

func (r *sqliteAuthRepository) createUser(ctx context.Context, email string) error {
	query := `
		INSERT INTO
			users(email, created_at, updated_at)
		VALUES
			(?, ?, ?)
	`

	now := time.Now().UTC()
	_, err := r.db.ExecContext(ctx, query, email, now, now)

	return err
}

func (r *sqliteAuthRepository) findUserByID(ctx context.Context, id int64) (entity.User, error) {
	query := `
		SELECT
			id,
			email,
			created_at,
			updated_at
		FROM users
		WHERE
			id = ?
			AND deleted_at IS NULL
		LIMIT 1`

	var user entity.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}
