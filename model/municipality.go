package model;

type Municipality struct {
	MunicipalityId int `gorm:"primaryKey" json:"municipality_id"`
	Name string `json:"name"`
	RegionId string `json:"region_id"`
	Enabled bool `json:"enabled"`
}

func (Municipality) TableName() string {
    return "municipality";
}