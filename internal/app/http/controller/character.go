package controller

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/leguminosa/kestrel/internal/app/http/request"
	"github.com/leguminosa/kestrel/internal/app/http/response"
	"github.com/leguminosa/kestrel/internal/app/http/utilities"
	"github.com/leguminosa/kestrel/internal/entity"
	"github.com/leguminosa/kestrel/internal/module"
	"github.com/leguminosa/kestrel/pkg/httpx"
)

type Character interface {
	RegisterRoutes(router *httpx.Router)
}

func NewCharacter(
	validate *validator.Validate,
	module module.CharacterInterface,
) Character {
	return &character{
		validate: validate,
		module:   module,
	}
}

type character struct {
	validate *validator.Validate
	module   module.CharacterInterface
}

func (c *character) RegisterRoutes(router *httpx.Router) {
	router.Get("/characters", utilities.GetHandler(c.validate, c.getCharacters))
	router.Get("/characters/{id}", utilities.GetHandler(c.validate, c.getCharacter))
	router.Post("/characters", utilities.PostHandler(c.validate, c.createCharacter))
	router.Put("/characters/{id}", utilities.PutHandler(c.validate, c.updateCharacter))
	router.Delete("/characters/{id}", utilities.GetHandler(c.validate, c.deleteCharacter))
}

func (c *character) getCharacters(
	ctx context.Context,
	req *request.GetCharacters,
) (*response.GetCharacters, error) {
	result, err := c.module.GetCharacters(ctx, &entity.CharacterFilter{
		Rarity:  req.Rarity,
		Faction: req.Faction,
		Name:    req.Name,
		Cost:    req.Cost,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}

	return toResponseGetCharacters(result), nil
}

func (c *character) getCharacter(
	ctx context.Context,
	req *request.GetCharacter,
) (*response.GetCharacter, error) {
	model, err := c.module.GetCharacterByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &response.GetCharacter{
		Character: toResponseCharacter(model),
	}, nil
}

func (c *character) createCharacter(
	ctx context.Context,
	req *request.CreateCharacter,
) (*response.CreateCharacter, error) {
	id, err := c.module.CreateCharacter(ctx, &entity.Character{
		Rarity:  req.Rarity,
		Faction: req.Faction,
		Name:    req.Name,
		Cost:    req.Cost,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &response.CreateCharacter{
		ID: id,
	}, nil
}

func (c *character) updateCharacter(
	ctx context.Context,
	req *request.UpdateCharacter,
) (*response.UpdateCharacter, error) {
	err := c.module.UpdateCharacter(ctx, &entity.Character{
		ID:      req.ID,
		Rarity:  req.Rarity,
		Faction: req.Faction,
		Name:    req.Name,
		Cost:    req.Cost,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &response.UpdateCharacter{
		ID: req.ID,
	}, nil
}

func (c *character) deleteCharacter(
	ctx context.Context,
	req *request.DeleteCharacter,
) (*response.DeleteCharacter, error) {
	err := c.module.DeleteCharacter(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &response.DeleteCharacter{
		ID: req.ID,
	}, nil
}

func toResponseGetCharacters(result entity.CharacterResult) *response.GetCharacters {
	resp := &response.GetCharacters{
		Characters: []*response.Character{},
	}
	for _, model := range result.List {
		resp.Characters = append(resp.Characters, toResponseCharacter(model))
	}
	return resp
}

func toResponseCharacter(model *entity.Character) *response.Character {
	return &response.Character{
		ID:        model.ID,
		Rarity:    model.Rarity,
		Faction:   model.Faction,
		Name:      model.Name,
		Cost:      model.Cost,
		Status:    model.Status,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
