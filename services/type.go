package Services

import (
	"github.com/abdlmunemdhaw/customs/model"
	"gorm.io/gorm"
)

type TypeService struct {
	Db *gorm.DB;
}

func NewTypeService(db *gorm.DB) *TypeService {
	return &TypeService{Db: db};
}

func (s *TypeService) GetMaxId()int{
	var maxId int;
	s.Db.Raw("SELECT MAX(type_id) FROM type").Scan(&maxId);
	return maxId;
}

func (s *TypeService) CreateType(data *model.Type) error {
	return s.Db.Create(data).Error;
}

func (s *TypeService) GetAllTypes()([]model.Type ,error) {
	var types []model.Type;
	if err := s.Db.Find(&types).Error;err != nil{
		return nil,err;
	}
	return types,nil;
}

func (s *TypeService) UpdateType(type_id string,data map[string] interface{}) error {
	if err := s.Db.Model(&model.Type{}).Where("type_id= ?",type_id).Updates(data).Error ; err != nil{
		return err;
	}
	return nil;
}

func (s *TypeService) DeleteType(type_id string) error{
	if err := s.Db.Model(&model.Type{}).Delete("type_id = ?",type_id).Error; err != nil {
		return err;
	}
	return nil;
}