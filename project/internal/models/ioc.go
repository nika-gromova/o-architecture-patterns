package models

import "context"

const (
	IoCFormulaDomain                     = "Formula."
	IoCFormulaInterpreterDomain          = "Formula.Interpreter."
	IoCFormulaInterpreterOperatorsDomain = IoCFormulaInterpreterDomain + "Operators."
	IoCFormulaInterpreterVariablesDomain = IoCFormulaInterpreterDomain + "Variables."
	IoCFormulaDataConverterDomain        = IoCFormulaDomain + "Data.Converter."
)

type Registrar interface {
	Register(ctx context.Context) (context.Context, error)
}
