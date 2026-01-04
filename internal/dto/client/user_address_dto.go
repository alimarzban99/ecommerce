package client

type ListUserAddressDTO struct {
	Page  int `form:"page" binding:"required,min=1"`
	Limit int `form:"limit" binding:"required,min=1"`
}

type StoreUserAddressDTO struct {
	Name    string  `json:"name" binding:"required"`
	City    string  `json:"city" binding:"required"`
	Address string  `json:"address" binding:"required"`
	Status  string  `json:"status" binding:"required"`
	Lat     float64 `json:"lat" binding:"required"`
	Lng     float64 `json:"lng" binding:"required"`
}

type UpdateUserAddressDTO struct {
	Name    string  `json:"name" binding:"required"`
	City    string  `json:"city" binding:"required"`
	Address string  `json:"address" binding:"required"`
	Status  string  `json:"status" binding:"required"`
	Lat     float64 `json:"lat" binding:"required"`
	Lng     float64 `json:"lng" binding:"required"`
}
