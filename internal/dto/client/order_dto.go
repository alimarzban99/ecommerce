package client

type ListOrderDTO struct {
	Page  int // Page number (default: 1)
	Limit int // Items per page (default: 10, max: 100)
}
