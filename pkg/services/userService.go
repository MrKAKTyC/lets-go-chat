package services

import (
	"errors"

	"github.com/MrKAKTyC/lets-go-chat/client/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/hasher"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Login    string
	Password string
}

var userDB = make(map[string]User)

func RegisterUser(user auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	if len(user.GetPassword()) < 8 || len(user.GetUserName()) < 4 {
		return nil, errors.New("bad request, empty username or id")
	}
	userPassword, err := hasher.HashPassword(user.GetPassword())
	if err != nil {
		return nil, err
	}
	_, userExist := userDB[user.GetUserName()]
	if userExist {
		return nil, errors.New("User already exists")
	}
	userUUID := uuid.New()
	userDB[user.GetUserName()] = User{userUUID, user.GetUserName(), userPassword}
	return auth.NewUserResponse(userUUID, userPassword), nil
}

func AuthorizeUser(user auth.Authorization) (string, error) {
	userInDB, ok := userDB[user.GetLogin()]
	if ok && hasher.CheckPasswordHash(user.GetPassword(), userInDB.Password) {
		return "url", nil
	}
	return "", errors.New("no user with such credentials")
}
