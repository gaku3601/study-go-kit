package main

import "context"

// userデータ
var (
	users = map[int]User{
		1: User{ID: 1, Name: "user1"},
		2: User{ID: 2, Name: "user2"},
		3: User{ID: 3, Name: "user3"},
	}
)

type UserService interface {
	FetchUser(ctx context.Context, id int) *User
}

type userService struct{}

func (userService) FetchUser(ctx context.Context, id int) *User {
	user, _ := users[id]
	return &user
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
