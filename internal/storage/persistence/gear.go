package persistence

import (
	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/pkg/database"
	"github.com/leguminosa/kestrel/pkg/gormx"
)

type (
	Gear interface {
		GetGearSetOption() ([]entity.GearSetOption, error)
	}
	gearImpl struct {
		master, slave gormx.Client
	}
)

func NewGearDatabase(db *database.Database) Gear {
	return &gearImpl{
		master: db.Master,
		slave:  db.Slave,
	}
}

func (db *gearImpl) GetGearSetOption() ([]entity.GearSetOption, error) {
	var (
		result = []entity.GearSetOption{}
		err    error
	)

	err = db.slave.Table("gear_set_option").Select(`id, name, set_count`).Find(&result).Error()
	if err != nil {
		return nil, err
	}

	return result, nil
}
