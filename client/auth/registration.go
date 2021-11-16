package auth

// CreateUserRequest defines model for CreateUserRequest.
type CreateUserRequest struct {
	Password string `json:"password"`
	UserName string `json:"userName"`
}

// CreateUserResponse defines model for CreateUserResponse.
type CreateUserResponse struct {
	Id       *string `json:"id,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody CreateUserRequest

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody
