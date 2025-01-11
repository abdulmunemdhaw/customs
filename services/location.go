package Services

import (
	"github.com/abdlmunemdhaw/customs/model"
	"gorm.io/gorm"
)

type LocationService struct {
	Db *gorm.DB;
}
func NewLocationService (db *gorm.DB)* LocationService{
	return &LocationService{Db:db}
}

func (s *LocationService) CreateLocation(data *model.Location) error {
	return s.Db.Create(data).Error;
}
func (s *LocationService) GetMaxId () int {
	var maxId int;
	s.Db.Raw("SELECT MAX(location_id) FROM locations").Scan(&maxId);
	return maxId;
}
func (s *LocationService) GetLocations() ([]model.Location,error) {
	var locations []model.Location;
	if err := s.Db.Find(&locations).Error; err != nil {
		return nil,err;
	}
	if err := s.Db.Preload("Municipality").Find(&locations).Error; err != nil {
        return nil, err
    }
	return locations,nil;
}
func (s *LocationService) UpdateLocation(location_id string, data map[string] interface{} ) error {
	return s.Db.Model(&model.Location{}).Where("location_id = ?",location_id).Updates(data).Error
}
func (s *LocationService) DeleteLocation(location_id string) error{
	return s.Db.Model(&model.Location{}).Delete("location_id = ?",location_id).Error;
}