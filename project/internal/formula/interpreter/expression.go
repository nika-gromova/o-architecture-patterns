package interpreter

import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/ioc"
)

type ExpressionNode interface {
	ToExpression(ctx context.Context) (any, error)
}

type NodeOperator struct {
	Value string
	Left  ExpressionNode
	Right ExpressionNode
}

func (ns *NodeOperator) ToExpression(ctx context.Context) (any, error) {
	var (
		left, right any
		err         error
	)
	if ns.Left != nil {
		left, err = ns.Left.ToExpression(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		left = &NilExpression[any]{}
	}

	if ns.Right != nil {
		right, err = ns.Right.ToExpression(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		right = &NilExpression[any]{}
	}

	expression, err := ioc.Resolve(ctx, models.IoCFormulaInterpreterOperatorsDomain+ns.Value, left, right)
	if err != nil {
		return nil, err
	}

	return expression, nil
}

type NodeLeaf struct {
	Value        string
	VariableName string
}

func (nl *NodeLeaf) ToExpression(ctx context.Context) (any, error) {
	expr, err := ioc.Resolve(ctx, models.IoCFormulaInterpreterVariablesDomain+nl.VariableName, nl.Value)
	if err != nil {
		return nil, fmt.Errorf("invalid operand: %w", err)
	}

	return expr, nil
}
