package users

import "context"

type Service interface {
	CreateUser(ctx context.Context, u User) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
}
