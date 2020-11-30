package model

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	UserTypeID int    `json:"user_type_id"`
}

func (u *User) NewUser(username, password string, userTypeID int) *User {
	return &User{
		Username:   username,
		Password:   password,
		UserTypeID: userTypeID,
	}
}
