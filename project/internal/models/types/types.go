package types

type Comparable interface {
	Equals(value Comparable) bool
	GreaterThan(value Comparable) bool
	LessThan(value Comparable) bool
}
