package entity

import (
	"math"
	"strings"
)

type Cart struct {
	ID          string  `json:"id"`
	Total       int     `json:"total"`
	Subtotal    int     `json:"subtotal"`
	CouponCode  string  `json:"coupon-code"`
	CouponValue int     `json:"coupon-value"`
	CouponType  string  `json:"coupon-type"`
	Items       []Items `json:"items"`
}

func (c *Cart) AddItem(item Items) error {
	for _, i := range c.Items {
		if i.ID == item.ID {
			return ErrDuplicatedItem
		}
	}
	c.Items = append(c.Items, item)
	c.Subtotal += (item.Price * item.Quantity)
	c.CalculateTotal()
	return nil
}

func (c *Cart) AddItemQuantity(item Items) error {
	for index, it := range c.Items {
		if it.ID == item.ID {
			c.Subtotal -= (c.Items[index].Price * c.Items[index].Quantity)
			c.Items[index].Quantity = item.Quantity
			c.Subtotal += (c.Items[index].Price * item.Quantity)
			c.CalculateTotal()
			return nil
		}
	}
	return ErrItemNotFound
}

func (c *Cart) RemoveItem(ID string) error {
	for index, it := range c.Items {
		if it.ID == ID {
			c.Items[index] = c.Items[len(c.Items)-1]
			c.Items = c.Items[:len(c.Items)-1]
			c.Subtotal -= (it.Price * it.Quantity)
			c.CalculateTotal()
			return nil
		}
	}
	return ErrItemNotFound
}

func (c *Cart) AddCoupon(coupon Coupon) {
	c.CouponCode = coupon.Code
	c.CouponType = coupon.Type
	c.CouponValue = coupon.Value
	c.CalculateTotal()
}

func (c *Cart) ClearCart() {
	c.Items = []Items{}
	c.CouponCode = ""
	c.CouponType = ""
	c.CouponValue = 0
	c.Subtotal = 0
	c.Total = 0
}

func (c *Cart) CalculateTotal() {
	if c.Subtotal < c.CouponDiscont() {
		c.Total = 0
		return
	}
	c.Total = c.Subtotal - c.CouponDiscont()
}

func (c *Cart) CouponDiscont() int {
	if strings.ToLower(c.CouponType) == "value" {
		return c.CouponValue
	}
	if strings.ToLower(c.CouponType) == "percentage" {
		discount := (float64(c.Subtotal) / 100) * (float64(c.CouponValue) / 100)
		return int(math.Ceil(discount * 100))
	}
	return 0
}
