package handler

import (
	"net/http"

	"github.com/leguminosa/kestrel/internal/app/http/wrapper"
	"github.com/leguminosa/kestrel/internal/module/gear"
)

type (
	GearHandler struct {
		module gear.GearModule
	}
)

func NewGearHandler(module gear.GearModule) *GearHandler {
	return &GearHandler{
		module: module,
	}
}

func (h *GearHandler) GetGearSetOption(w http.ResponseWriter, r *http.Request) {
	result, err := h.module.GetGearSetOption()
	if err != nil {
		wrapper.InternalServerError(w, err, nil)
		return
	}

	wrapper.OK(w, result)
}
