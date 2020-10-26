package model

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (u *User) NewUser(name, username, email, phone, password string) *User {
	return &User{
		Name:     name,
		Username: username,
		Email:    email,
		Phone:    phone,
		Password: password,
	}
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ul *UserLogin) NewUserLogin(username, password string) *UserLogin {
	return &UserLogin{
		Username: username,
		Password: password,
	}
}
