package users

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("unable to Handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db: db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (r *repo) CreateUser(ctx context.Context, user *User) error {
	if user == nil {
		return RepoErr
	}

	if _, err := r.db.ExecContext(ctx, createUserQuery, user.ID, user.Name, user.Email, user.Username, user.Password); err != nil {
		return err
	}

	return nil
}

func (r *repo) GetUser(ctx context.Context, id string) (*User, error) {
	var user User
	if err := r.db.QueryRow(getUserQuery, id).Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}
