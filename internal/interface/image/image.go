package image

// image interface
import (
	"vision/internal/model"
)

type Store interface {
	//Create(*model.Image) error
	GetImageByID(uint) (*model.Image, error)
}
