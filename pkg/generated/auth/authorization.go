package auth

// LoginUserRequest defines model for LoginUserRequest.
type LoginUserRequest struct {
	// The password for login in clear text
	Password string `json:"password"`

	// The user name for login
	UserName string `json:"userName"`
}

// LoginUserResonse defines model for LoginUserResonse.
type LoginUserResonse struct {
	// A url for websoket API with a one-time token for starting chat
	Url string `json:"url"`
}

// LoginUserJSONBody defines parameters for LoginUser.
type LoginUserJSONBody LoginUserRequest

// LoginUserJSONRequestBody defines body for LoginUser for application/json ContentType.
type LoginUserJSONRequestBody LoginUserJSONBody
