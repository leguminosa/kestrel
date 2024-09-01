package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/leguminosa/kestrel/internal/app/http/handler/response"
	"github.com/leguminosa/kestrel/internal/module"
	"github.com/leguminosa/kestrel/pkg/convert"
)

type CharacterHandler struct {
	module module.CharacterInterface
}

func NewCharacterHandler(module module.CharacterInterface) *CharacterHandler {
	return &CharacterHandler{
		module: module,
	}
}

func (h *CharacterHandler) GetByID(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		id  = convert.ToInt(c.Param("id"))
	)

	result, err := h.module.GetByID(ctx, id)
	if err != nil {
		return response.InternalServerError(c, err, "failed to get character by id")
	}

	return response.OK(c, result)
}
