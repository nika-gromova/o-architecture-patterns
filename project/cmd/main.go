package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/api"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/mw/auth"
	grpcservice "github.com/nika-gromova/o-architecture-patterns/project/libs/service"
	log "github.com/sirupsen/logrus"
)

const (
	grpcPort  = 50051
	httpPort  = 8080
	adminPort = 8081

	appName = "rules"
)

func main() {
	service := &api.Service{}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	authService := &auth.Interceptor{}
	manager, err := grpcservice.New(service,
		grpcservice.WithAuthInterceptor(authService),
		grpcservice.WithServiceName(appName),
		grpcservice.WithPorts(httpPort, grpcPort, adminPort))
	if err != nil {
		log.Fatalf("failed to create service manager: %v", err)
	}

	manager.RunService()

	<-ctx.Done()
}
