package model;

type Sector struct {
	SectorId int `gorm:"primaryKey" json:"sector_id"`
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
}

func (Sector) TableName() string {
	return "sector";
}