package model;

type Field struct {
	FieldId int `gorm:"primaryKey"  json:"field_id"`
	Name string `json:"name"`
	Details string `json:"details"`
}

func (Field) TableName () string{
	return "field";
}