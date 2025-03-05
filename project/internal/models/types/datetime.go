package types

import (
	"time"
)

type DateTimeType struct {
	Value time.Time
}

func (dt *DateTimeType) Equals(value Comparable) bool {
	valueTime := convertTime(value)

	if valueTime.IsZero() {
		return false
	}
	return dt.Value.Equal(valueTime)
}

func (dt *DateTimeType) GreaterThan(value Comparable) bool {
	valueTime := convertTime(value)

	if valueTime.IsZero() {
		return false
	}
	return dt.Value.After(valueTime)
}

func (dt *DateTimeType) LessThan(value Comparable) bool {
	valueTime := convertTime(value)

	if valueTime.IsZero() {
		return false
	}
	return dt.Value.Before(valueTime)
}

func convertTime(value any) time.Time {
	var result time.Time
	valueTime, ok := value.(*DateTimeType)
	if !ok {
		return result
	}
	if valueTime == nil {
		return result
	}
	return valueTime.Value
}
