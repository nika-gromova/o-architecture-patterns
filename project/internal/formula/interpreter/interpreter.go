package interpreter

import (
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models/types"
)

var (
	ErrTypeConversion = fmt.Errorf("type conversion error")
)

type AbstractExpression[T any] interface {
	Interpret(context models.Data[T]) (bool, error)
}

type NilExpression[T any] struct{}

func (ne *NilExpression[T]) Interpret(context models.Data[T]) (bool, error) {
	return false, nil
}

type AndExpression[T any] struct {
	Left  AbstractExpression[T]
	Right AbstractExpression[T]
}

func (ae *AndExpression[T]) Interpret(context models.Data[T]) (bool, error) {
	left, err := ae.Left.Interpret(context)
	if err != nil {
		return false, err
	}
	right, err := ae.Right.Interpret(context)
	if err != nil {
		return false, err
	}
	return left && right, nil
}

type OrExpression[T any] struct {
	Left  AbstractExpression[T]
	Right AbstractExpression[T]
}

func (oe *OrExpression[T]) Interpret(context models.Data[T]) (bool, error) {
	left, err := oe.Left.Interpret(context)
	if err != nil {
		return false, err
	}
	right, err := oe.Right.Interpret(context)
	if err != nil {
		return false, err
	}
	return left || right, nil
}

type ComparableInterpreter[T any] interface {
	InterpretComparable(context models.Data[T]) (types.Comparable, error)
}

type GraterExpression[T any] struct {
	Left  ComparableInterpreter[T]
	Right ComparableInterpreter[T]
}

func (ge *GraterExpression[T]) Interpret(context models.Data[T]) (bool, error) {
	lValue, err := ge.Left.InterpretComparable(context)
	if err != nil {
		return false, err
	}
	rValue, err := ge.Right.InterpretComparable(context)
	if err != nil {
		return false, err
	}
	return lValue.GreaterThan(rValue), nil
}

type LessExpression[T any] struct {
	Left  ComparableInterpreter[T]
	Right ComparableInterpreter[T]
}

func (le *LessExpression[T]) Interpret(context models.Data[T]) (bool, error) {
	lValue, err := le.Left.InterpretComparable(context)
	if err != nil {
		return false, err
	}
	rValue, err := le.Right.InterpretComparable(context)
	if err != nil {
		return false, err
	}
	return lValue.LessThan(rValue), nil
}

type EqualExpression[T any] struct {
	Left  ComparableInterpreter[T]
	Right ComparableInterpreter[T]
}

func (ee *EqualExpression[T]) Interpret(context models.Data[T]) (bool, error) {
	lValue, err := ee.Left.InterpretComparable(context)
	if err != nil {
		return false, err
	}
	rValue, err := ee.Right.InterpretComparable(context)
	if err != nil {
		return false, err
	}
	return lValue.Equals(rValue), nil
}

type Variable[T any] struct {
	Name string
}

func (v Variable[T]) InterpretComparable(context models.Data[T]) (types.Comparable, error) {
	value, err := context.GetValue(v.Name)
	if err != nil {
		return nil, err
	}
	valueConverted, ok := any(value).(types.Comparable)
	if !ok {
		return nil, ErrTypeConversion
	}
	return valueConverted, nil
}

type Const[T any] struct {
	Value types.Comparable
}

func (c *Const[T]) InterpretComparable(_ models.Data[T]) (types.Comparable, error) {
	return c.Value, nil
}
