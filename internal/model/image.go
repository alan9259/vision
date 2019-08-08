package model

import "time"

type Image struct {
	ID            int32
	BodyPath      string
	ThumbnailPath string
	CreatedAt     time.Time
}
