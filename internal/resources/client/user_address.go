package client

type UserAddressResource struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	City      string  `json:"city"`
	Address   string  `json:"address"`
	Status    string  `json:"status"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	CreatedAt string  `json:"created_at"`
}
type UserAddressListResource struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
}
