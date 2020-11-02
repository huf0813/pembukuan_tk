package model

type AuthLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (al *AuthLogin) NewUserLogin(username, password string) *AuthLogin {
	return &AuthLogin{
		Username: username,
		Password: password,
	}
}
