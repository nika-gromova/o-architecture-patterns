package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/lalamove/konfig"
	"github.com/lalamove/konfig/loader/klfile"
	"github.com/lalamove/konfig/parser/kpyaml"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
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

func (c *Config) GetValue(key string) string {
	return c.cfg.String(key)
}

func (c *Config) GetValues(key string) []string {
	return c.cfg.StringSlice(key)
}

func (c *Config) GetMap(key string) map[string]string {
	pairs := c.cfg.StringSlice(key)
	result := make(map[string]string, len(pairs))
	for _, pair := range pairs {
		split := strings.Split(pair, ":")
		if len(split) != 2 {
			continue
		}
		result[split[0]] = split[1]
	}
	return result
}

func (c *Config) GetHeaderVariables() []*models.HeaderVariable {
	variables := c.cfg.StringMap(VariablesToHeaders)
	result := make([]*models.HeaderVariable, 0, len(variables))
	for key, value := range variables {
		values, ok := value.(map[string]string)
		if !ok {
			continue
		}

		result = append(result, &models.HeaderVariable{
			Name:   key,
			Header: values[VariableHeader],
			Type:   values[VariableType],
		})
	}
	return result
}
