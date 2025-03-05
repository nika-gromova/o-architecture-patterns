package formula

import (
	"context"
	"fmt"
	"strings"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models/types"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/ioc"
)

type IoCFormulaOperatorsRegistrar struct {
	next models.Registrar
}

func (r *IoCFormulaOperatorsRegistrar) Register(oldCtx context.Context) (context.Context, error) {
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

	err = ioc.Register(ctx, models.IoCFormulaInterpreterOperatorsDomain+models.GraterOperator, func(args ...any) (any, error) {
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

	err = ioc.Register(ctx, models.IoCFormulaInterpreterOperatorsDomain+models.LessOperator, func(args ...any) (any, error) {
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

	err = ioc.Register(ctx, models.IoCFormulaInterpreterOperatorsDomain+models.OrOperator, func(args ...any) (any, error) {
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

	err = ioc.Register(ctx, models.IoCFormulaInterpreterOperatorsDomain+models.ANDOperator, func(args ...any) (any, error) {
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

	if r.next != nil {
		return r.next.Register(ctx)
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

type IoCFormulaStringVariableRegistrar struct {
	variableName string
	next         models.Registrar
}

func (r *IoCFormulaStringVariableRegistrar) Register(oldCtx context.Context) (context.Context, error) {
	ctx := ioc.NewScope(oldCtx)

	err := ioc.Register(ctx, models.IoCFormulaInterpreterVariablesDomain+r.variableName, func(args ...any) (any, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
		}
		value, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("failed to convert value to string")
		}

		if strings.EqualFold(value, r.variableName) {
			return &interpreter.Variable[any]{
				Name: value,
			}, nil
		}
		constValue, err := types.NewStringTypeFromString(value)
		if err != nil {
			return nil, err
		}
		return &interpreter.Const[any]{
			Value: constValue,
		}, nil
	})

	if err != nil {
		return nil, err
	}

	if r.next != nil {
		return r.next.Register(ctx)
	}
	return ctx, nil
}

type IoCFormulaDateTimeVariableRegistrar struct {
	variableName string
	next         models.Registrar
}

func (r *IoCFormulaDateTimeVariableRegistrar) Register(oldCtx context.Context) (context.Context, error) {
	ctx := ioc.NewScope(oldCtx)

	err := ioc.Register(ctx, models.IoCFormulaInterpreterVariablesDomain+r.variableName, func(args ...any) (any, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
		}
		value, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("failed to convert value to string")
		}

		if strings.EqualFold(value, r.variableName) {
			return &interpreter.Variable[any]{
				Name: value,
			}, nil
		}
		constValue, err := types.NewDateTimeTypeFromString(value)
		if err != nil {
			return nil, err
		}
		return &interpreter.Const[any]{
			Value: constValue,
		}, nil
	})

	if err != nil {
		return nil, err
	}

	if r.next != nil {
		return r.next.Register(ctx)
	}
	return ctx, nil
}
