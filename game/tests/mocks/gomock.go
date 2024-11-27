package mocks

//go:generate mockgen -destination moving_obj.go -package mocks github.com/nika-gromova/o-architecture-patterns/game/move MovingObject
//go:generate mockgen -destination rotating_obj.go -package mocks github.com/nika-gromova/o-architecture-patterns/game/rotate RotatingObject
//go:generate mockgen -destination fuel_obj.go -package mocks github.com/nika-gromova/o-architecture-patterns/game/fuel UsingFuelObject
//go:generate mockgen -destination mf_obj_obj.go -package mocks github.com/nika-gromova/o-architecture-patterns/game/macro_command/move MovingWithFuelObj
//go:generate mockgen -destination mr_obj_obj.go -package mocks github.com/nika-gromova/o-architecture-patterns/game/change_velocity RotatingAndMovingObject

//go:generate mockgen -destination cmd.go -package mocks github.com/nika-gromova/o-architecture-patterns/game/base Command
