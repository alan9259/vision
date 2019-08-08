package store

import (
	"vision/internal/model"

	"github.com/jinzhu/gorm"
)

//store image in DB
type ImageStore struct {
	db *gorm.DB
}

func NewImageStore(db *gorm.DB) *ImageStore {
	return &ImageStore{
		db: db,
	}
}

func (is *ImageStore) GetApiKey(e string) (*model.Config, error) {
	var m model.Config
	if err := is.db.Where(&model.Config{Name: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (is *ImageStore) GetImageByID(id uint) (*model.Image, error) {
	var i model.Image

	return &i, nil
}
