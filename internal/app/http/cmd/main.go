package cmd

import (
	"context"
	"log"

	"github.com/leguminosa/kestrel/internal/app/http"
	"github.com/leguminosa/kestrel/internal/config"
)

func Main() {
	var ctx = context.Background()

	// initialize config
	cfg := &config.Config{}

	cfg.Server.Port = ":9000"

	cfg.DatabaseConfig.ConnectionString = "user=kestrel_user password=magical_password dbname=counterside host=127.0.0.1 port=5432 sslmode=disable pool_health_check_period=10s"

	log.Println("http running on port", cfg.Server.Port)
	log.Fatalln(http.InitHTTPServer(ctx, cfg).Start())
}
