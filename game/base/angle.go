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
