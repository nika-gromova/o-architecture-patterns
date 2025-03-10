package config

import (
	"fmt"
	"time"

	"github.com/lalamove/konfig"
	"github.com/lalamove/konfig/loader/klfile"
	"github.com/lalamove/konfig/parser/kpyaml"
	log "github.com/sirupsen/logrus"
)

var configFiles = []klfile.File{
	{
		Path:   "config.yaml",
		Parser: kpyaml.Parser,
	},
}

type Config struct {
	cfg konfig.Store
}

func New() *Config {
	s := konfig.New(konfig.DefaultConfig())

	fileLoader := klfile.New(&klfile.Config{
		Files: configFiles,
		Watch: true,
		Rate:  1 * time.Second,
	})
	s.RegisterLoader(fileLoader)
	if err := s.Load(); err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	return &Config{
		cfg: s,
	}
}

func (c *Config) GetSecret(key string) string {
	secret := c.cfg.Get(fmt.Sprintf("secrets.%s", key))
	secretValue, ok := secret.(string)
	if !ok {
		return ""
	}
	return secretValue
}
