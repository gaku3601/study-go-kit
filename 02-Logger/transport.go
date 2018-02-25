package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func makeFetchUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(fetchUserRequest)
		v := svc.FetchUser(ctx, req.ID)
		return fetchUserResponse{v.ID, v.Name}, nil
	}
}

type fetchUserRequest struct {
	ID int `json:"id"`
}

type fetchUserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func decodeFetchUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request fetchUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
