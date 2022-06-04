package resolver

import (
	"context"
	"strconv"
	"time"

	"github.com/dewadg/go-playground-api/internal/auth"
	"github.com/graph-gophers/graphql-go"
)

type User struct {
	ID        graphql.ID `json:"id"`
	Email     string     `json:"email"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
}

func (r *resolver) Whoami(ctx context.Context) (User, error) {
	user, err := auth.Whoami(ctx)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:        graphql.ID(strconv.FormatInt(user.ID, 10)),
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}
