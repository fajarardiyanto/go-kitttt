package users

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPServer(ctx context.Context, endpoint Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoint.CreateUser,
		decodeRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/user/update/{id}").Handler(httptransport.NewServer(
		endpoint.UpdateUser,
		decodeUpdateRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/users").Handler(httptransport.NewServer(
		endpoint.GetAllUsers,
		decodeAllUserReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoint.GetUser,
		decodeUserReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/delete/{id}").Handler(httptransport.NewServer(
		endpoint.DeleteUser,
		decodeUserReq,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
