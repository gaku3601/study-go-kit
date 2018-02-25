package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var svc UserService
	svc = userService{}
	logger := log.NewLogfmtLogger(os.Stderr)
	svc = loggingMiddleware{logger, svc}

	fetchUserHandler := httptransport.NewServer(
		makeFetchUserEndpoint(svc),
		decodeFetchUserRequest,
		encodeResponse,
	)

	http.Handle("/fetchUser", fetchUserHandler)
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
