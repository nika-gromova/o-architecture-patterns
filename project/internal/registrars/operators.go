package registrars

import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/ioc"
)

type IoCFormulaOperatorsEqualRegistrar struct {
}

func (r *IoCFormulaOperatorsEqualRegistrar) Register(oldCtx context.Context) (context.Context, error) {
	ctx := ioc.NewScope(oldCtx)

	err := ioc.Register(ctx, models.IoCFormulaInterpreterOperatorsDomain+models.EqualOperator, func(args ...any) (any, error) {
		leftTyped, rightTyped, errIn := forComparableOperator(args...)
		if errIn != nil {
			return nil, errIn
		}
		return &interpreter.EqualExpression[any]{
			Left:  leftTyped,
			Right: rightTyped,
		}, nil
	})
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

type IoCFormulaOperatorsGreaterRegistrar struct {
}

func (r *IoCFormulaOperatorsGreaterRegistrar) Register(oldCtx context.Context) (context.Context, error) {
	ctx := ioc.NewScope(oldCtx)

	err := ioc.Register(ctx, models.IoCFormulaInterpreterOperatorsDomain+models.GraterOperator, func(args ...any) (any, error) {
		leftTyped, rightTyped, errIn := forComparableOperator(args...)
		if errIn != nil {
			return nil, errIn
		}
		return &interpreter.GraterExpression[any]{
			Left:  leftTyped,
			Right: rightTyped,
		}, nil
	})
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

type IoCFormulaOperatorsLessRegistrar struct {
}

func (r *IoCFormulaOperatorsLessRegistrar) Register(oldCtx context.Context) (context.Context, error) {
	ctx := ioc.NewScope(oldCtx)

	err := ioc.Register(ctx, models.IoCFormulaInterpreterOperatorsDomain+models.LessOperator, func(args ...any) (any, error) {
		leftTyped, rightTyped, errIn := forComparableOperator(args...)
		if errIn != nil {
			return nil, errIn
		}
		return &interpreter.LessExpression[any]{
			Left:  leftTyped,
			Right: rightTyped,
		}, nil
	})
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

type IoCFormulaOperatorsOrRegistrar struct {
}

func (r *IoCFormulaOperatorsOrRegistrar) Register(oldCtx context.Context) (context.Context, error) {
	ctx := ioc.NewScope(oldCtx)

	err := ioc.Register(ctx, models.IoCFormulaInterpreterOperatorsDomain+models.OrOperator, func(args ...any) (any, error) {
		leftTyped, rightTyped, errIn := forAbstractOperator(args...)
		if errIn != nil {
			return nil, errIn
		}
		return &interpreter.OrExpression[any]{
			Left:  leftTyped,
			Right: rightTyped,
		}, nil
	})
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

type IoCFormulaOperatorsAndRegistrar struct {
}

func (r *IoCFormulaOperatorsAndRegistrar) Register(oldCtx context.Context) (context.Context, error) {
	ctx := ioc.NewScope(oldCtx)

	err := ioc.Register(ctx, models.IoCFormulaInterpreterOperatorsDomain+models.ANDOperator, func(args ...any) (any, error) {
		leftTyped, rightTyped, errIn := forAbstractOperator(args...)
		if errIn != nil {
			return nil, errIn
		}
		return &interpreter.AndExpression[any]{
			Left:  leftTyped,
			Right: rightTyped,
		}, nil
	})
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func forComparableOperator(args ...any) (interpreter.ComparableInterpreter[any], interpreter.ComparableInterpreter[any], error) {
	if len(args) != 2 {
		return nil, nil, fmt.Errorf("expected 2 arguments, got %d", len(args))
	}
	left := args[0]
	right := args[1]
	leftTyped, ok := left.(interpreter.ComparableInterpreter[any])
	if !ok {
		return nil, nil, fmt.Errorf("failed to convert left expression to comparable")
	}
	rightTyped, ok := right.(interpreter.ComparableInterpreter[any])
	if !ok {
		return nil, nil, fmt.Errorf("failed to convert right expression to comparable")
	}
	return leftTyped, rightTyped, nil
}

func forAbstractOperator(args ...any) (interpreter.AbstractExpression[any], interpreter.AbstractExpression[any], error) {
	if len(args) != 2 {
		return nil, nil, fmt.Errorf("expected 2 arguments, got %d", len(args))
	}
	left := args[0]
	right := args[1]
	leftTyped, ok := left.(interpreter.AbstractExpression[any])
	if !ok {
		return nil, nil, fmt.Errorf("failed to convert left expression to abstrat")
	}
	rightTyped, ok := right.(interpreter.AbstractExpression[any])
	if !ok {
		return nil, nil, fmt.Errorf("failed to convert right expression to abstract")
	}
	return leftTyped, rightTyped, nil
}
