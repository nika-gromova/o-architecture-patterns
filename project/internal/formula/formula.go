package formula

import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/ioc"
)

type Parser interface {
	Parse(input string) (*models.ParsingNode, error)
}

type Formula struct {
	text                string
	parser              Parser
	context             models.InterpreterContext[any]
	knownVariableTokens map[string]struct{}
}

func (f *Formula) Evaluate() (bool, error) {
	parsed, err := f.parser.Parse(f.text)
	if err != nil {
		return false, err
	}

	expression, err := f.buildExpression(context.Background(), parsed) // TODO
	if err != nil {
		return false, err
	}

	return expression.Interpret(f.context)
}

func (f *Formula) buildExpression(ctx context.Context, node *models.ParsingNode) (interpreter.AbstractExpression[any], error) {
	if !node.IsOperator || node.Left == nil || node.Right == nil {
		return &interpreter.NilExpression[any]{}, nil
	}

	return f.toExpression(ctx, node.Value, node.Left, node.Right)
}

func (f *Formula) toExpression(ctx context.Context, operator string, left *models.ParsingNode, right *models.ParsingNode) (interpreter.AbstractExpression[any], error) {
	var (
		leftExpression, rightExpression any
		err                             error
	)

	if left.IsOperator {
		leftExpression, err = f.toExpression(ctx, left.Value, left.Left, left.Right)
		if err != nil {
			return nil, err
		}
	}
	if right.IsOperator {
		rightExpression, err = f.toExpression(ctx, right.Value, right.Left, right.Right)
		if err != nil {
			return nil, err
		}
	}
	if !left.IsOperator && !right.IsOperator {
		var variableName string
		if _, known := f.knownVariableTokens[left.Value]; known {
			variableName = left.Value
		}
		if _, known := f.knownVariableTokens[right.Value]; known {
			variableName = right.Value
		}
		if variableName == "" {
			return nil, fmt.Errorf("invalid operands: %s, %s", left.Value, right.Value)
		}

		leftExpression, err = ioc.Resolve(ctx, "Formula.Interpreter.Variables."+variableName, left.Value)
		if err != nil {
			return nil, err
		}
		rightExpression, err = ioc.Resolve(ctx, "Formula.Interpreter.Variables."+variableName, right.Value)
		if err != nil {
			return nil, err
		}
	}

	if leftExpression == nil {
		leftExpression = &interpreter.NilExpression[any]{}
	}
	if rightExpression == nil {
		rightExpression = &interpreter.NilExpression[any]{}
	}

	var (
		result    interpreter.AbstractExpression[any]
		tmpResult any
		ok        bool
	)

	tmpResult, err = ioc.Resolve(ctx, fmt.Sprintf("Formula.Interpreter.Operators.%s", operator), leftExpression, rightExpression)
	if err != nil {
		return nil, err
	}

	result, ok = tmpResult.(interpreter.AbstractExpression[any])
	if !ok {
		return nil, fmt.Errorf("failed to convert result expression to abstract, operator: %s", operator)
	}
	return result, nil
}
