package Services

import (
	"github.com/abdlmunemdhaw/customs/model"
	"gorm.io/gorm"
)

type SectorService struct {
	Db *gorm.DB
}

func NewSectorService(db *gorm.DB) *SectorService{
	return &SectorService{Db: db};
}

func (s *SectorService) CreateSector (data *model.Sector) error {
	return s.Db.Create(data).Error;
}
func (s *SectorService) GetMaxId () int {
	var maxId int;
	s.Db.Raw("SELECT MAX(sector_id) FROM sector").Scan(&maxId);
	return maxId;
}
func (s *SectorService) GetAllSector() ([]model.Sector,error) {
	var sectors []model.Sector;
	if err := s.Db.Find(&sectors).Error; err != nil {
		return nil,err;
	}
	return sectors,nil;
}

func (s *SectorService) UpdateSector(sector_id string, data map[string] interface{} ) error {
	return s.Db.Model(&model.Sector{}).Where("sector_id = ?",sector_id).Updates(data).Error
}
func (s *SectorService) DeleteSector(sector_id string) error{
	return s.Db.Model(&model.Sector{}).Delete("sector_id = ?",sector_id).Error;
}