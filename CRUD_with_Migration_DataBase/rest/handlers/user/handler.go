package user

import (
	"ecomerce/config"
	"ecomerce/repo"
)



type Handler struct {
	cnf *config.Config
     userRepo repo.UserRepo
}


func NewHandler(cnf *config.Config, userRepo repo.UserRepo) *Handler{
	return &Handler{
		userRepo: userRepo,
		cnf: cnf,
	}
}