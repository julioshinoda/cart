package cart

import "github.com/julioshinoda/cart/entity"

//Repository interface
type Repository interface {
	Get(id string) (*entity.Cart, error)
	Update(e *entity.Cart) error
}

//UseCase interface
type UseCase interface {
	GetCart(id string) (*entity.Cart, error)
	AddItem(id string, item entity.Items) (*entity.Cart, error)
	RemoveItem(ID, itemID string) error
	ClearCart(ID string) error
	UpdateCart(ID string, item entity.Items) error
	AddCoupon(ID string, coupon entity.Coupon) error
}
