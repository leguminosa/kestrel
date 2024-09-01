package http

import (
	"context"

	"github.com/leguminosa/kestrel/internal/app/http/handler"
	"github.com/leguminosa/kestrel/internal/config"
	character_module "github.com/leguminosa/kestrel/internal/module/character"
	character_repo "github.com/leguminosa/kestrel/internal/repository/character"
	"github.com/leguminosa/kestrel/internal/storage/database/postgres"
	"github.com/leguminosa/kestrel/internal/tools/database/pgx"
)

func InitHTTPServer(ctx context.Context, cfg *config.Config) *serverImpl {
	dbClient, err := pgx.NewPgxClient(ctx, cfg.DatabaseConfig)
	if err != nil {
		return newServerWithError(err)
	}
	// defer dbClient.Close()

	characterDB := postgres.NewCharacterDB(dbClient)
	characterRepo := character_repo.NewCharacterRepository(characterDB)
	characterModule := character_module.NewCharacterModule(characterRepo)
	characterHandler := handler.NewCharacterHandler(characterModule)

	echo := provideHTTPServer(&controller{
		characterHandler: characterHandler,
	})

	return newServer(*cfg, echo)
}
