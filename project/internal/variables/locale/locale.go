package locale

import (
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models/types"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/registrars"
)

const (
	name   = "locale"
	header = "Accept-Language"
)

type Variable struct {
	name   string
	header string
}

type opts func(v *Variable)

func NewVariable(opts ...opts) *Variable {
	v := &Variable{
		name:   name,
		header: header,
	}

	for _, opt := range opts {
		opt(v)
	}

	return v
}

func (v *Variable) GetRegistrars() []models.Registrar {
	return []models.Registrar{
		&registrars.IoCRequestHeaderDataConverterRegistrar{
			Header: v.header,
			Converter: func(s string) (string, any, error) {
				res, err := types.NewStringTypeFromString(s)
				return v.name, res, err
			},
		},
		&registrars.IoCFormulaStringVariableRegistrar{
			VariableName: v.name,
		},
	}
}
