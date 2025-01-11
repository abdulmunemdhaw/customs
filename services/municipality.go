package Services

import (


	"github.com/abdlmunemdhaw/customs/model"
	"gorm.io/gorm"
)

type MunicipalityService struct {
	Db *gorm.DB;
}

func NewMunicipalityService(db *gorm.DB) *MunicipalityService {
	return &MunicipalityService{Db: db};
}
func (o *MunicipalityService) GetAllMunicipality()([]model.Municipality,error){
	var municipalities []model.Municipality;
	if err := o.Db.Find(&municipalities).Error;err != nil {
		return nil,err;
	}
	return municipalities,nil;
}
func (o *MunicipalityService) GetMunicipality(municipality_id string,data *model.Municipality)(*model.Municipality,error){
	var municipality model.Municipality;
	if err := o.Db.Model(&municipality).Where("municipality_id=?",municipality_id).Update("name",data.Name).Error;err != nil {
		return nil,err;
	}
	return &municipality,nil;
}
func (o *MunicipalityService) CreateMunicipality(data *model.Municipality) error {
	return o.Db.Create(data).Error;
}

func (o *MunicipalityService) UpdateMunicipality(municipality_id string,updates map[string] interface {}) error {
	if err := o.Db.Model(&model.Municipality{}).Where("municipality_id = ?", municipality_id).Updates(updates).Error; err != nil {
        return err
    }
	return nil;
}
func (o *MunicipalityService) GetMaxId() int {
	var maxId int;
	o.Db.Raw("SELECT MAX(municipality_id) FROM municipality").Scan(&maxId);
	return maxId;
}
func (o *MunicipalityService) DeleteMunicipality(municipality_id string,data *model.Municipality) error {
	if err := o.Db.Delete(data,"municipality_id = ?",municipality_id).Error; err != nil {
		return err;
	}
	return nil;
}
