package base

type Vector struct {
	X int
	Y int
}

func (v Vector) Plus(other Vector) Vector {
	return Vector{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}
