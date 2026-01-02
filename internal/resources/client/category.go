package client

type CategoryCollection struct{}

type CategoryResource struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type CategoryPluckResource struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
