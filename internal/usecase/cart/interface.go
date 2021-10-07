package cart

import "github.com/julioshinoda/cart/entity"

//Repository interface
type Repository interface {
	Get(id string) (entity.Cart, error)
	Update(e entity.Cart) error
}

//UseCase interface
type UseCase interface {
	GetCart(id string) (entity.Cart, error)
}
