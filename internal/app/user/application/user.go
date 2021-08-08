package application

import (
	"context"

	"layout/api/user/v1"
	"layout/internal/app/user/domain/repository"
)

// User provides use-case
type User struct {
	userRepo repository.UserRepository
	// userRepo    repository.UserRepository
}

// NewUserUseCase .
func NewUserUseCase(rep repository.UserRepository) *User {
	return &User{userRepo: rep}
}

// GetUser returns user
func (i User) GetUser(ctx context.Context, id int64) (*v1.GetUserReply, error) {
	a, err := i.userRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserReply{
		Id:   a.ID,
		Name: a.Name,
	}, nil
}

// NewUser create a user use case
func NewUser(repo repository.UserRepository) *User {
	return &User{
		userRepo: repo,
	}
}
