package client

type DiscountValidationResource struct {
	FinalPrice     float64 `json:"final_price"`
	DiscountAmount float64 `json:"discount_amount"`
}
