package model;

type Type struct {
	TypeId int `gorm:"primaryKey" json:"type_id"`
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
}

func (Type) TableName () string {
	return "type"
}