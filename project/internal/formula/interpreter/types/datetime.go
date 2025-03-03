package types

import (
	"time"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

type DateTimeType struct {
	Value time.Time
}

func (dt *DateTimeType) Equals(value models.Comparable) bool {
	valueTime := convertTime(value)

	if valueTime.IsZero() {
		return false
	}
	return dt.Value.Equal(valueTime)
}

func (dt *DateTimeType) GreaterThan(value models.Comparable) bool {
	valueTime := convertTime(value)

	if valueTime.IsZero() {
		return false
	}
	return dt.Value.After(valueTime)
}

func (dt *DateTimeType) LessThan(value models.Comparable) bool {
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
