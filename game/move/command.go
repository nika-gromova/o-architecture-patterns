package move

import (
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
)

type MovingObject interface {
	GetLocation() (base.Vector, bool)
	GetVelocity() (base.Vector, bool)
	SetLocation(base.Vector) bool
}

type MoveCommand struct {
	obj MovingObject
}

func (c *MoveCommand) Execute() error {
	location, ok := c.obj.GetLocation()
	if !ok {
		return fmt.Errorf("MoveCommand: location %w", base.ErrGetProperty)
	}
	velocity, ok := c.obj.GetVelocity()
	if !ok {
		return fmt.Errorf("MoveCommand: velocity %w", base.ErrGetProperty)
	}
	if !c.obj.SetLocation(location.Plus(velocity)) {
		return fmt.Errorf("MoveCommand: location %w", base.ErrSetProperty)
	}
	return nil
}
