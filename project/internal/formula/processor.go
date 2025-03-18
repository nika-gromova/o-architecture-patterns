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

type Cache interface {
	Set(key string, object any)
	Get(key string) (any, bool)
}

type Processor struct {
	parser Parser
	cache  Cache
}

func New(parser Parser, cache Cache) *Processor {
	return &Processor{
		parser: parser,
		cache:  cache,
	}
}

func (p *Processor) Evaluate(ctx context.Context, input *models.Formula, data models.Data[any]) (bool, error) {
	var expression interpreter.AbstractExpression[any]
	object, exists := p.cache.Get(input.Expression)
	if exists {
		expression = object.(interpreter.AbstractExpression[any])
	}
	if expression == nil {
		var err error
		expression, err = p.buildExpression(ctx, input.Expression)
		if err != nil {
			return false, fmt.Errorf("failed to build expression for `%s`: %w", input.Expression, err)
		}
		p.cache.Set(input.Expression, expression)
	}

	return expression.Interpret(data)
}

func (p *Processor) buildExpression(ctx context.Context, input string) (interpreter.AbstractExpression[any], error) {
	parsed, err := p.parser.Parse(input)
	if err != nil {
		return nil, err
	}

	expressionNode := p.toExpressionNode(ctx, parsed)
	if expressionNode == nil {
		return nil, fmt.Errorf("failed to build expression node for `%s`", input)
	}
	expression, err := expressionNode.ToExpression(ctx)
	if err != nil {
		return nil, err
	}

	result, ok := expression.(interpreter.AbstractExpression[any])
	if !ok {
		return nil, fmt.Errorf("failed to convert result expression to abstract, input: %s", input)
	}
	return result, nil
}

func (p *Processor) toExpressionNode(ctx context.Context, node *models.ParsingNode) interpreter.ExpressionNode {
	if !node.IsOperator || node.Left == nil || node.Right == nil {
		return nil
	}

	return p.toExpression(ctx, node.Value, node.Left, node.Right)
}

func (p *Processor) toExpression(ctx context.Context, operator string, left *models.ParsingNode, right *models.ParsingNode) interpreter.ExpressionNode {
	var leftExpression, rightExpression interpreter.ExpressionNode

	if left.IsOperator {
		leftExpression = p.toExpression(ctx, left.Value, left.Left, left.Right)
	}
	if right.IsOperator {
		rightExpression = p.toExpression(ctx, right.Value, right.Left, right.Right)
	}
	if !left.IsOperator && !right.IsOperator {
		// determine the variable - time, locale, etc
		var variableName string
		if p.isKnownVariableToken(ctx, left.Value) {
			variableName = left.Value
		}
		if p.isKnownVariableToken(ctx, right.Value) {
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

func (p *Processor) isKnownVariableToken(ctx context.Context, token string) bool {
	_, err := ioc.Resolve(ctx, models.IoCFormulaInterpreterVariablesDomain+token, token)
	return err == nil
}
