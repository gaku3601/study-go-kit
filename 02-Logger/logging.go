package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

func (mw loggingMiddleware) FetchUser(ctx context.Context, id int) (user *User) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "FetchUser",
			"input:id", id,
			"output:id", user.ID,
			"output:name", user.Name,
			"took", time.Since(begin),
		)
	}(time.Now())

	user = mw.next.FetchUser(ctx, id)
	return
}
