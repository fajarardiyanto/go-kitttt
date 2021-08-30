package users

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
)

type service struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

func (s *service) CreateUser(ctx context.Context, u User) (*User, error) {
	logger := log.With(s.logger, "POST", "CreateUser")

	uid := uuid.New().String()
	user := User{
		ID:       uid,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}

	if err := s.repostory.CreateUser(ctx, &user); err != nil {
		level.Error(logger).Log("err", err.Error())
		return nil, err
	}

	logger.Log("Create User", user)

	return &user, nil
}

func (s *service) UpdateUser(ctx context.Context, u User, id string) (*User, error) {
	logger := log.With(s.logger, "POST", "CreateUser")

	user := User{
		ID:       id,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}

	if err := s.repostory.UpdateUser(ctx, &user, id); err != nil {
		level.Error(logger).Log("err", err.Error())
		return nil, err
	}

	logger.Log("Update User", user)

	return &user, nil
}

func (s *service) GetAllUsers(ctx context.Context) (*[]User, error) {
	logger := log.With(s.logger, "GET", "GetAllUser")

	user, err := s.repostory.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	logger.Log("Get All User")

	return user, err
}

func (s *service) GetUser(ctx context.Context, id string) (*User, error) {
	logger := log.With(s.logger, "GET", "GetUser")

	user, err := s.repostory.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	logger.Log("Get User", id)

	return user, err
}

func (s *service) DeleteUser(ctx context.Context, id string) error {
	logger := log.With(s.logger, "GET", "DeleteUser")

	err := s.repostory.DeleteUser(ctx, id)


	logger.Log("Delete User")

	return err
}
