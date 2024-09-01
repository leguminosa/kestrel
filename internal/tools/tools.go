package tools

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxClient interface {
	GetMaster(ctx context.Context) (*pgxpool.Pool, error)
	GetSlave(ctx context.Context) (*pgxpool.Pool, error)
	Close()
}
