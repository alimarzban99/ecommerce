package client

type CartAddDTO struct {
	ProductId uint `json:"product_id" binding:"required"`
}

type CartRemoveDTO struct {
	Search   string // Search by product name/title
	Category uint   // Filter by category ID
	Page     int    // Page number (default: 1)
	Limit    int    // Items per page (default: 10, max: 100)
}

type CartFinalizeDTO struct {
	Search   string // Search by product name/title
	Category uint   // Filter by category ID
	Page     int    // Page number (default: 1)
	Limit    int    // Items per page (default: 10, max: 100)
}
