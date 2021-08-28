package users

import (
	"context"
	"encoding/json"
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

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req User
	vars := mux.Vars(r)

	req = User{
		ID: vars["id"],
	}
	return req, nil
}
