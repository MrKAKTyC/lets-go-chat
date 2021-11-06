package auth

type Authorization struct {
	login    string
	password string
}

func New(login, password string) Authorization {
	return Authorization{login, password}
}

func (auth *Authorization) Login() string {
	return auth.login
}

func (auth *Authorization) Password() string {
	return auth.password
}
