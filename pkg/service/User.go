package service

import (
	"errors"

	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/hasher"
	"github.com/MrKAKTyC/lets-go-chat/pkg/repository"
)

var url = "url"

type User struct {
	repository repository.User
}

func New(userRepository repository.User) User {
	return User{userRepository}
}

func (u *User) Register(user auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	if len(user.Password) < 8 || len(user.UserName) < 4 {
		return nil, errors.New("bad request, empty username or id")
	}
	userPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = userPassword
	userDao, err := u.repository.Create(user.UserName, user.Password)
	if userDao == nil {
		return nil, err
	}
	return &auth.CreateUserResponse{Id: &userDao.ID, UserName: &user.UserName}, nil

}

func (u *User) Authorize(user auth.LoginUserRequest) (*auth.LoginUserResonse, error) {
	userPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = userPassword
	_, err = u.repository.Get(user.UserName, user.Password)
	if err != nil {
		return nil, err
	}
	return &auth.LoginUserResonse{Url: url}, nil
}
