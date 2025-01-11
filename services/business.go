package Services

import (
	"github.com/abdlmunemdhaw/customs/model"
	"gorm.io/gorm"
)

type BusinessService struct {
	Db *gorm.DB;
}
func NewBusinessService (db *gorm.DB)* BusinessService{
	return &BusinessService{Db:db}
}

func (s *BusinessService) CreateBusiness(data *model.Business) error {
	return s.Db.Create(data).Error;
}
func (s *BusinessService) GetMaxId () int {
	var maxId int;
	s.Db.Raw("SELECT MAX(business_id) FROM business").Scan(&maxId);
	return maxId;
}
func (s *BusinessService) GetAllBusinesses() ([]model.Business,error) {
	var businesses []model.Business;
	if err := s.Db.Find(&businesses).Error; err != nil {
		return nil,err;
	}
	return businesses,nil;
}
func (s *BusinessService) UpdateBusiness(field_id string, data map[string] interface{} ) error {
	return s.Db.Model(&model.Business{}).Where("business_id = ?",field_id).Updates(data).Error
}
func (s *BusinessService) DeleteBusines(business_id string) error{
	return s.Db.Model(&model.Business{}).Delete("business_id = ?",business_id).Error;
}