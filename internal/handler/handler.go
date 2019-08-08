package handler

import (
	"vision/internal/interface/config"
	"vision/internal/interface/image"
)

type Handler struct {
	configStore config.Store
	imageStore  image.Store
}

func NewHandler(cs config.Store, is image.Store) *Handler {
	return &Handler{
		configStore: cs,
		imageStore:  is,
	}
}
