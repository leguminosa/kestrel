package character

import (
	"context"

	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/internal/storage/database"
)

type CharacterRepository struct {
	db database.CharacterInterface
}

func NewCharacterRepository(db database.CharacterInterface) *CharacterRepository {
	return &CharacterRepository{
		db: db,
	}
}

func (r *CharacterRepository) GetByID(ctx context.Context, id int) (*entity.Character, error) {
	return r.db.GetByID(ctx, id)
}
