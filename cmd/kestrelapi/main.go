package main

import (
	"flag"
	"log"

	"github.com/leguminosa/kestrel/internal/app/http"
	"github.com/leguminosa/kestrel/pkg/config"
	"github.com/leguminosa/kestrel/pkg/database"
)

var (
	configTest bool
)

func init() {
	// read flags (logging file location, config test, debug mode, etc)
	flag.BoolVar(&configTest, "t", false, "")
	flag.Parse()
}

func main() {
	var err error

	var cfg *config.Config
	cfg, err = config.InitConfig()
	if err != nil {
		log.Fatalln("error config", err)
	}

	var db *database.Database
	db, err = database.InitDatabase(cfg)
	if err != nil {
		log.Fatalln("error init db", err)
	}

	if configTest {
		log.Println("config test success")
		return
	}

	log.Println("running on port", cfg.Port)
	log.Fatalln(http.InitApp(cfg, db).Run())
}
