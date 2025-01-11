package Services

import (
	"github.com/abdlmunemdhaw/customs/model"
	"gorm.io/gorm"
)

type BulletinIssuerService struct {
	Db *gorm.DB;
}
func NewBulletinIssuerService (db *gorm.DB)* BulletinIssuerService{
	return &BulletinIssuerService{Db:db}
}

func (s *BulletinIssuerService) CreateBulletinIssuer(data *model.BulletinIssuer) error {
	return s.Db.Create(data).Error;
}
func (s *BulletinIssuerService) GetMaxId () int {
	var maxId int;
	s.Db.Raw("SELECT MAX(bulletin_issuer_id) FROM bulletin_issuer").Scan(&maxId);
	return maxId;
}
func (s *BulletinIssuerService) GetAllBulletinIssueres() ([]model.BulletinIssuer,error) {
	var BulletinIssueres []model.BulletinIssuer;
	if err := s.Db.Find(&BulletinIssueres).Error; err != nil {
		return nil,err;
	}
	return BulletinIssueres,nil;
}
func (s *BulletinIssuerService) UpdateBulletinIssuer(field_id string, data map[string] interface{} ) error {
	return s.Db.Model(&model.BulletinIssuer{}).Where("bulletin_issuer_id = ?",field_id).Updates(data).Error
}
func (s *BulletinIssuerService) DeleteBulletinIssuer(BulletinIssuer_id string) error{
	return s.Db.Model(&model.BulletinIssuer{}).Delete("bulletin_issuer_id = ?",BulletinIssuer_id).Error;
}