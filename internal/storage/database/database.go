package database

import (
	"context"

	"github.com/leguminosa/kestrel/internal/entity"
)

type CharacterInterface interface {
	Find(ctx context.Context, filter *entity.CharacterFilter) (entity.CharacterResult, error)
	FindByID(ctx context.Context, id int) (*entity.Character, error)
	Save(ctx context.Context, model *entity.Character) (int, error)
	Update(ctx context.Context, model *entity.Character) error
	Delete(ctx context.Context, id int) error
}
