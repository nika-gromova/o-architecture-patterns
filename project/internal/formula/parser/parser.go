package parser

import (
	"strings"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

type Strategy func(data *models.ParsingData) (*models.ParsingNode, error)

func defaultPriorities() map[string]int {
	return map[string]int{
		models.GraterOperator: 100,
		models.LessOperator:   100,
		models.EqualOperator:  100,
		models.ANDOperator:    50,
		models.OrOperator:     30,
		"(":                   1000,
		")":                   1000,
	}
}

type Parser struct {
	operatorsPriorities map[string]int
	parseStrategy       Strategy
	tokenSeparator      string
	tokenFrame          string
}

type opts func(p *Parser)

func WithPriorities(priorities map[string]int) opts {
	return func(p *Parser) {
		p.operatorsPriorities = priorities
	}
}

func WithParseStrategy(f Strategy) opts {
	return func(p *Parser) {
		p.parseStrategy = f
	}
}

func WithTokenSeparator(s string) opts {
	return func(p *Parser) {
		p.tokenSeparator = s
	}
}

func WithTokenFrame(s string) opts {
	return func(p *Parser) {
		p.tokenFrame = s
	}
}

func New(opts ...opts) *Parser {
	p := &Parser{
		operatorsPriorities: defaultPriorities(),
		tokenSeparator:      " ",
		tokenFrame:          "'",
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *Parser) Parse(input string) (*models.ParsingNode, error) {
	tokens := p.parseToTokens(input)

	parsingData := &models.ParsingData{
		OperandsPriorities: p.operatorsPriorities,
		Tokens:             tokens,
	}
	return p.parseStrategy(parsingData)
}

func (p *Parser) parseToTokens(input string) []string {
	input = strings.TrimSpace(input)
	result := make([]string, 0, len(input))
	var (
		inToken      bool
		currentToken string
	)
	for _, ch := range input {
		val := string(ch)
		if val == p.tokenFrame {
			if !inToken {
				inToken = true
			} else {
				inToken = false
				if currentToken != "" {
					result = append(result, currentToken)
				}
				currentToken = ""
			}
		} else if val == p.tokenSeparator && !inToken {
			if currentToken != "" {
				result = append(result, currentToken)
			}
			currentToken = ""
		} else {
			currentToken += val
		}
	}
	return result
}
