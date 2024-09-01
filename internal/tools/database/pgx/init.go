package pgx

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/leguminosa/kestrel/internal/config"
)

type PgxPoolWrapper struct {
	master, slave *pgxpool.Pool
}

func NewPgxClient(ctx context.Context, cfg config.DatabaseConfig) (*PgxPoolWrapper, error) {
	masterPool, err := connectPgx(ctx, cfg)
	if err != nil {
		return nil, err
	}

	slavePool, err := connectPgx(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &PgxPoolWrapper{
		master: masterPool,
		slave:  slavePool,
	}, nil
}

func connectPgx(ctx context.Context, cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	dbCfg, err := pgxpool.ParseConfig(cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, dbCfg)
	if err != nil {
		return nil, err
	}

	c, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Release()

	err = c.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
