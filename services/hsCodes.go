package Services

import (
	"github.com/abdlmunemdhaw/customs/model"
	"gorm.io/gorm"
)

type HsCodeService struct {
	Db *gorm.DB
}

func NewHsCodeService(db *gorm.DB) *HsCodeService {
	return &HsCodeService{Db: db}
}
func (s *HsCodeService) CreateHsCode(data *model.HsCode) error {
	return s.Db.Create(data).Error
}
func (s *HsCodeService) GetMaxId() int {
	var maxId int
	s.Db.Raw("SELECT MAX(hs_code_id) FROM hs_codes").Scan(&maxId)
	return maxId
}
func (s *HsCodeService) GetAllHsCodes() ([]model.HsCode, error) {
	var hsCodes []model.HsCode
	if err := s.Db.Find(&hsCodes).Error; err != nil {
		return nil, err
	}
	return hsCodes, nil
}
func (s *HsCodeService) UpdateHsCode(hs_code_id string, data map[string]interface{}) error {
	return s.Db.Model(&model.HsCode{}).Where("hs_code_id = ?",hs_code_id).Updates(data).Error
}
func (s *HsCodeService) DeleteHsCode(hs_code_id string) error {
	return s.Db.Model(&model.HsCode{}).Delete("hs_code_id = ?", hs_code_id).Error
}
