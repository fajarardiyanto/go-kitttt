package users

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/jinzhu/gorm"
)

var RepoErr = errors.New("unable to Handle Repo Request")

type repo struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepo(db *gorm.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (r *repo) CreateUser(ctx context.Context, user *User) error {
	if user == nil {
		return RepoErr
	}

	if err := r.db.Debug().Model(&User{}).Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateUser(ctx context.Context, user *User, id string) error {
	var err error
	if err = r.db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).UpdateColumns(&user).Error; err != nil {
		return err
	}

	return err
}

func (r *repo) GetAllUsers(ctx context.Context) (*[]User, error) {
	var user []User

	var err error
	if err = r.db.Debug().Model(&User{}).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repo) GetUser(ctx context.Context, id string) (*User, error) {
	var user User

	var err error
	if err = r.db.Debug().Model(&User{}).Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repo) DeleteUser(ctx context.Context, id string) error {
	var user User

	var err error
	if err = r.db.Debug().Model(&User{}).Where("id = ?", id).Take(&user).Delete(&User{}).Error; err != nil {
		return err
	}

	return err
}
