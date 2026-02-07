package user

import (
	"ecomerce/domain"
	userHandler "ecomerce/rest/handlers/user"
)


type Service interface {
	userHandler.Service // embedding
}

type UserRepo interface {

	Creat(user domain.User) (*domain.User,error)

	Find(email,pass string) (*domain.User,error)

	// List() ([]*User,error)

	// Delete(userId int) error

	// Update(user User) (*User,error)
}