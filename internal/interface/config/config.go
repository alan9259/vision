package config

import (
	"vision/internal/model"
)

type Store interface {
	GetApiKey(string) (*model.Config, error)
}
