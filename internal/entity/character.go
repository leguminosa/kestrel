package entity

type (
	Character struct {
		ID                int
		Name              string
		CharacterRarityID int
		CharacterTypeID   int
	}
	CharacterRarity struct {
		ID   int
		Name string
	}
	CharacterType struct {
		ID   int
		Name string
	}
)
