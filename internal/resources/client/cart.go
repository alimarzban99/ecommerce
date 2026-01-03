package client

type CartItemResource struct {
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	ProductSlug string `json:"product_slug"`
	Image       string `json:"image"`
	Price       int64  `json:"price"` // Price in smallest currency unit (e.g., cents)
	Quantity    int    `json:"quantity"`
	Subtotal    int64  `json:"subtotal"` // Price * Quantity
}

type CartResource struct {
	Items      []CartItemResource `json:"items"`
	TotalItems int                `json:"total_items"`
	Subtotal   int64              `json:"subtotal"` // Sum of all item subtotals
}
