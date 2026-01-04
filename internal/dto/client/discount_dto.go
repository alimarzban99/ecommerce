package client

type ValidateDiscountDTO struct {
	Code    string  `json:"code"`
	OrderId float64 `json:"order_id"`
}
