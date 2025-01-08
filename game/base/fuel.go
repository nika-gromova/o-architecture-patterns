package base

type FuelInfo struct {
	Value int32
	Capacity     uint64
}

func (f FuelInfo) Less(other FuelInfo) bool {
	return f.Value < other.Value
}

func (f FuelInfo) Burn(other FuelInfo) FuelInfo {
	f.Value -= other.Value
	return f
}
