package grpcservice

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/mw/logging"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/mw/panic"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/mw/tracer"
	tracerlib "github.com/nika-gromova/o-architecture-patterns/project/libs/tracer"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ServiceManager struct {
	grpcServer  *grpc.Server
	httpServer  *http.Server
	adminServer *http.Server
}

type Service interface {
	RegisterGRPC(gs *grpc.Server)
	RegisterHTTP(ctx context.Context, mux *runtime.ServeMux) error
}

func New(serviceName string, service Service) (*ServiceManager, error) {
	s := &ServiceManager{}

	if err := tracerlib.InitGlobal(serviceName); err != nil {
		return nil, err
	}

	s.initGRPCServer(service)
	s.initHTTPServer(service)
	s.initAdminServer()

	return s, nil
}

func (s *ServiceManager) initGRPCServer(service Service) {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(panic.Interceptor),
		grpc.ChainUnaryInterceptor(tracer.Interceptor),
		grpc.ChainUnaryInterceptor(logging.Interceptor),
	)
	reflection.Register(server)
	service.RegisterGRPC(server)

	s.grpcServer = server
}

func (s *ServiceManager) initHTTPServer(service Service) {
	mux := runtime.NewServeMux()
	err := service.RegisterHTTP(context.Background(), mux)
	if err != nil {
		log.Fatalf("failed to register gateway: %s", err)
	}

	s.httpServer = &http.Server{
		Handler: mux,
	}
}

func (s *ServiceManager) initAdminServer() {
	mux := http.NewServeMux()

	s.adminServer = &http.Server{
		Handler: mux,
	}
}

func (s *ServiceManager) RunService(httpPort, grpcPort, adminPort int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen tcp: %v", err)
	}

	go func() {
		if err = s.grpcServer.Serve(lis); err != nil {
			log.Fatalf("grpc server failed to serve: %v", err)
		}
	}()

	s.httpServer.Addr = fmt.Sprintf(":%d", httpPort)
	go func() {
		if err = s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server failed to serve: %v", err)
		}
	}()

	s.adminServer.Addr = fmt.Sprintf(":%d", adminPort)
	go func() {
		if err = s.adminServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("admin server failed to serve: %v", err)
		}
	}()

	log.Warnf("service started, grpc port: %d, http port: %d, admin port: %d", grpcPort, httpPort, adminPort)
}
