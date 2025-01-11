package model;

type Category struct {
	CategoryId int `gorm:"primaryKey" json:"category_id"`
	SectorId int `json:"sector_id"`
	FieldId int `json:"field_id"`
	TypeId int `json:"type_id"`
	BusinessId int `json:"business_id"`
	Sector Sector `gorm:"foreignKey:SectorId"`
	Field Field `gorm:"foreignKey:FieldId"`
	Type Type `gorm:"foreignKey:TypeId"`
	Business Business `gorm:"foreignKey:BusinessId"`
}

func (Category) TableName() string {
	return "category"
}