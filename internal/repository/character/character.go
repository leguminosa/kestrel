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

func (r *CharacterRepository) Find(ctx context.Context, filter *entity.CharacterFilter) (entity.CharacterResult, error) {
	return r.db.Find(ctx, filter)
}

func (r *CharacterRepository) FindByID(ctx context.Context, id int) (*entity.Character, error) {
	return r.db.FindByID(ctx, id)
}

func (r *CharacterRepository) Save(ctx context.Context, model *entity.Character) (int, error) {
	return r.db.Save(ctx, model)
}

func (r *CharacterRepository) Update(ctx context.Context, model *entity.Character) error {
	return r.db.Update(ctx, model)
}

func (r *CharacterRepository) Delete(ctx context.Context, id int) error {
	return r.db.Delete(ctx, id)
}
