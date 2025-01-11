package Services

import (
	"github.com/abdlmunemdhaw/customs/model"
	"gorm.io/gorm"
)

type DocumentService struct {
	Db *gorm.DB
}

func NewDocumentService(db *gorm.DB) *DocumentService {
	return &DocumentService{Db: db}
}

func (s *DocumentService) CreateDocument(data *model.Document) error {
	return s.Db.Create(data).Error
}
func (s *DocumentService) GetMaxId() int {
	var maxId int
	s.Db.Raw("SELECT MAX(document_id) FROM documents").Scan(&maxId)
	return maxId
}
func (s *DocumentService) GetAllDocuments() ([]model.Document, error) {
	var documents []model.Document
	if err := s.Db.Find(&documents).Error; err != nil {
		return nil, err
	}
	return documents, nil
}
func (s *DocumentService) UpdateDocument(document_id string, data map[string]interface{}) error {
	return s.Db.Model(&model.Document{}).Where("document_id = ?", document_id).Updates(data).Error
}
func (s *DocumentService) DeleteDocument(business_id string) error {
	return s.Db.Model(&model.Document{}).Delete("document_id = ?", business_id).Error
}
