package handler

import (
	"github.com/labstack/echo"
)

func (h *Handler) Register(v1 *echo.Group) {
	//jwtMiddleware := middleware.JWT(platform.JWTSecret)
	guestUsers := v1.Group("/connect")
	guestUsers.POST("", h.UploadImage)
}
