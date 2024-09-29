package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/internal/enum"
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

func (s *CharacterDB) Find(ctx context.Context, filter *entity.CharacterFilter) (entity.CharacterResult, error) {
	var (
		result = entity.CharacterResult{
			List: []*entity.Character{},
		}
		err error
	)

	var db *pgxpool.Pool
	db, err = s.client.GetSlave(ctx)
	if err != nil {
		return result, err
	}

	query := `
		SELECT
			id,
			rarity,
			faction,
			name,
			cost,
			status,
			created_at,
			updated_at
		FROM
			characters
		;
	`

	var rows pgx.Rows
	rows, err = db.Query(ctx, query)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var model entity.Character
		err = rows.Scan(
			&model.ID,
			&model.Rarity,
			&model.Faction,
			&model.Name,
			&model.Cost,
			&model.Status,
			&model.CreatedAt,
			&model.UpdatedAt,
		)
		if err != nil {
			return result, err
		}

		result.List = append(result.List, &model)
	}

	err = rows.Err()
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *CharacterDB) FindByID(ctx context.Context, id int) (*entity.Character, error) {
	db, err := s.client.GetSlave(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT
			id,
			rarity,
			faction,
			name,
			cost,
			status,
			created_at,
			updated_at
		FROM
			characters
		WHERE
			id = $1
		;
	`

	var model entity.Character
	err = db.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.Rarity,
		&model.Faction,
		&model.Name,
		&model.Cost,
		&model.Status,
		&model.CreatedAt,
		&model.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (s *CharacterDB) Save(ctx context.Context, model *entity.Character) (int, error) {
	db, err := s.client.GetMaster(ctx)
	if err != nil {
		return 0, err
	}

	query := `
		INSERT INTO characters (
			rarity,
			faction,
			name,
			cost,
			status
		) VALUES (
		 	$1,
			$2,
			$3,
			$4,
			$5
		) RETURNING id;
	`

	var id int
	err = db.QueryRow(
		ctx,
		query,
		model.Rarity,
		model.Faction,
		model.Name,
		model.Cost,
		model.Status,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *CharacterDB) Update(ctx context.Context, model *entity.Character) error {
	db, err := s.client.GetMaster(ctx)
	if err != nil {
		return err
	}

	query := `
		UPDATE characters
		SET
			rarity = $2,
			faction = $3,
			name = $4,
			cost = $5,
			status = $6,
			updated_at = NOW()
		WHERE
			id = $1
		;
	`

	_, err = db.Exec(
		ctx,
		query,
		model.ID,
		model.Rarity,
		model.Faction,
		model.Name,
		model.Cost,
		model.Status,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *CharacterDB) Delete(ctx context.Context, id int) error {
	db, err := s.client.GetMaster(ctx)
	if err != nil {
		return err
	}

	query := `
		UPDATE characters
		SET
			status = $2,
			updated_at = NOW()
		WHERE
			id = $1
		;
	`

	_, err = db.Exec(
		ctx,
		query,
		id,
		enum.CharacterStatusDeleted,
	)
	if err != nil {
		return err
	}

	return nil
}
