package users

import "context"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
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