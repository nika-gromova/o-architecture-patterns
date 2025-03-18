package types

import (
	"time"
)

type DateTimeType struct {
	Value time.Time
}

func NewDateTimeTypeFromString(str string) (*DateTimeType, error) {
	result, err := time.Parse(time.RFC1123, str)
	if err != nil {
		return nil, err
	}
	return &DateTimeType{
		Value: result,
	}, nil
}

func (dt *DateTimeType) Equals(value Comparable) bool {
	valueTime := convertTimeFromAny(value)

	if valueTime.IsZero() {
		return false
	}
	return dt.Value.Equal(valueTime)
}

func (dt *DateTimeType) GreaterThan(value Comparable) bool {
	valueTime := convertTimeFromAny(value)

	if valueTime.IsZero() {
		return false
	}
	return dt.Value.After(valueTime)
}

func (dt *DateTimeType) LessThan(value Comparable) bool {
	valueTime := convertTimeFromAny(value)

	if valueTime.IsZero() {
		return false
	}
	return dt.Value.Before(valueTime)
}

func convertTimeFromAny(value any) time.Time {
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
