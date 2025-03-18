package types

import (
	"strings"
)

type StringType struct {
	Value string
}

func NewStringTypeFromString(s string) (*StringType, error) {
	return &StringType{
		Value: s,
	}, nil
}

func (s *StringType) Equals(value Comparable) bool {
	str, ok := value.(*StringType)
	if !ok {
		return false
	}
	if str == nil {
		return false
	}
	return strings.EqualFold(str.Value, s.Value)
}

func (s *StringType) GreaterThan(_ Comparable) bool {
	return false
}

func (s *StringType) LessThan(_ Comparable) bool {
	return false
}
