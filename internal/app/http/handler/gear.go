package handler

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leguminosa/kestrel/internal/app/http/wrapper"
	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/internal/module/gear"
	"github.com/leguminosa/kestrel/pkg/jsonx"
	"github.com/leguminosa/kestrel/pkg/util/convert"
	"github.com/leguminosa/kestrel/pkg/util/xcontext"
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

func (h *GearHandler) GetGearSetOptions(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = xcontext.GetAllContextFromIncomingRequest(r)
		filter = &entity.GearSetOptionFilter{}
	)

	query := r.URL.Query()

	filter.Datatable.Pagination.Page = convert.ToInt(query.Get("page"))
	filter.Datatable.Pagination.Limit = convert.ToInt(query.Get("row"))
	filter.Datatable.Sort.Field = query.Get("sorted_field")
	filter.Datatable.Sort.Direction = query.Get("sorted_direction")

	filter.Name = query.Get("name")
	filter.SetCount = convert.ToInt(query.Get("set_count"))
	filter.StatusRaw = query.Get("status")
	filter.Status = convert.ToInt(filter.StatusRaw)

	filter.Validate()

	result, err := h.module.FindGearSetOptions(ctx, filter)
	if err != nil {
		wrapper.InternalServerError(w, err, nil)
		return
	}

	wrapper.OK(w, result)
}

func (h *GearHandler) GetGearSetOption(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = xcontext.GetAllContextFromIncomingRequest(r)
		id  int
	)

	routeVars := mux.Vars(r)
	rawID, found := routeVars["id"]
	if !found {
		wrapper.BadRequest(w, errors.New("missing id"), routeVars)
		return
	}

	id = convert.ToInt(rawID)
	if !(id > 0) {
		wrapper.BadRequest(w, errors.New("invalid id"), rawID)
		return
	}

	result, err := h.module.FindGearSetOptionByID(ctx, id)
	if err != nil {
		wrapper.InternalServerError(w, err, nil)
		return
	}

	wrapper.OK(w, result)
}

func (h *GearHandler) CreateGearSetOption(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = xcontext.GetAllContextFromIncomingRequest(r)
		model = &entity.GearSetOption{}
	)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrapper.BadRequest(w, err, body)
		return
	}

	err = jsonx.GetClient().Unmarshal(body, model)
	if err != nil {
		wrapper.BadRequest(w, err, body)
		return
	}

	err = h.module.InsertGearSetOption(ctx, model)
	if err != nil {
		wrapper.InternalServerError(w, err, nil)
		return
	}

	wrapper.OK(w, nil)
}

func (h *GearHandler) UpdateGearSetOption(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = xcontext.GetAllContextFromIncomingRequest(r)
		model = &entity.GearSetOption{}
	)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrapper.BadRequest(w, err, body)
		return
	}

	err = jsonx.GetClient().Unmarshal(body, model)
	if err != nil {
		wrapper.BadRequest(w, err, body)
		return
	}

	routeVars := mux.Vars(r)
	rawID, found := routeVars["id"]
	if !found {
		wrapper.BadRequest(w, errors.New("missing id"), routeVars)
		return
	}

	model.ID = convert.ToInt(rawID)
	if !(model.ID > 0) {
		wrapper.BadRequest(w, errors.New("invalid id"), rawID)
		return
	}

	err = h.module.UpdateGearSetOption(ctx, model)
	if err != nil {
		wrapper.InternalServerError(w, err, nil)
		return
	}

	wrapper.OK(w, nil)
}

func (h *GearHandler) DeleteGearSetOption(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = xcontext.GetAllContextFromIncomingRequest(r)
		id  int
	)

	routeVars := mux.Vars(r)
	rawID, found := routeVars["id"]
	if !found {
		wrapper.BadRequest(w, errors.New("missing id"), routeVars)
		return
	}

	id = convert.ToInt(rawID)
	if !(id > 0) {
		wrapper.BadRequest(w, errors.New("invalid id"), rawID)
		return
	}

	err := h.module.DeleteGearSetOption(ctx, id)
	if err != nil {
		wrapper.InternalServerError(w, err, nil)
		return
	}

	wrapper.OK(w, nil)
}
