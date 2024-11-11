package base

type Vector struct {
	Coordinates []int
}

func (v Vector) Plus(other Vector) Vector {
	if len(v.Coordinates) != len(other.Coordinates) {
		return v
	}
	result := Vector{
		Coordinates: make([]int, len(v.Coordinates)),
	}
	for i := 0; i < len(v.Coordinates); i++ {
		result.Coordinates[i] = v.Coordinates[i] + other.Coordinates[i]
	}
	return result
}
