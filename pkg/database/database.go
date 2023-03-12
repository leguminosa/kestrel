package database

import (
	"github.com/leguminosa/kestrel/pkg/config"
	"github.com/leguminosa/kestrel/pkg/gormx"
)

type (
	Database struct {
		Master gormx.Client
		Slave  gormx.Client
	}
)

func InitDatabase(cfg *config.Config) (*Database, error) {
	var (
		db  = &Database{}
		err error
	)

	db.Master, err = gormx.NewClient(cfg.Database.Master)
	if err != nil {
		return nil, err
	}

	db.Slave, err = gormx.NewClient(cfg.Database.Slave)
	if err != nil {
		return nil, err
	}

	return db, nil
}
