package utilities

import (
	"github.com/go-playground/validator/v10"
	"github.com/leguminosa/kestrel/internal/enum"
)

func RegisterValidator() (*validator.Validate, error) {
	validate := validator.New()

	err := validate.RegisterValidation("rarity", validateRarity)
	if err != nil {
		return nil, err
	}

	err = validate.RegisterValidation("faction", validateFaction)
	if err != nil {
		return nil, err
	}

	err = validate.RegisterValidation("characterStatus", validateCharacterStatus)
	if err != nil {
		return nil, err
	}

	return validate, nil
}

func validateRarity(fl validator.FieldLevel) bool {
	return enum.IsValid(enum.Rarity(fl.Field().String()))
}

func validateFaction(fl validator.FieldLevel) bool {
	return enum.IsValid(enum.Faction(fl.Field().String()))
}

func validateCharacterStatus(fl validator.FieldLevel) bool {
	return enum.IsValid(enum.CharacterStatus(fl.Field().String()))
}
