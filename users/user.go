package users

import "context"

type User struct {
	ID       string `gorm:"primary_key" json:"id"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:255" json:"email"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
}

type Repository interface{
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id string) (*User, error)
}

//CREATE TABLE users (
//id VARCHAR ( 50 ) PRIMARY KEY,
//name VARCHAR ( 50 ) NOT NULL,
//username VARCHAR ( 50 ) UNIQUE NOT NULL,
//password VARCHAR ( 50 ) NOT NULL,
//email VARCHAR ( 255 )
//);