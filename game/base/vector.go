package base

import "math"

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

func (v Vector) ToInt() int {
	sum := 0
	for i := 0; i < len(v.Coordinates); i++ {
		sum += v.Coordinates[i] * v.Coordinates[i]
	}
	return int(math.Round(math.Sqrt(float64(sum))))
}
