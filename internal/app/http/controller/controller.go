package controller

import (
	"github.com/leguminosa/kestrel/internal/app/http/handler"
	"github.com/leguminosa/kestrel/internal/module/gear"
	"github.com/leguminosa/kestrel/internal/repository"
	"github.com/leguminosa/kestrel/internal/storage/persistence"
	"github.com/leguminosa/kestrel/pkg/database"
)

type (
	Controller struct {
		*handler.GearHandler
	}
)

func New(db *database.Database) *Controller {
	// storage, db
	gearDB := persistence.NewGearDatabase(db)

	// repo
	gearRepo := repository.NewGearRepository(gearDB)

	// module
	gearModule := gear.NewGearModule(gearRepo)

	// handler
	ctrl := &Controller{}
	ctrl.GearHandler = handler.NewGearHandler(gearModule)

	return ctrl
}
