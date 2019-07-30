package interfaces

import (
	"vision/internal/model"
)

type ConfigStore interface {
	GetApiKey(string) (*model.Config, error)
}
