package request

import "github.com/leguminosa/kestrel/internal/enum"

type (
	GetCharacters struct {
		Rarity  enum.Rarity          `validate:"omitempty,rarity" json:"rarity"`
		Faction enum.Faction         `validate:"omitempty,faction" json:"faction"`
		Name    string               `json:"name"`
		Cost    int                  `json:"cost"`
		Status  enum.CharacterStatus `validate:"omitempty,characterStatus" json:"status"`
	}
	GetCharacter struct {
		ID int `validate:"required" json:"id"`
	}
	CreateCharacter struct {
		Rarity  enum.Rarity          `validate:"required,rarity" json:"rarity"`
		Faction enum.Faction         `validate:"required,faction" json:"faction"`
		Name    string               `validate:"required" json:"name"`
		Cost    int                  `validate:"required" json:"cost"`
		Status  enum.CharacterStatus `validate:"required,characterStatus" json:"status"`
	}
	UpdateCharacter struct {
		ID      int                  `validate:"required" json:"id"`
		Rarity  enum.Rarity          `validate:"required,rarity" json:"rarity"`
		Faction enum.Faction         `validate:"required,faction" json:"faction"`
		Name    string               `validate:"required" json:"name"`
		Cost    int                  `validate:"required" json:"cost"`
		Status  enum.CharacterStatus `validate:"required,characterStatus" json:"status"`
	}
	DeleteCharacter struct {
		ID int `validate:"required" json:"id"`
	}
)
