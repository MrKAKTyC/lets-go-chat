package service

import (
	"errors"
	"github.com/MrKAKTyC/lets-go-chat/pkg/repository"

	"github.com/MrKAKTyC/lets-go-chat/pkg/generated/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/hasher"
)

type UserService interface {
	Register(user auth.CreateUserRequest) (*auth.CreateUserResponse, error)
	Authorize(user auth.LoginUserRequest) (*auth.LoginUserResponse, error)
}

type User struct {
	repository *repository.UserRepository
	otpService *OTPService
	url        string
}

func NewUser(userRepository *repository.UserRepository, otpService *OTPService) *UserService {
	var us UserService
	us = &User{
		repository: userRepository,
		otpService: otpService,
		url:        "/chat/ws.rtm.start?token=",
	}
	return &us
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
	userDao, err := (*u.repository).Create(user.UserName, user.Password)
	if err != nil {
		return nil, err
	}
	return &auth.CreateUserResponse{Id: &userDao.ID, UserName: &user.UserName}, nil

}

func (u *User) Authorize(user auth.LoginUserRequest) (*auth.LoginUserResponse, error) {
	userPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = userPassword
	userDAO, err := (*u.repository).Get(user.UserName, user.Password)
	if err != nil {
		return nil, err
	}
	return &auth.LoginUserResponse{Url: u.url + (*u.otpService).GenerateOTP(userDAO.ID)}, nil
}
