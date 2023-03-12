package repository

import (
	"context"

	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/internal/storage/persistence"
)

type (
	Gear interface {
		FindGearSetOption(ctx context.Context, filter *entity.GearSetOptionFilter) (entity.GearSetOptionResult, error)
		InsertGearSetOption(ctx context.Context, model *entity.GearSetOption) error
		UpdateGearSetOption(ctx context.Context, model *entity.GearSetOption) error
		DeleteGearSetOption(ctx context.Context, id int) error
	}
	gearImpl struct {
		db persistence.GearDB
	}
)

func NewGearRepository(db persistence.GearDB) Gear {
	return &gearImpl{
		db: db,
	}
}

func (repo *gearImpl) FindGearSetOption(ctx context.Context, filter *entity.GearSetOptionFilter) (entity.GearSetOptionResult, error) {
	// get from cache

	// get from db if cache miss
	return repo.db.FindGearSetOption(ctx, filter)
}

func (repo *gearImpl) InsertGearSetOption(ctx context.Context, model *entity.GearSetOption) error {
	return repo.db.InsertGearSetOption(ctx, model)
}

func (repo *gearImpl) UpdateGearSetOption(ctx context.Context, model *entity.GearSetOption) error {
	return repo.db.UpdateGearSetOption(ctx, model)
}

func (repo *gearImpl) DeleteGearSetOption(ctx context.Context, id int) error {
	return repo.db.DeleteGearSetOption(ctx, id)
}
