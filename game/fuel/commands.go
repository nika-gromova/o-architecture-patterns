package fuel

import (
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
)

type UsingFuelObject interface {
	GetFuel() (base.FuelInfo, bool)
	SetFuel(f base.FuelInfo) bool
}

type CheckFuelCommand struct {
	Obj        UsingFuelObject
	NeededFuel base.FuelInfo
}

func (c *CheckFuelCommand) Execute() error {
	f, ok := c.Obj.GetFuel()
	if !ok {
		return fmt.Errorf("CheckFuelCommand: fuel %w", base.ErrGetProperty)
	}
	if f.Less(c.NeededFuel) {
		return fmt.Errorf("CheckFuelCommand: %w", base.ErrCommandExecution)
	}
	return nil
}

type BurnFuelCommand struct {
	Obj    UsingFuelObject
	ToBurn base.FuelInfo
}

func (c *BurnFuelCommand) Execute() error {
	f, ok := c.Obj.GetFuel()
	if !ok {
		return fmt.Errorf("BurnFuelCommand: fuel %w", base.ErrGetProperty)
	}
	newValue := f.Burn(c.ToBurn)
	ok = c.Obj.SetFuel(newValue)
	if !ok {
		return fmt.Errorf("BurnFuelCommand: fuel %w", base.ErrSetProperty)
	}
	return nil
}
