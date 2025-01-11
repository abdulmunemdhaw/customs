package Services

import (
	"github.com/abdlmunemdhaw/customs/model"
	"gorm.io/gorm"
)

type FieldService struct {
	Db *gorm.DB
}

func NewFieldService(db *gorm.DB) *FieldService{
	return &FieldService{Db: db}
}
func (s *FieldService) CreateSector (data *model.Field) error {
	return s.Db.Create(data).Error;
}
func (s *FieldService) GetMaxId () int {
	var maxId int;
	s.Db.Raw("SELECT MAX(field_id) FROM field").Scan(&maxId);
	return maxId;
}
func (s *FieldService) GetAllFields() ([]model.Field,error) {
	var fields []model.Field;
	if err := s.Db.Find(&fields).Error; err != nil {
		return nil,err;
	}
	return fields,nil;
}
func (s *FieldService) UpdateField(field_id string, data map[string] interface{} ) error {
	return s.Db.Model(&model.Field{}).Where("field_id = ?",field_id).Updates(data).Error
}
func (s *FieldService) DeleteField(field_id string) error{
	return s.Db.Model(&model.Field{}).Delete("field_id = ?",field_id).Error;
}