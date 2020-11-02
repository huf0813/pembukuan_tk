package model

type User struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	UserTypeID int    `json:"user_type_id"`
}

func (u *User) NewUser(name, username, email, phone, password string) *User {
	return &User{
		Name:       name,
		Username:   username,
		Email:      email,
		Phone:      phone,
		Password:   password,
		UserTypeID: 2,
	}
}
