package base

import "math"

type Angle struct {
	Direction  int
	TotalCount int
}

func (a Angle) Plus(other Angle) Angle {
	coefficient := float64(a.TotalCount) / float64(other.TotalCount)
	da := float64(other.Direction) * coefficient
	return Angle{
		Direction:  (a.Direction + int(math.Round(da))) % a.TotalCount,
		TotalCount: a.TotalCount,
	}
}

func (a Angle) ToDouble() float64 {
	coefficient := float64(360) / float64(a.TotalCount)
	degrees := float64(a.Direction) * coefficient
	return degrees * (math.Pi / 180.0)
}
