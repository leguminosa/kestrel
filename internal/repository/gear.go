package repository

import (
	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/internal/storage/persistence"
)

type (
	Gear interface {
		GetGearSetOption() ([]entity.GearSetOption, error)
	}
	gearImpl struct {
		db persistence.Gear
	}
)

func NewGearRepository(db persistence.Gear) Gear {
	return &gearImpl{
		db: db,
	}
}

func (repo *gearImpl) GetGearSetOption() ([]entity.GearSetOption, error) {
	// get from cache

	// get from db if cache miss
	return repo.db.GetGearSetOption()
}
