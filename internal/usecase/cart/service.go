package cart

import "github.com/julioshinoda/cart/entity"

//Service  interface
type Service struct {
	repo Repository
}

func NewService(r Repository) UseCase {
	return &Service{
		repo: r,
	}
}

//GetUser Get an user
func (s *Service) GetCart(id string) (entity.Cart, error) {
	return s.repo.Get(id)
}
