package inniter

import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/api"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/config"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/parser"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/parser/shunting_yard"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/registrars"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/rules"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/rules/storage/in_memory"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/variables/datetime"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/variables/locale"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/cache"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/service"
)

type Variable interface {
	GetRegistrars() []models.Registrar
}

func InitService(ctx context.Context, cfg *config.Config) (service.Service, error) {
	variables := []Variable{
		locale.NewVariable(),
		datetime.NewVariable(),
	}

	chain := registrars.NewChain(
		&registrars.IoCFormulaOperatorsAndRegistrar{},
		&registrars.IoCFormulaOperatorsOrRegistrar{},
		&registrars.IoCFormulaOperatorsEqualRegistrar{},
		&registrars.IoCFormulaOperatorsGreaterRegistrar{},
		&registrars.IoCFormulaOperatorsLessRegistrar{},
	)
	for _, variable := range variables {
		chain.Append(variable.GetRegistrars()...)
	}

	baseCtx, err := chain.Register(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to register dependencies: %w", err)
	}

	expressionParser := parser.New(
		parser.WithParseStrategy(
			shunting_yard.New(),
		),
	)
	processor := formula.New(
		expressionParser,
		cache.New(),
	)
	rulesService := rules.NewService(
		in_memory.NewStorage(),
		processor,
		rules.WithBaseCtx(baseCtx),
	)

	return api.NewService(rulesService), nil
}
