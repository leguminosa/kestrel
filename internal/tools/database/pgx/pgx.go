package pgx

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var errEmptyDatabasePool = errors.New("empty pgx pool connection")

func (db *PgxPoolWrapper) GetMaster(ctx context.Context) (*pgxpool.Pool, error) {
	if db.master == nil {
		return nil, errEmptyDatabasePool
	}

	return db.master, nil
}

func (db *PgxPoolWrapper) GetSlave(ctx context.Context) (*pgxpool.Pool, error) {
	if db.slave == nil {
		return nil, errEmptyDatabasePool
	}

	return db.slave, nil
}

func (db *PgxPoolWrapper) Close() {
	if db.master != nil {
		log.Println("closing master pgx pool connection")
		db.master.Close()
	}

	if db.slave != nil {
		log.Println("closing slave pgx pool connection")
		db.slave.Close()
	}
}
