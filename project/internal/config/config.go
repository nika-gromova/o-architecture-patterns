package config

import (
	"fmt"
	"os"
	"time"

	"github.com/lalamove/konfig"
	"github.com/lalamove/konfig/loader/klfile"
	"github.com/lalamove/konfig/parser/kpyaml"
	log "github.com/sirupsen/logrus"
)

const configPathEnv = "CONFIG_PATH"

type Config struct {
	cfg konfig.Store
}

func New() *Config {
	s := konfig.New(konfig.DefaultConfig())

	fileLoader := klfile.New(&klfile.Config{
		Files: []klfile.File{
			{
				Path:   os.Getenv(configPathEnv),
				Parser: kpyaml.Parser,
			},
		},
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

func (c *Config) GetValue(key string) string {
	return c.cfg.String(key)
}

func (c *Config) GetInt(key string) int {
	return c.cfg.Int(key)
}
