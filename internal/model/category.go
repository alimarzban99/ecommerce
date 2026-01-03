package model

type Category struct {
	BaseModel
	Title    string    `gorm:"type:varchar(250);not null"`
	Products []Product `gorm:"foreignKey:CategoryID;constraint:OnDelete:RESTRICT"`
}
