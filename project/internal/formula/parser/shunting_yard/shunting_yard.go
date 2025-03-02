package shunting_yard

import (
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/collections"
)

var (
	ErrMismatchedParentheses = fmt.Errorf("mismatched parentheses found")
)

type ShuntingYardStrategy struct{}

func (s *ShuntingYardStrategy) Parse(data *models.ParsingData) (*models.ParsingNode, error) {
	postfix, err := parse(data.OperandsPriorities, data.Tokens)
	if err != nil {
		return nil, fmt.Errorf("unable to parse tokens with Shunting-yard algorithm: %w", err)
	}
	return toNode(postfix)
}

type token struct {
	value      string
	isOperator bool
}

// parses an array of strings and returns an array of abstract tokens using Shunting-yard algorithm
func parse(priorities map[string]int, tokens []string) ([]token, error) {
	var (
		ret       collections.Stack[token]
		operators collections.Stack[string]
	)

	for _, tk := range tokens {
		if _, found := priorities[tk]; !found {
			ret.Push(token{
				value:      tk,
				isOperator: false,
			})
			continue
		}

		switch tk {
		case "(":
			operators.Push(tk)
		case ")":
			foundLeftParenthesis := false
			for isEmpty := operators.IsEmpty(); !isEmpty; {
				// pop until "(" is found
				op := operators.Pop()
				if op == "(" {
					foundLeftParenthesis = true
					break
				} else {
					ret.Push(token{
						value:      op,
						isOperator: true,
					})
				}
			}
			if !foundLeftParenthesis {
				return nil, ErrMismatchedParentheses
			}
		default:
			// operator priority
			priority, ok := priorities[tk]
			if !ok {
				return nil, fmt.Errorf("unknown operator: %v", tk)
			}

			// pop till less priority found
			for isEmpty := operators.IsEmpty(); !isEmpty; {
				op, _ := operators.Top()
				if op == "(" {
					break
				}
				prevPriority := priorities[op]
				if prevPriority >= priority {
					operators.Pop()
					ret.Push(token{
						value:      op,
						isOperator: true,
					})
				} else {
					break
				}
			}
			operators.Push(tk)
		}
	}

	// process remaining operators
	for isEmpty := operators.IsEmpty(); !isEmpty; {
		op := operators.Pop()
		if op == "(" {
			return nil, ErrMismatchedParentheses
		}
		ret.Push(token{
			value:      op,
			isOperator: true,
		})
	}
	return ret.ToSlice(), nil
}

// form tree of nodes from postfix notation
func toNode(postfix []token) (*models.ParsingNode, error) {
	var (
		values collections.Stack[*models.ParsingNode]
	)

	for _, tk := range postfix {
		node := &models.ParsingNode{
			Value:      tk.value,
			IsOperator: tk.isOperator,
		}

		if tk.isOperator {
			right, err := values.Top()
			if err != nil {
				return nil, err
			}
			values.Pop()
			left, err := values.Top()
			if err != nil {
				return nil, err
			}
			values.Pop()
			node.Left = left
			node.Right = right

		}
		values.Push(node)
	}

	result := values.ToSlice()
	if len(result) > 0 {
		// get the root node
		return result[0], nil
	}
	return nil, fmt.Errorf("invalid postfix sequense")
}
