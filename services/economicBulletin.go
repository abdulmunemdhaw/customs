package Services

import (
	"github.com/abdlmunemdhaw/customs/model"
	"gorm.io/gorm"
)

type EconomicBulletinService struct {
	Db *gorm.DB;
}
func NewEconomicBulletinService (db *gorm.DB)* EconomicBulletinService{
	return &EconomicBulletinService{Db:db}
}

func (s *EconomicBulletinService) CreateEconomicBulletin(data *model.EconomicBulletin) error {
	return s.Db.Create(data).Error;
}
func (s *EconomicBulletinService) GetMaxId () int {
	var maxId int;
	s.Db.Raw("SELECT MAX(economic_bulletin_id) FROM economic_bulletin").Scan(&maxId);
	return maxId;
}
func (s *EconomicBulletinService) GetAllEconomicBulletins() ([]model.EconomicBulletin,error) {
	var EconomicBulletins []model.EconomicBulletin;
	if err := s.Db.Find(&EconomicBulletins).Error; err != nil {
		return nil,err;
	}
	if err := s.Db.Preload("BulletinIssuer").Find(&EconomicBulletins).Error; err != nil {
        return nil, err
    }
	return EconomicBulletins,nil;
}
func (s *EconomicBulletinService) UpdateEconomicBulletin(economic_bulletin_id string, data map[string] interface{} ) error {
	return s.Db.Model(&model.EconomicBulletin{}).Where("economic_bulletin_id = ?",economic_bulletin_id).Updates(data).Error
}
func (s *EconomicBulletinService) DeleteEconomicBulletin(economic_bulletin_id string) error{
	return s.Db.Model(&model.BulletinIssuer{}).Delete("economic_bulletin_id = ?",economic_bulletin_id).Error;
}