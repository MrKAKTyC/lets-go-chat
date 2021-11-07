package auth

type Authorization struct {
	login    string
	password string
}

func New(login, password string) Authorization {
	return Authorization{login, password}
}

func (auth *Authorization) GetLogin() string {
	return auth.login
}

func (auth *Authorization) GetPassword() string {
	return auth.password
}
