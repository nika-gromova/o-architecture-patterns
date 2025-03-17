package registrars

import (
	"context"
	"fmt"
	"strings"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models/types"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/ioc"
)

type IoCFormulaStringVariableRegistrar struct {
	VariableName string
}

func (r *IoCFormulaStringVariableRegistrar) Register(oldCtx context.Context) (context.Context, error) {
	ctx := ioc.NewScope(oldCtx)

	err := ioc.Register(ctx, models.IoCFormulaInterpreterVariablesDomain+r.VariableName, func(args ...any) (any, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
		}
		value, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("failed to convert value to string")
		}

		if strings.EqualFold(value, r.VariableName) {
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

	return ctx, nil
}

type IoCFormulaDateTimeVariableRegistrar struct {
	VariableName string
}

func (r *IoCFormulaDateTimeVariableRegistrar) Register(oldCtx context.Context) (context.Context, error) {
	ctx := ioc.NewScope(oldCtx)

	err := ioc.Register(ctx, models.IoCFormulaInterpreterVariablesDomain+r.VariableName, func(args ...any) (any, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("expected 1 argument, got %d", len(args))
		}
		value, ok := args[0].(string)
		if !ok {
			return nil, fmt.Errorf("failed to convert value to string")
		}

		if strings.EqualFold(value, r.VariableName) {
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

	return ctx, nil
}
