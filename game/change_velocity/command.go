package change_velocity

import (
	"fmt"
	"math"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	log "github.com/sirupsen/logrus"
)

type RotatingAndMovingObject interface {
	GetVelocityVector() (base.Vector, bool)
	SetVelocityVector(vector base.Vector) bool
}

type ChangeVelocityCommand struct {
	Obj   RotatingAndMovingObject
	Angle base.Angle
}

// поворот по часовой стрелке
func (c ChangeVelocityCommand) Execute() error {
	velocity, ok := c.Obj.GetVelocityVector()
	if !ok {
		log.Errorf("ChangeVelocityCommand: %s", base.ErrGetProperty)
		return nil
	}
	if len(velocity.Coordinates) < 2 {
		return base.ErrCommandExecution
	}
	angle := c.Angle.ToDouble()
	angleCos := math.Cos(angle)
	angleSin := math.Sin(angle)
	newVelocity := base.Vector{
		Coordinates: []int{
			int(math.Round(float64(velocity.Coordinates[0])*angleCos + float64(velocity.Coordinates[1])*angleSin)),
			int(math.Round(float64(-velocity.Coordinates[0])*angleSin + float64(velocity.Coordinates[1])*angleCos)),
		},
	}
	ok = c.Obj.SetVelocityVector(newVelocity)
	if !ok {
		return fmt.Errorf("ChangeVelocityCommand: %w", base.ErrSetProperty)
	}
	return nil
}
