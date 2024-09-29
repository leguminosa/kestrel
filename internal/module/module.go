package module

import (
	"context"

	"github.com/leguminosa/kestrel/internal/entity"
)

type CharacterInterface interface {
	GetCharacters(ctx context.Context, filter *entity.CharacterFilter) (entity.CharacterResult, error)
	GetCharacterByID(ctx context.Context, id int) (*entity.Character, error)
	CreateCharacter(ctx context.Context, model *entity.Character) (int, error)
	UpdateCharacter(ctx context.Context, model *entity.Character) error
	DeleteCharacter(ctx context.Context, id int) error
}
