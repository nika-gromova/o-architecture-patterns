package rotate

import (
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/change_velocity"
	"github.com/nika-gromova/o-architecture-patterns/game/macro_command"
	rotate2 "github.com/nika-gromova/o-architecture-patterns/game/rotate"
)

type RotatingAndMovingObject interface {
	GetAngle() (base.Angle, bool)
	GetAngularVelocity() (base.Angle, bool)
	SetAngle(base.Angle) bool
	GetVelocityVector() (base.Vector, bool)
	SetVelocityVector(vector base.Vector) bool
}

type RotateCommand struct {
	Obj RotatingAndMovingObject
}

func (c *RotateCommand) Execute() error {
	rotate := &rotate2.RotateCommand{
		Obj: c.Obj,
	}
	angle, ok := c.Obj.GetAngularVelocity()
	if !ok {
		return fmt.Errorf("RotateCommand: %w", base.ErrGetProperty)
	}
	changeVelocity := &change_velocity.ChangeVelocityCommand{
		Obj:   c.Obj,
		Angle: angle,
	}
	mc := macro_command.MacroCommand{
		Commands: []base.Command{
			rotate,
			changeVelocity,
		},
	}
	return mc.Execute()
}
