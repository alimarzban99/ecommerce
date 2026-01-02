package client

type ProductCollection struct{}

type ProductResource struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Slug        string            `json:"slug"`
	Description string            `json:"description,omitempty"`
	Image       string            `json:"image"`
	Category    *CategoryResource `json:"category,omitempty"`
	Status      string            `json:"status"`
	CreatedAt   string            `json:"created_at"`
}
