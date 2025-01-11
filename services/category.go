package Services

import (
	"github.com/abdlmunemdhaw/customs/model"
	"gorm.io/gorm"
)

type CategoryService struct {
	Db *gorm.DB;
}
func NewCategoryService (db *gorm.DB)* CategoryService{
	return &CategoryService{Db:db}
}

func (s *CategoryService) CreateCategory(data *model.Category) error {
	return s.Db.Create(data).Error;
}
func (s *CategoryService) GetMaxId () int {
	var maxId int;
	s.Db.Raw("SELECT MAX(category_id) FROM category").Scan(&maxId);
	return maxId;
}
func (s *CategoryService) GetAllCategories() ([]model.Category,error) {
	var Categoryes []model.Category;
	if err := s.Db.Find(&Categoryes).Error; err != nil {
		return nil,err;
	}
	if err := s.Db.Preload("Sector").Preload("Field").Preload("Type").Preload("Business").Find(&Categoryes).Error; err != nil {
        return nil, err
    }
	return Categoryes,nil;
}
func (s *CategoryService) UpdateCategory(category_id string, data map[string] interface{} ) error {
	return s.Db.Model(&model.Category{}).Where("category_id = ?",category_id).Updates(data).Error
}
func (s *CategoryService) DeleteCategory(Category_id string) error{
	return s.Db.Model(&model.Category{}).Delete("category_id = ?",Category_id).Error;
}