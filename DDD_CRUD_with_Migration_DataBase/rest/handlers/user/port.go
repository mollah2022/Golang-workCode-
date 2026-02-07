package user

import "ecomerce/domain"

type Service interface {
	Creat(user domain.User) (*domain.User, error)
	Find(email string, pass string) (*domain.User, error)
}