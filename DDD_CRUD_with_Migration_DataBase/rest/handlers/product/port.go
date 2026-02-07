package product

import "ecomerce/domain"

type Service interface {
	Creat(domain.Products) (*domain.Products, error)
	Get(id int) (*domain.Products, error)
	List() ([]*domain.Products, error)
	Update(domain.Products) (*domain.Products, error)
	Delete(id int) error
}