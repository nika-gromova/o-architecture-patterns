package models

type ParsingData struct {
	OperandsPriorities map[string]int
	Tokens             []string
}

type ParsingNode struct {
	Value      string
	IsOperator bool
	Left       *ParsingNode
	Right      *ParsingNode
}

type InterpreterContext[T any] interface {
	GetValue(key string) (T, error)
}

const (
	ANDOperator    = "AND"
	OrOperator     = "OR"
	GraterOperator = ">"
	LessOperator   = "<"
	EqualOperator  = "="
)
