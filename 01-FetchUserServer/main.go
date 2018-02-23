package main

import (
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := userService{}

	fetchUserHandler := httptransport.NewServer(
		makeFetchUserEndpoint(svc),
		decodeFetchUserRequest,
		encodeResponse,
	)

	http.Handle("/fetchUser", fetchUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
