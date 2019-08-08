package handler

import (
	"net/http"
	"vision/internal/model"

	"github.com/labstack/echo"
)

func (h *Handler) GetImageByID(c echo.Context) error {
	var i model.Image
	return c.JSON(http.StatusCreated, NewImageResponse(&i))
}
