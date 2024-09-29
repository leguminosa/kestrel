package entity

import (
	"time"

	"github.com/leguminosa/kestrel/internal/enum"
)

type (
	Character struct {
		ID        int                  `db:"id"`
		Rarity    enum.Rarity          `db:"rarity"`
		Faction   enum.Faction         `db:"faction"`
		Name      string               `db:"name"`
		Cost      int                  `db:"cost"`
		Status    enum.CharacterStatus `db:"status"`
		CreatedAt time.Time            `db:"created_at"`
		UpdatedAt *time.Time           `db:"updated_at"`
	}
	CharacterFilter struct {
		Rarity  enum.Rarity
		Faction enum.Faction
		Name    string
		Cost    int
		Status  enum.CharacterStatus
	}
	CharacterResult struct {
		List []*Character
	}
)
