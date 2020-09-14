package entities

import "fmt"

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UserID struct {
	Id int64 `json:"id"`
}

type UserName struct {
	Name string `json:"name"`
}

func (user User) ToString() string {
	return fmt.Sprintf("Name : %s", user.Name)
}
