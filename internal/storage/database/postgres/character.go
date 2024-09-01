package postgres

import (
	"context"

	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/internal/tools"
)

type CharacterDB struct {
	client tools.PgxClient
}

func NewCharacterDB(client tools.PgxClient) *CharacterDB {
	return &CharacterDB{
		client: client,
	}
}

func (s *CharacterDB) GetByID(ctx context.Context, id int) (*entity.Character, error) {
	db, err := s.client.GetSlave(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT
			c.id AS "character_id",
			c.name AS "character_name",
			c.cost AS "character_cost",
			c.character_rarity_id,
			cr.code AS "character_rarity_code",
			c.character_faction_id,
			cf.name AS "character_faction_name"
		FROM
			characters c
			LEFT JOIN character_rarities cr ON cr.id = c.character_rarity_id
			LEFT JOIN character_factions cf ON cf.id = c.character_faction_id
		WHERE
			c.id = $1
	`

	var result entity.Character
	err = db.QueryRow(ctx, query, id).Scan(
		&result.ID,
		&result.Name,
		&result.Cost,
		&result.CharacterRarityID,
		&result.CharacterRarityCode,
		&result.CharacterFactionID,
		&result.CharacterFactionName,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
