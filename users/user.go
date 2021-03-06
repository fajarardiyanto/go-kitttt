package users

import "context"

type User struct {
	ID       string `gorm:"primary_key" json:"id"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:255" json:"email"`
	Username string `gorm:"size:255;index:username;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User, id string) error
	GetAllUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

//CREATE TABLE users (
//id VARCHAR ( 50 ) PRIMARY KEY,
//name VARCHAR ( 50 ) NOT NULL,
//username VARCHAR ( 50 ) UNIQUE NOT NULL,
//password VARCHAR ( 50 ) NOT NULL,
//email VARCHAR ( 255 )
//);
