package service

import (
	"context"

	"layout/api/user/v1"
	"layout/internal/app/user/application"
)

// User provides use-case
type User struct {
	*v1.UnimplementedUserServer
	UseCase *application.User
}

// GetUser returns user
func (u User) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {

	a, err := u.UseCase.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return a, nil
}
