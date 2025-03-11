package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/api"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/config"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/mw/errors"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/rules"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/rules/storage/in_memory"
	auth_lib "github.com/nika-gromova/o-architecture-patterns/project/libs/auth"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/mw/auth"
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

	rulesService := rules.NewService(in_memory.NewStorage())

	service := api.NewService(rulesService)
	authService := &auth.Interceptor{
		Authenticator: auth_lib.NewAuthenticator(cfg.GetSecret(config.JWTPublicKey)),
	}
	manager, err := grpcservice.New(service,
		grpcservice.WithGRPCInterceptors(
			authService.InterceptorGRPC,
			errors.InterceptorGRPC),
		grpcservice.WithHTTPInterceptors(
			authService.InterceptorHTTP),
		grpcservice.WithCustomErrorHandler(errors.CustomHTTPErrorHandler),
		grpcservice.WithServiceName(os.Getenv("APP_NAME")),
		grpcservice.WithPorts(httpPort, grpcPort, adminPort))
	if err != nil {
		log.Fatalf("failed to create service manager: %v", err)
	}

	manager.RunService()

	<-ctx.Done()
}
