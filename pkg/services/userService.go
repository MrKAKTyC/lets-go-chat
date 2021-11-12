package services

import (
	"errors"

	"github.com/MrKAKTyC/lets-go-chat/client/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/hasher"
	"github.com/google/uuid"
)

type User struct {
	ID       string
	Login    string
	Password string
}

type UserService struct {
	storage map[string]User
}

func New() UserService {
	return UserService{make(map[string]User)}
}

func (userService *UserService) RegisterUser(user auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	if len(user.Password) < 8 || len(user.UserName) < 4 {
		return nil, errors.New("bad request, empty username or id")
	}
	userPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	_, userExist := userService.storage[user.UserName]
	if userExist {
		return nil, errors.New("User already exists")
	}
	userUUID := uuid.New().String()
	userService.storage[user.UserName] = User{userUUID, user.UserName, userPassword}

	return &auth.CreateUserResponse{&userUUID, &userPassword}, nil
}

func (userService *UserService) AuthorizeUser(user auth.LoginUserRequest) (*auth.LoginUserResonse, error) {
	userInDB, ok := userService.storage[user.UserName]
	userResponse := new(auth.LoginUserResonse)
	if ok && hasher.CheckPasswordHash(user.Password, userInDB.Password) {
		userResponse.Url = "url"
		return userResponse, nil
	}
	return userResponse, errors.New("no user with such credentials")
}
