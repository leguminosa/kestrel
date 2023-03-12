package gear

import (
	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/internal/repository"
)

type (
	GearModule interface {
		GetGearSetOption() ([]entity.GearSetOption, error)
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

func (m *moduleImpl) GetGearSetOption() ([]entity.GearSetOption, error) {
	return m.repo.GetGearSetOption()
}
