package product

import (
	"ecomerce/domain"
	prdctHandlr "ecomerce/rest/handlers/product"
)
type Service interface {
    prdctHandlr.Service
}

type ProductRepo interface {

	Creat(p domain.Products) (*domain.Products,error)

	Get(productId int) (*domain.Products,error)

	List() ([]*domain.Products,error)

	Delete(productId int) error

	Update(p domain.Products) (*domain.Products,error)
}