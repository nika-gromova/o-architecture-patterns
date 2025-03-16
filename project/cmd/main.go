package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/api"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/config"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/data/request"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/parser"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/parser/shunting_yard"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/mw/errors"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/rules"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/rules/storage/in_memory"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/cache"
	grpcservice "github.com/nika-gromova/o-architecture-patterns/project/libs/service"
	log "github.com/sirupsen/logrus"
)

const (
	grpcPort  = 50051
	httpPort  = 8080
	adminPort = 8081
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.New()

	converters := request.GetInitConverters()

	var registrar models.Registrar
	registrar = &formula.IoCFormulaOperatorsRegistrar{}
	variablesToHeaders := cfg.GetHeaderVariables()
	for _, variable := range variablesToHeaders {
		converter, found := converters[variable.Name]
		if found {
			var headerRegistrar models.Registrar = &request.IoCRequestHeaderDataConverterRegistrar{
				Header:    variable.Header,
				Converter: converter,
			}
			headerRegistrar.SetNext(registrar)
			registrar = headerRegistrar
		}
		if variable.Type == "string" {
			var stringRegistrar models.Registrar = &formula.IoCFormulaStringVariableRegistrar{
				VariableName: variable.Name,
			}
			stringRegistrar.SetNext(registrar)
			registrar = stringRegistrar
		}
		if variable.Type == "time" {
			var timeRegistrar models.Registrar = &formula.IoCFormulaDateTimeVariableRegistrar{
				VariableName: variable.Name,
			}
			timeRegistrar.SetNext(registrar)
			registrar = timeRegistrar
		}
	}

	baseCtx, err := registrar.Register(context.Background())
	if err != nil {
		log.Fatalf("failed to register dependencies: %v", err)
	}

	expressionParser := parser.New(
		parser.WithParseStrategy(shunting_yard.New()),
	)
	processor := formula.New(expressionParser, cfg, cache.New())
	rulesService, err := rules.NewService(
		in_memory.NewStorage(),
		processor,
		rules.WithBaseCtx(baseCtx),
	)

	service := api.NewService(rulesService)
	//authService := &auth.Interceptor{
	//	Authenticator: auth_lib.NewAuthenticator(cfg.GetSecret(config.JWTPublicKey)),
	//}
	manager, err := grpcservice.New(service,
		grpcservice.WithGRPCInterceptors(
			//authService.InterceptorGRPC,
			errors.InterceptorGRPC),
		//grpcservice.WithHTTPInterceptors(
		//	authService.InterceptorHTTP),
		grpcservice.WithCustomErrorHandler(errors.CustomHTTPErrorHandler),
		grpcservice.WithServiceName(os.Getenv("APP_NAME")),
		grpcservice.WithPorts(httpPort, grpcPort, adminPort))
	if err != nil {
		log.Fatalf("failed to create service manager: %v", err)
	}

	manager.RunService()

	<-ctx.Done()
}
