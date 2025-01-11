package model;

type Location struct {
	LocationId int `json:"location_id"`
	Name string `json:"name"`
	Address string `json:"address"`
	PhoneNo string `json:"phone_no"`
	MunicipalityId int `json:"municipality_id"`
	Municipality Municipality `gorm:"foreignKey:MunicipalityId"`
}

func (Location) TableName ()string{
	return "locations";
}