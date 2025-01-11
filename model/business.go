package model;

type Business struct {
	BusinessId int `gorm:"primaryKey" json:"business_id"`
	Name string `json:"name"`
}

func (Business) TableName () string {
	return "business";
}