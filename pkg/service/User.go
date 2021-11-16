package service

import (
	"errors"

	"github.com/MrKAKTyC/lets-go-chat/pkg/dao"
	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/hasher"
	"github.com/google/uuid"
)

type User struct {
	storage map[string]dao.User
}

func New() User {
	return User{make(map[string]dao.User)}
}

func (u *User) Register(user auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	if len(user.Password) < 8 || len(user.UserName) < 4 {
		return nil, errors.New("bad request, empty username or id")
	}
	userPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	_, userExist := u.storage[user.UserName]
	if userExist {
		return nil, errors.New("User already exists")
	}
	userUUID := uuid.New().String()
	u.storage[user.UserName] = dao.User{ID: userUUID, Login: user.UserName, Password: userPassword}

	return &auth.CreateUserResponse{Id: &userUUID, UserName: &user.UserName}, nil
}

func (userService *User) Authorize(user auth.LoginUserRequest) (*auth.LoginUserResonse, error) {
	userInDB, ok := userService.storage[user.UserName]
	userResponse := new(auth.LoginUserResonse)
	if ok && hasher.CheckPasswordHash(user.Password, userInDB.Password) {
		userResponse.Url = "url"
		return userResponse, nil
	}
	return userResponse, errors.New("no user with such credentials")
}
