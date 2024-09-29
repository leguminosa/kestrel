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

func (m *CharacterModule) GetCharacters(ctx context.Context, filter *entity.CharacterFilter) (entity.CharacterResult, error) {
	return m.repo.Find(ctx, filter)
}

func (m *CharacterModule) GetCharacterByID(ctx context.Context, id int) (*entity.Character, error) {
	return m.repo.FindByID(ctx, id)
}

func (m *CharacterModule) CreateCharacter(ctx context.Context, model *entity.Character) (int, error) {
	return m.repo.Save(ctx, model)
}

func (m *CharacterModule) UpdateCharacter(ctx context.Context, model *entity.Character) error {
	return m.repo.Update(ctx, model)
}

func (m *CharacterModule) DeleteCharacter(ctx context.Context, id int) error {
	return m.repo.Delete(ctx, id)
}
