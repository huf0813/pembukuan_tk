package model

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (al *UserReq) NewUserLogin(username, password string) *UserReq {
	return &UserReq{
		Username: username,
		Password: password,
	}
}
