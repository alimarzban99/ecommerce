package client

type CartAddDTO struct {
	ProductId uint `json:"product_id" binding:"required"`
}

type CartUpdateQuantityDTO struct {
	ProductId uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type CartRemoveDTO struct {
	ProductId uint `json:"product_id" binding:"required"`
}

type CartFinalizeDTO struct {
	PaymentMethod string  `json:"payment_method" binding:"required,oneof=gateway wallet"`
	DiscountCode  *string `json:"discount_code,omitempty"`
}
