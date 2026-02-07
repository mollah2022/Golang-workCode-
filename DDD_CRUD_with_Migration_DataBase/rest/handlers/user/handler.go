package user

import (
	"ecomerce/config"
)



type Handler struct {
	cnf *config.Config
    svc Service
}


func NewHandler(cnf *config.Config,  svc Service) *Handler{
	return &Handler{
		svc: svc,
		cnf: cnf,
	}
}