package ioc

import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter/types"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/ioc"
)

// TODO prettify, add all operators

func InitForFormula() (context.Context, error) {
	container := ioc.New()
	ctx := container.NewScope(context.Background())

	err := ioc.Register(ctx, "Formula.Interpreter.Operators."+models.EqualOperator, func(args ...any) (any, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("expected 2 arguments, got %d", len(args))
		}
		left := args[0]
		right := args[1]
		leftTyped, ok := left.(interpreter.ComparableInterpreter[any])
		if !ok {
			return nil, fmt.Errorf("failed to convert left expression to comparable")
		}
		rightTyped, ok := right.(interpreter.ComparableInterpreter[any])
		if !ok {
			return nil, fmt.Errorf("failed to convert right expression to comparable")
		}
		return &interpreter.EqualExpression[any]{
			Left:  leftTyped,
			Right: rightTyped,
		}, nil
	})

	if err != nil {
		return ctx, err
	}

	err = ioc.Register(ctx, "Formula.Interpreter.Operators."+models.OrOperator, func(args ...any) (any, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("expected 2 arguments, got %d", len(args))
		}
		left := args[0]
		right := args[1]
		leftTyped, ok := left.(interpreter.AbstractExpression[any])
		if !ok {
			return nil, fmt.Errorf("failed to convert left expression to abstrat")
		}
		rightTyped, ok := right.(interpreter.AbstractExpression[any])
		if !ok {
			return nil, fmt.Errorf("failed to convert right expression to abstract")
		}
		return &interpreter.OrExpression[any]{
			Left:  leftTyped,
			Right: rightTyped,
		}, nil
	})

	if err != nil {
		return ctx, err
	}

	err = ioc.Register(ctx, "Formula.Interpreter.Variables.Locale", func(args ...any) (any, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
		}
		value, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("failed to convert value to string")
		}
		if value == "Locale" {
			return &interpreter.Variable[any]{
				Name: value,
			}, nil
		}
		return &interpreter.Const[any]{
			Value: &types.StringType{
				Value: value,
			},
		}, nil
	})

	if err != nil {
		return nil, err
	}

	return ctx, nil
}
