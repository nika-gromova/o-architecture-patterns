package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/config"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/inniter"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/mw/errors"
	grpcservice "github.com/nika-gromova/o-architecture-patterns/project/libs/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.New()

	service, err := inniter.InitService(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	//authHelper := &auth.Interceptor{
	//	Authenticator: authlib.NewAuthenticator(
	//		cfg.GetSecret(config.JWTPublicKey),
	//	),
	//}
	manager, err := grpcservice.New(service,
		grpcservice.WithGRPCInterceptors(
			//authHelper.InterceptorGRPC,
			errors.InterceptorGRPC,
		),
		grpcservice.WithHTTPInterceptors(
		//authHelper.InterceptorHTTP,
		),
		grpcservice.WithCustomErrorHandler(
			errors.CustomHTTPErrorHandler,
		),
		grpcservice.WithServiceName(
			cfg.GetValue(config.AppName),
		),
		grpcservice.WithPorts(
			cfg.GetInt(config.HttpPort),
			cfg.GetInt(config.GrpcPort),
			cfg.GetInt(config.AdminPort)))
	if err != nil {
		log.Fatalf("failed to create service manager: %v", err)
	}

	manager.RunService()

	<-ctx.Done()
}
