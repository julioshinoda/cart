package entity

type Coupon struct {
	Code  string `json:"coupon-code"`
	Value int    `json:"coupon-value"`
	Type  string `json:"coupon-type"`
}
