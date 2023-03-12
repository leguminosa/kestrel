package persistence

import (
	"context"
	"time"

	"github.com/leguminosa/kestrel/internal/constant"
	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/pkg/database"
	"github.com/leguminosa/kestrel/pkg/gormx"
)

type (
	GearDB interface {
		FindGearSetOption(ctx context.Context, filter *entity.GearSetOptionFilter) (entity.GearSetOptionResult, error)
		InsertGearSetOption(ctx context.Context, model *entity.GearSetOption) error
		UpdateGearSetOption(ctx context.Context, model *entity.GearSetOption) error
		DeleteGearSetOption(ctx context.Context, id int) error
	}
	gearDBImpl struct {
		master, slave gormx.Client
	}
)

func NewGearDatabase(db *database.Database) GearDB {
	return &gearDBImpl{
		master: db.Master,
		slave:  db.Slave,
	}
}

func (db *gearDBImpl) FindGearSetOption(
	ctx context.Context,
	filter *entity.GearSetOptionFilter,
) (entity.GearSetOptionResult, error) {
	var (
		result = entity.GearSetOptionResult{
			List: []entity.GearSetOption{},
		}
		err error
	)

	query := db.slave.Table("gear_set_option").
		Where("status <> ?", constant.GearSetOptionStatusDeleted)

	query = query.Select(`
		id,
		name,
		set_count,
		status,
		create_time,
		update_time
	`)

	if filter.ID > 0 {
		query = query.Where("id = ?", filter.ID)
	}

	if filter.Name != "" {
		query = query.Where("LOWER(name) LIKE ?", filter.NameFilter())
	}

	if filter.SetCount > 0 {
		query = query.Where("set_count = ?", filter.SetCount)
	}

	if filter.UseStatusFilter() {
		query = query.Where("status = ?", filter.Status)
	}

	if filter.Datatable.IsPaginated() {
		query = query.Count(&result.Count)
	}

	err = query.Find(&result.List).Error()
	if err != nil {
		return result, err
	}

	result.Page = filter.Datatable.Pagination.Page
	result.Row = filter.Datatable.Pagination.Limit

	return result, nil
}

func (db *gearDBImpl) InsertGearSetOption(
	ctx context.Context,
	model *entity.GearSetOption,
) error {
	var err error

	query := db.master.Begin()
	defer func() {
		if err != nil {
			query.Rollback()
		}
	}()

	err = query.Table("gear_set_option").Create(map[string]interface{}{
		"name":        model.Name,
		"set_count":   model.SetCount,
		"status":      model.Status,
		"create_time": time.Now(),
	}).Error()
	if err != nil {
		return err
	}

	err = query.Commit().Error()
	if err != nil {
		return err
	}

	return nil
}

func (db *gearDBImpl) UpdateGearSetOption(
	ctx context.Context,
	model *entity.GearSetOption,
) error {
	var err error

	query := db.master.Begin()
	defer func() {
		if err != nil {
			query.Rollback()
		}
	}()

	err = query.Table("gear_set_option").Where("id = ?", model.ID).Updates(map[string]interface{}{
		"name":        model.Name,
		"set_count":   model.SetCount,
		"status":      model.Status,
		"update_time": time.Now(),
	}).Error()
	if err != nil {
		return err
	}

	err = query.Commit().Error()
	if err != nil {
		return err
	}

	return nil
}

func (db *gearDBImpl) DeleteGearSetOption(
	ctx context.Context,
	id int,
) error {
	var err error

	query := db.master.Begin()
	defer func() {
		if err != nil {
			query.Rollback()
		}
	}()

	err = query.Table("gear_set_option").Where("id = ?", id).Updates(map[string]interface{}{
		"status":      constant.GearSetOptionStatusDeleted,
		"update_time": time.Now(),
	}).Error()
	if err != nil {
		return err
	}

	err = query.Commit().Error()
	if err != nil {
		return err
	}

	return nil
}
