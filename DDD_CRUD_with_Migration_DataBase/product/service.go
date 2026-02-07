package product

import "ecomerce/domain"


type service struct {
	prdctRepo ProductRepo
}

func NewService(prdctRepo ProductRepo) Service {
	return &service{
		prdctRepo: prdctRepo,
	}
}

func (svc *service) Creat(p domain.Products) (*domain.Products,error){
    return svc.prdctRepo.Creat(p)
}

func (svc *service) Get(id int) (*domain.Products,error){
    return svc.prdctRepo.Get(id)
}

func (svc *service) List() ([]*domain.Products,error){
   return svc.prdctRepo.List()
}

func (svc *service) Delete(productId int) error{
  return svc.prdctRepo.Delete(productId)
}

func (svc *service) Update(p domain.Products) (*domain.Products,error){
   return svc.prdctRepo.Update(p)
}