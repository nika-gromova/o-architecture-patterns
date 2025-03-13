package models

import "context"

const (
	IoCFormulaDomain                     = "Formula."
	IoCFormulaInterpreterDomain          = "Formula.Interpreter."
	IoCFormulaInterpreterOperatorsDomain = IoCFormulaInterpreterDomain + "Operators."
	IoCFormulaInterpreterVariablesDomain = IoCFormulaInterpreterDomain + "Variables."
	IoCFormulaDataConverterHeadersDomain = IoCFormulaDomain + "Data.Converter.Headers."
)

type Registrar interface {
	Register(ctx context.Context) (context.Context, error)
	// Append(func(context.Context) (context.Context, error))
}
