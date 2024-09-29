package response

import (
	"time"

	"github.com/leguminosa/kestrel/internal/enum"
)

type (
	GetCharacters struct {
		Characters []*Character `json:"characters"`
	}
	GetCharacter struct {
		*Character
	}
	Character struct {
		ID        int                  `json:"id"`
		Rarity    enum.Rarity          `json:"rarity"`
		Faction   enum.Faction         `json:"faction"`
		Name      string               `json:"name"`
		Cost      int                  `json:"cost"`
		Status    enum.CharacterStatus `json:"status"`
		CreatedAt time.Time            `json:"created_at"`
		UpdatedAt *time.Time           `json:"updated_at,omitempty"`
	}
	CreateCharacter struct {
		ID int `json:"id"`
	}
	UpdateCharacter struct {
		ID int `json:"id"`
	}
	DeleteCharacter struct {
		ID int `json:"id"`
	}
)
