package services

import (
	"errors"

	"github.com/MrKAKTyC/lets-go-chat/client/auth"
	"github.com/MrKAKTyC/lets-go-chat/pkg/hasher"
	"github.com/MrKAKTyC/lets-go-chat/pkg/repository"
)

var url = "url"

type User struct {
	ID       string
	Login    string
	Password string
}

type UserService struct {
	userRepo repository.UserRepository
}

func New(userRepository repository.UserRepository) UserService {
	return UserService{userRepository}
}

func (userService *UserService) RegisterUser(user auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	if len(user.Password) < 8 || len(user.UserName) < 4 {
		return nil, errors.New("bad request, empty username or id")
	}
	userPassword, err := hasher.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = userPassword
	userDao := userService.userRepo.CreateUser(user.UserName, user.Password)

	return &auth.CreateUserResponse{Id: &userDao.ID, UserName: &user.Password}, nil

}

func (userService *UserService) AuthorizeUser(user auth.LoginUserRequest) (*auth.LoginUserResonse, error) {
	userService.userRepo.GetUser(user.UserName, user.Password)

	// userInDB, ok := userService.storage[user.UserName]
	// userResponse := new(auth.LoginUserResonse)
	// if ok && hasher.CheckPasswordHash(user.Password, userInDB.Password) {
	// 	userResponse.Url = "url"
	// 	return userResponse, nil
	// }
	return &auth.LoginUserResonse{Url: url}, errors.New("no user with such credentials")
}
