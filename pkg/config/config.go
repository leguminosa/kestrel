package config

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/leguminosa/kestrel/pkg/envx"
	"github.com/leguminosa/kestrel/pkg/gomodfinder"
	"github.com/leguminosa/kestrel/pkg/viperx"
)

var (
	cfg      = &Config{}
	initOnce sync.Once
)

func InitConfig() (*Config, error) {
	var err error

	initOnce.Do(func() {
		err = readConfig(cfg)
		if err != nil {
			return
		}
	})

	return cfg, err
}

func readConfig(cfg *Config) error {
	var (
		basePath    = "/etc"
		serviceName = "kestrel"
		currentEnv  = envx.ServiceEnv()
	)

	if envx.IsDevelopment() {
		localRoot, _ := gomodfinder.Find()
		basePath = filepath.Join(localRoot, "files", basePath)
	}

	err := viperx.NewViper().Read(cfg, filepath.Join(basePath, serviceName), strings.Join([]string{"config", currentEnv}, "."))
	if err != nil {
		return err
	}

	return nil
}
