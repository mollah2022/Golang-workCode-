package user

import "ecomerce/domain"


type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Creat(user domain.User) (*domain.User, error){
     usr, err := svc.userRepo.Creat(user)
	 if usr!=nil {
		return nil,err
	 }
	 return usr,nil
}
func (svc *service) Find(email string, pass string) (*domain.User, error){
     usr,err := svc.userRepo.Find(email,pass)
	  if usr!=nil {
		return nil,err
	 }
	 return usr,nil
}