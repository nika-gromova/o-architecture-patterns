package rotate

import (
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
)

type RotatingObject interface {
	GetAngle() (base.Angle, bool)
	GetAngularVelocity() (base.Angle, bool)
	SetAngle(base.Angle) bool
}

type RotateCommand struct {
	obj RotatingObject
}

func (c *RotateCommand) Execute() error {
	angle, ok := c.obj.GetAngle()
	if !ok {
		return fmt.Errorf("RotateCommand: angle %w", base.ErrGetProperty)
	}
	velocity, ok := c.obj.GetAngularVelocity()
	if !ok {
		return fmt.Errorf("RotateCommand: velocity %w", base.ErrGetProperty)
	}

	if !c.obj.SetAngle(angle.Plus(velocity)) {
		return fmt.Errorf("RotateCommand: angle %w", base.ErrSetProperty)
	}

	return nil
}
