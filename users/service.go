package users

import "context"

type Service interface {
	CreateUser(ctx context.Context, u User) (*User, error)
	UpdateUser(ctx context.Context, u User, id string) (*User, error)
	GetAllUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}
