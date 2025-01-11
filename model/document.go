package model;

type Document struct {
	DocumentId int `json:"document_id"`
	Name string `json:"name"`
	SubjectToVal bool `json:"subject_to_val"`
	IdRestriction string `json:"id_restriction"`
}

func (Document) TableName () string {
	return "documents"
}