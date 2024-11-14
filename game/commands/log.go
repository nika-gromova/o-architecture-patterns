package commands

import (
	"reflect"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	log "github.com/sirupsen/logrus"
)

type LogCommand struct {
	Err error
	Cmd base.Command
}

func (c *LogCommand) Execute() error {
	log.Errorf("failed to execute %s: %s", reflect.TypeOf(c.Cmd).Name(), c.Err)
	return nil
}
