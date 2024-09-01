package database

import (
	"context"

	"github.com/leguminosa/kestrel/internal/entity"
)

type CharacterInterface interface {
	GetByID(ctx context.Context, id int) (*entity.Character, error)
}
