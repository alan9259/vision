package handler

import (
	"time"
	"vision/internal/model"
)

type ImageResponse struct {
	ID            int32     `json:"id,omitempty"`
	BodyPath      string    `json:"body_path,omitempty"`
	ThumbnailPath string    `json:"thumbnail_path,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

func NewImageResponse(i *model.Image) *ImageResponse {
	r := new(ImageResponse)
	r.ID = i.ID
	r.BodyPath = i.BodyPath
	r.ThumbnailPath = i.ThumbnailPath
	r.CreatedAt = i.CreatedAt

	return r
}
