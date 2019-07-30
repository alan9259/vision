package handler

import (
	config "vision/internal/interfaces"
)

type Handler struct {
	configStore config.ConfigStore
}

func NewHandler(cs config.ConfigStore) *Handler {
	return &Handler{
		configStore: cs,
	}
}
