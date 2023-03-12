package gear

import (
	"context"
	"errors"

	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/internal/repository"
)

type (
	GearModule interface {
		FindGearSetOptions(ctx context.Context, filter *entity.GearSetOptionFilter) (entity.GearSetOptionResult, error)
		FindGearSetOptionByID(ctx context.Context, id int) (entity.GearSetOption, error)
		InsertGearSetOption(ctx context.Context, model *entity.GearSetOption) error
		UpdateGearSetOption(ctx context.Context, model *entity.GearSetOption) error
		DeleteGearSetOption(ctx context.Context, id int) error
	}
	moduleImpl struct {
		repo repository.Gear
	}
)

func NewGearModule(repo repository.Gear) GearModule {
	return &moduleImpl{
		repo: repo,
	}
}

func (m *moduleImpl) FindGearSetOptions(ctx context.Context, filter *entity.GearSetOptionFilter) (entity.GearSetOptionResult, error) {
	return m.repo.FindGearSetOption(ctx, filter)
}

func (m *moduleImpl) FindGearSetOptionByID(ctx context.Context, id int) (entity.GearSetOption, error) {
	result, err := m.repo.FindGearSetOption(ctx, &entity.GearSetOptionFilter{
		Datatable: entity.Datatable{
			Pagination: entity.DatatablePagination{
				Disable: true,
			},
		},
		ID: id,
	})
	if err != nil {
		return entity.GearSetOption{}, err
	}

	if !(len(result.List) > 0) {
		return entity.GearSetOption{}, errors.New("no result")
	}

	return result.List[0], nil
}

func (m *moduleImpl) InsertGearSetOption(ctx context.Context, model *entity.GearSetOption) error {
	return m.repo.InsertGearSetOption(ctx, model)
}

func (m *moduleImpl) UpdateGearSetOption(ctx context.Context, model *entity.GearSetOption) error {
	return m.repo.UpdateGearSetOption(ctx, model)
}

func (m *moduleImpl) DeleteGearSetOption(ctx context.Context, id int) error {
	return m.repo.DeleteGearSetOption(ctx, id)
}
