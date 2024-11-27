package move

import (
	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/fuel"
	"github.com/nika-gromova/o-architecture-patterns/game/macro_command"
	move2 "github.com/nika-gromova/o-architecture-patterns/game/move"
)

type MovingWithFuelObj interface {
	GetLocation() (base.Vector, bool)
	GetVelocity() (base.Vector, bool)
	SetLocation(base.Vector) bool
	GetFuel() (base.FuelInfo, bool)
	SetFuel(f base.FuelInfo) bool
}

type MoveWithFuelCommand struct {
	Obj  MovingWithFuelObj
	Fuel base.FuelInfo
}

func (c *MoveWithFuelCommand) Execute() error {
	check := &fuel.CheckFuelCommand{
		Obj:        c.Obj,
		NeededFuel: c.Fuel,
	}
	move := &move2.MoveCommand{
		Obj: c.Obj,
	}
	burn := &fuel.BurnFuelCommand{
		Obj:    c.Obj,
		ToBurn: c.Fuel,
	}
	mc := macro_command.MacroCommand{
		Commands: []base.Command{
			check,
			move,
			burn,
		},
	}
	return mc.Execute()
}
