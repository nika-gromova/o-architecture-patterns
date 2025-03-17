package registrars

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/ioc"
)

type IoCRequestHeaderDataConverterRegistrar struct {
	Header    string
	Converter func(string) (string, any, error)
}

func (r *IoCRequestHeaderDataConverterRegistrar) Register(oldCtx context.Context) (context.Context, error) {
	ctx := ioc.NewScope(oldCtx)

	err := ioc.Register(ctx, models.IoCFormulaDataConverterHeadersDomain+r.Header, func(args ...any) (any, error) {
		return r.Converter, nil
	})

	if err != nil {
		return nil, err
	}

	return ctx, nil
}
