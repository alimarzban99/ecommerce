package client

type OrderResource struct {
	ID           uint    `json:"id"`
	TrackingCode string  `json:"tracking_code"`
	Amount       float64 `json:"amount"`
	CreatedAt    string  `json:"created_at"`
}
