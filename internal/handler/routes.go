package handler

import (
	"vision/internal/platform"
	middleware "vision/internal/router/middleware"

	"github.com/labstack/echo"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(platform.JWTSecret)
	images := v1.Group("/images")
	images.POST("", h.UploadImage, jwtMiddleware)
	images.GET("/:id", h.GetImageByID, jwtMiddleware)
}
