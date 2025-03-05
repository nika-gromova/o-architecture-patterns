package formula

import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

type Parser interface {
	Parse(input string) (*models.ParsingNode, error)
}

type Storage interface {
	IsKnownVariableToken(string) bool
}

type Cache interface {
	Set(key string, expression interpreter.AbstractExpression[any])
	Get(key string) (expression interpreter.AbstractExpression[any])
}

type Processor struct {
	parser  Parser
	storage Storage
	cache   Cache
}

func New(parser Parser, storage Storage, cache Cache) *Processor {
	return &Processor{
		parser:  parser,
		storage: storage,
		cache:   cache,
	}
}

func (p *Processor) Evaluate(ctx context.Context, input string, data models.Data[any]) (bool, error) {
	expression := p.cache.Get(input)
	if expression == nil {
		var err error
		expression, err = p.buildExpression(ctx, input)
		if err != nil {
			return false, fmt.Errorf("failed to build expression for `%s`: %w", input, err)
		}
		p.cache.Set(input, expression)
	}

	return expression.Interpret(data)
}

func (p *Processor) buildExpression(ctx context.Context, input string) (interpreter.AbstractExpression[any], error) {
	parsed, err := p.parser.Parse(input)
	if err != nil {
		return nil, err
	}

	expression, err := p.toExpressionNode(parsed).ToExpression(ctx)
	if err != nil {
		return nil, err
	}

	result, ok := expression.(interpreter.AbstractExpression[any])
	if !ok {
		return nil, fmt.Errorf("failed to convert result expression to abstract, input: %s", input)
	}
	return result, nil
}

func (p *Processor) toExpressionNode(node *models.ParsingNode) interpreter.ExpressionNode {
	if !node.IsOperator || node.Left == nil || node.Right == nil {
		return nil
	}

	return p.toExpression(node.Value, node.Left, node.Right)
}

func (p *Processor) toExpression(operator string, left *models.ParsingNode, right *models.ParsingNode) interpreter.ExpressionNode {
	var leftExpression, rightExpression interpreter.ExpressionNode

	if left.IsOperator {
		leftExpression = p.toExpression(left.Value, left.Left, left.Right)
	}
	if right.IsOperator {
		rightExpression = p.toExpression(right.Value, right.Left, right.Right)
	}
	if !left.IsOperator && !right.IsOperator {
		// determine the variable - time, locale, etc
		var variableName string
		if p.storage.IsKnownVariableToken(left.Value) {
			variableName = left.Value
		}
		if p.storage.IsKnownVariableToken(right.Value) {
			variableName = right.Value
		}

		if variableName == "" {
			return nil
		}

		leftExpression = &interpreter.NodeLeaf{
			Value:        left.Value,
			VariableName: variableName,
		}
		rightExpression = &interpreter.NodeLeaf{
			Value:        right.Value,
			VariableName: variableName,
		}
	}

	return &interpreter.NodeOperator{
		Value: operator,
		Left:  leftExpression,
		Right: rightExpression,
	}
}
