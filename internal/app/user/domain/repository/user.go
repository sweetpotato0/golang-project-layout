package repository

import (
	"context"

	"layout/internal/app/user/domain/entity"
)

// UserRepository represent repository of the user
// Expect implementation by the infrastructure layer
type UserRepository interface {
	GetUser(ctx context.Context, id int64) (*entity.User, error)
}
