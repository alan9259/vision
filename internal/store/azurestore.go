package store

import (
	"vision/internal/model"

	"github.com/jinzhu/gorm"
)

type AzureStore struct {
	db *gorm.DB
}

func NewAzureStore(db *gorm.DB) *AzureStore {
	return &AzureStore{
		db: db,
	}
}

func (us *AzureStore) GetApiKey(e string) (*model.Config, error) {
	var m model.Config
	if err := us.db.Where(&model.Config{Name: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
