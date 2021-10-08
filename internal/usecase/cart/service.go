package cart

import (
	"errors"

	"github.com/julioshinoda/cart/entity"
)

//Service  interface
type Service struct {
	repo Repository
}

func NewService(r Repository) UseCase {
	return &Service{
		repo: r,
	}
}

//GetCart Get an user
func (s *Service) GetCart(ID string) (*entity.Cart, error) {
	return s.repo.Get(ID)
}

//AddItem Add item to cart
func (s *Service) AddItem(ID string, item entity.Items) (*entity.Cart, error) {
	cart, err := s.repo.Get(ID)
	if err != nil && !errors.Is(err, entity.ErrNotFound) {
		return nil, err
	}
	if errors.Is(err, entity.ErrNotFound) {
		cart = &entity.Cart{ID: ID}
	}
	if err := cart.AddItem(item); err != nil {
		return nil, err
	}
	if err := s.repo.Update(cart); err != nil {
		return nil, err
	}

	return cart, nil
}

//RemoveItem Remove item to cart
func (s *Service) RemoveItem(ID, itemID string) error {
	cart, err := s.repo.Get(ID)

	if errors.Is(err, entity.ErrNotFound) {
		return entity.ErrCartNotFound
	}
	if err != nil {
		return err
	}
	if err := cart.RemoveItem(itemID); err != nil {
		return err
	}
	if err := s.repo.Update(cart); err != nil {
		return err
	}
	return nil
}

//ClearCart clear cart item and coupon
func (s *Service) ClearCart(ID string) error {
	cart, err := s.repo.Get(ID)

	if errors.Is(err, entity.ErrNotFound) {
		return entity.ErrCartNotFound
	}
	if err != nil {
		return err
	}
	cart.ClearCart()
	if err := s.repo.Update(cart); err != nil {
		return err
	}
	return nil
}

//UpdateCart update cart item
func (s *Service) UpdateCart(ID string, item entity.Items) error {
	cart, err := s.repo.Get(ID)

	if errors.Is(err, entity.ErrNotFound) {
		return entity.ErrCartNotFound
	}
	if err != nil {
		return err
	}
	if err := cart.AddItemQuantity(item); err != nil {
		return err
	}
	if err := s.repo.Update(cart); err != nil {
		return err
	}
	return nil
}

//AddCoupon add coupon cart
func (s *Service) AddCoupon(ID string, coupon entity.Coupon) error {
	cart, err := s.repo.Get(ID)

	if errors.Is(err, entity.ErrNotFound) {
		return entity.ErrCartNotFound
	}
	if err != nil {
		return err
	}
	cart.AddCoupon(coupon)
	if err := s.repo.Update(cart); err != nil {
		return err
	}
	return nil
}
