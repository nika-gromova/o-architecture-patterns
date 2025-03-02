package formula

import (
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/interpreter"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

type Parser interface {
	Parse(input string) (*models.ParsingNode, error)
}

type Formula struct {
	text                string
	parser              Parser
	context             models.InterpreterContext[any]
	knownVariableTokens map[string]struct{}
	ioc                 map[string]interpreter.Comparable
}

func (f *Formula) Evaluate() (bool, error) {
	parsed, err := f.parser.Parse(f.text)
	if err != nil {
		return false, err
	}

	expression, err := f.buildExpression(parsed)
	if err != nil {
		return false, err
	}

	return expression.Interpret(f.context)
}

func (f *Formula) buildExpression(node *models.ParsingNode) (interpreter.AbstractExpression[any], error) {
	if !node.IsOperator || node.Left == nil || node.Right == nil {
		return &interpreter.NilExpression[any]{}, nil
	}

	return f.toExpression(node.Value, node.Left, node.Left)
}

func (f *Formula) toExpression(operator string, left *models.ParsingNode, right *models.ParsingNode) (interpreter.AbstractExpression[any], error) {
	var (
		leftExpression, rightExpression interpreter.AbstractExpression[any]
		err                             error
	)

	if left.IsOperator {
		leftExpression, err = f.toExpression(left.Value, left.Left, left.Right)
		if err != nil {
			return nil, err
		}
	} else if right.IsOperator {
		rightExpression, err = f.toExpression(right.Value, right.Left, right.Right)
		if err != nil {
			return nil, err
		}
	} else {
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
		var leftValue, rightValue interpreter.ComparableInterpreter[any]
		// leftValue = &interpreter.Variable[any]{
		//				Name: left.Value,
		//			}
		// rightValue = &interpreter.Variable[any]{
		//				Name: right.Value,
		//			}

		// TODO create ioc container, IoC.Resolve("Variables.VariableName", value) any

		if leftValue == nil {
			leftValue = &interpreter.Const[any]{
				Value: f.ioc[variableName], // (left.value)// TODO create ioc container, IoC.Resolve("VariableName", value) Comparable
			}
		}
		if rightValue == nil {
			rightValue = &interpreter.Const[any]{
				Value: f.ioc[variableName], // (right.Value)// TODO create ioc container, IoC.Resolve("VariableName", value) Comparable
			}
		}

		// TODO ioc.Resolve("Operators.>", left, right) AbstractExpression
		// inside ioc convert any to needed interface or return error
		switch operator {
		case models.GraterOperator:
			return &interpreter.GraterExpression[any]{
				Left:  leftValue,
				Right: rightValue,
			}, nil
		case models.LessOperator:
			return &interpreter.LessExpression[any]{
				Left:  leftValue,
				Right: rightValue,
			}, nil
		case models.EqualOperator:
			return &interpreter.EqualExpression[any]{
				Left:  leftValue,
				Right: rightValue,
			}, nil
		default:
			return nil, fmt.Errorf("invalid operator: %s", operator)
		}
	}

	// TODO ioc.Resolve("Operators.AND", leftExpression, rightExpression) AbstractExpression
	if operator == models.OrOperator {
		return &interpreter.OrExpression[any]{
			Left:  leftExpression,
			Right: rightExpression,
		}, nil
	}
	if operator == models.ANDOperator {
		return &interpreter.AndExpression[any]{
			Left:  leftExpression,
			Right: rightExpression,
		}, nil
	}

	return nil, fmt.Errorf("invalid operator: %s", operator)
}
