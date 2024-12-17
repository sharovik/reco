package services

import "fmt"

type BaseUsersServiceInterface interface {
	Store(user User) error
}

type User struct {
	ID   string
	Name string
}

type UsersService struct {
}

func (s UsersService) Store(user User) error {
	//@todo: do something with
	fmt.Println(user)

	return nil
}
