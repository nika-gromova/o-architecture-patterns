package queue

import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/sirupsen/logrus"
)

type ErrorsHandler interface {
	Handle(cmd base.Command, err error) base.Command
}

type Queue struct {
	commands      chan base.Command
	ctx           context.Context
	errorsHandler ErrorsHandler
}

func New(ctx context.Context, handler ErrorsHandler, cmds chan base.Command) *Queue {
	return &Queue{
		ctx:           ctx,
		commands:      cmds,
		errorsHandler: handler,
	}
}

func (q *Queue) Process() error {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered. Error: %s", r)
		}
	}()

	for {
		select {
		case <-q.ctx.Done():
			return q.ctx.Err()
		case cmd := <-q.commands:
			if err := cmd.Execute(); err != nil {
				err = q.errorsHandler.Handle(cmd, err).Execute()
				if err != nil {
					panic(fmt.Errorf("failed to handler error: %w", err))
				}
			}
		}
	}
}
