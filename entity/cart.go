package entity

type Cart struct {
	ID          string  `json:"id"`
	Total       int     `json:"total"`
	Subtotal    int     `json:"subtotal"`
	CouponCode  string  `json:"coupon-code"`
	CouponValue int     `json:"coupon-value"`
	CouponType  string  `json:"coupon-type"`
	Items       []Items `json:"items"`
}
