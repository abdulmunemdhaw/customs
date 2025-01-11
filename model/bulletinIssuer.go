package model;

type BulletinIssuer struct {
	BulletinIssuerId int `gorm:"primaryKey" json:"bulletin_issuer_id"`
	Name string `json:"name"`
}

func (BulletinIssuer) TableName() string {
	return "bulletin_issuer";
}