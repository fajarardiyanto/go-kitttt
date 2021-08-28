package users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoint struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoint {
	return Endpoint{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser: makeGetUserEnpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return  func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(User)
		ok, err := s.CreateUser(ctx, req)
		return ok, err
	}
}

func makeGetUserEnpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(User)
		user, err := s.GetUser(ctx, req.ID)
		return user, err
	}
}
