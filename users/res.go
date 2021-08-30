package users

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req User
	vars := mux.Vars(r)

	req = User{
		ID: vars["id"],
		Username: req.Username,
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req User
	vars := mux.Vars(r)

	req = User{
		ID: vars["id"],
	}
	return req, nil
}

func decodeAllUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req []User

	return req, nil
}

func decodeDeleteUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req User
	vars := mux.Vars(r)

	req = User{
		ID: vars["id"],
	}
	fmt.Println(req)
	return "Success Delete User", nil
}

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
