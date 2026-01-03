package model

type UserAddress struct {
	BaseModel
	Name    string  `gorm:"not null"`
	City    string  `gorm:"not null"`
	Address string  `gorm:"not null"`
	Status  string  `gorm:"default:'active';type:enum('active', 'inactive', 'blocked')"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}
