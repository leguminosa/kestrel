package character

import (
	"context"

	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/internal/repository"
)

type CharacterModule struct {
	repo repository.CharacterInterface
}

func NewCharacterModule(repo repository.CharacterInterface) *CharacterModule {
	return &CharacterModule{
		repo: repo,
	}
}

func (m *CharacterModule) GetByID(ctx context.Context, id int) (*entity.Character, error) {
	return m.repo.GetByID(ctx, id)
}
