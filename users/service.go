package users

import "context"

type Service interface {
	CreateUser(ctx context.Context, u User) (*User, error)
	GetAllUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}
