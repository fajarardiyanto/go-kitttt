package users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoint struct {
	CreateUser  endpoint.Endpoint
	GetAllUsers endpoint.Endpoint
	GetUser     endpoint.Endpoint
	DeleteUser  endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoint {
	return Endpoint{
		CreateUser:  makeCreateUserEndpoint(s),
		GetAllUsers: makeGetAllUserEnpoint(s),
		GetUser:     makeGetUserEnpoint(s),
		DeleteUser:  makeDeleteUserEnpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(User)
		ok, err := s.CreateUser(ctx, req)

		res := map[string]interface{}{
			"error": false,
			"message": "",
			"data": ok,
		}

		return res, err
	}
}

func makeGetAllUserEnpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		user, err := s.GetAllUsers(ctx)

		res := map[string]interface{}{
			"error": false,
			"message": "",
			"data": user,
		}

		return res, err
	}
}

func makeGetUserEnpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(User)
		user, err := s.GetUser(ctx, req.ID)

		res := map[string]interface{}{
			"error": false,
			"message": "",
			"data": user,
		}

		return res, err
	}
}

func makeDeleteUserEnpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(User)
		err := s.DeleteUser(ctx, req.ID)

		res := map[string]interface{}{
			"error": false,
			"message": "Success Delete User",
			"data": "",
		}

		return res, err
	}
}
