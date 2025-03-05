package interpreter

import (
	"context"
	"fmt"

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
	if ns == nil {
		return &NilExpression[any]{}, nil
	}

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

	expression, err := ioc.Resolve(ctx, fmt.Sprintf("Formula.Interpreter.Operators.%s", ns.Value), left, right)
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
	if nl == nil {
		return &NilExpression[any]{}, nil
	}
	if nl.VariableName == "" {
		return nil, fmt.Errorf("invalid operand: %s", nl.Value)
	}

	expr, err := ioc.Resolve(ctx, "Formula.Interpreter.Variables."+nl.VariableName, nl.Value)
	if err != nil {
		return nil, err
	}

	return expr, nil
}
