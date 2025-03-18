package registrars

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

type Chain struct {
	registrars []models.Registrar
}

func NewChain(registrars ...models.Registrar) *Chain {
	return &Chain{
		registrars: registrars,
	}
}

func (ch *Chain) Append(registrars ...models.Registrar) {
	ch.registrars = append(ch.registrars, registrars...)
}

func (ch *Chain) Register(ctx context.Context) (context.Context, error) {
	var err error
	for _, registrar := range ch.registrars {
		ctx, err = registrar.Register(ctx)
		if err != nil {
			return ctx, err
		}
	}
	return ctx, nil
}
