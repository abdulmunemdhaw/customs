package model;

type HsCode struct {
	HsCodeId int `json:"hs_code_id"`
	Name string `json:"name"`
	HsCode string `json:"hscode"`
}

func (HsCode) TableName () string {
	return "hs_codes"
}