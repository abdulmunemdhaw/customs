package model;

type EconomicBulletin struct {
	EconomicBulletinId int `gorm:"primaryKey" json:"economic_bulletin_id"`
	IssuanceYear string `json:"issuance_year"`
	DecreeNo string `json:"decree_no"`
	BulletinIssuerId int `json:"bulletin_issuer_Id"`
	BulletinIssuer BulletinIssuer `gorm:"foreignKey:BulletinIssuerId"`
}
func (EconomicBulletin) TableName () string {
	return "economic_bulletin"
}