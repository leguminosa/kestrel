package entity

type (
	Character struct {
		ID   int    `db:"character_id"`
		Name string `db:"character_name"`
		Cost int    `db:"character_cost"`

		CharacterRarityID   int    `db:"character_rarity_id"`
		CharacterRarityCode string `db:"character_rarity_code"`

		CharacterFactionID   int    `db:"character_faction_id"`
		CharacterFactionName string `db:"character_faction_name"`
	}
)
