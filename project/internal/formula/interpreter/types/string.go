package types

import "github.com/nika-gromova/o-architecture-patterns/project/internal/models"

type StringType struct {
	Value string
}

func (s *StringType) Equals(value models.Comparable) bool {
	str, ok := value.(*StringType)
	if !ok {
		return false
	}
	if str == nil {
		return false
	}
	return str.Value == s.Value
}

func (s *StringType) GreaterThan(_ models.Comparable) bool {
	return false
}

func (s *StringType) LessThan(_ models.Comparable) bool {
	return false
}
