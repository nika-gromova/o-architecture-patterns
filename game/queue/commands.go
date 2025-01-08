package queue

import (
	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/sirupsen/logrus"
)

type Processor interface {
	Process() error
	Add(cmd base.Command)
	Stop()
}

type Thread struct {
	Processor Processor
	OnStart   func()
	OnStop    func()
}

func (t *Thread) Start() {
	t.OnStart()
	defer t.OnStop()

	if err := t.Processor.Process(); err != nil {
		logrus.Fatalf("Failed to process queue in thread: %s", err)
	}
}

func (t *Thread) Stop() {
	t.Processor.Stop()
}

type StartThreadCommand struct {
	thread *Thread
}

func (c *StartThreadCommand) Execute() error {
	go c.thread.Start()
	return nil
}

type HardStopCommand struct {
	thread *Thread
}

func (c *HardStopCommand) Execute() error {
	c.thread.Stop()
	return nil
}

type SoftStopCommand struct {
	thread *Thread
}

func (c *SoftStopCommand) Execute() error {
	hardStop := &HardStopCommand{
		thread: c.thread,
	}
	c.thread.Processor.Add(hardStop)
	return nil
}
