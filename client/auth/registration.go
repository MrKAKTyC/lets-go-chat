package auth

import (
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	userName string
	password string
}

type CreateUserResponse struct {
	Id       uuid.UUID `json:"id"`
	Password string    `json:"password"`
}

func NewUserResponse(userName uuid.UUID, password string) *CreateUserResponse {
	return &CreateUserResponse{userName, password}
}

func (user *CreateUserResponse) GetUserName() string {
	return user.Password
}

func (user *CreateUserResponse) GetId() uuid.UUID {
	return user.Id
}

func NewUserRequest(userName, password string) *CreateUserRequest {
	return &CreateUserRequest{userName, password}
}

func (user *CreateUserRequest) GetUserName() string {
	return user.userName
}

func (user *CreateUserRequest) GetPassword() string {
	return user.password
}
