package dto

import "github.com/axidex/api-example/server/pkg/tables"

type User struct {
	Name string `json:"name" form:"name"`
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

func (u *User) Storage() *tables.User {
	return &tables.User{
		Name: u.Name,
	}
}

func UserFromStorage(user *tables.User) *User {
	return &User{
		Name: user.Name,
	}
}
