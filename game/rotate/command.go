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
	Obj RotatingObject
}

func (c *RotateCommand) Execute() error {
	angle, ok := c.Obj.GetAngle()
	if !ok {
		return fmt.Errorf("RotateCommand: angle %w", base.ErrGetProperty)
	}
	velocity, ok := c.Obj.GetAngularVelocity()
	if !ok {
		return fmt.Errorf("RotateCommand: velocity %w", base.ErrGetProperty)
	}

	if !c.Obj.SetAngle(angle.Plus(velocity)) {
		return fmt.Errorf("RotateCommand: angle %w", base.ErrSetProperty)
	}

	return nil
}
