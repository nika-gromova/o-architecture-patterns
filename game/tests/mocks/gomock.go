package mocks

//go:generate mockgen -destination moving_obj.go -package mocks github.com/nika-gromova/o-architecture-patterns/game/move MovingObject
//go:generate mockgen -destination rotating_obj.go -package mocks github.com/nika-gromova/o-architecture-patterns/game/rotate RotatingObject
