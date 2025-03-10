package grpcservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/mw/cors"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/mw/logging"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/mw/panic"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/mw/tracer"
	tracerlib "github.com/nika-gromova/o-architecture-patterns/project/libs/tracer"
	log "github.com/sirupsen/logrus"
	"github.com/swaggest/swgui/v5emb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ServiceManager struct {
	grpcServer                    *grpc.Server
	httpServer                    *http.Server
	adminServer                   *http.Server
	serviceName                   string
	auth                          AuthInterceptor
	httpPort, grpcPort, adminPort int
}

type Service interface {
	RegisterGRPC(gs *grpc.Server)
	RegisterHTTP(ctx context.Context, mux *runtime.ServeMux) error
}

type AuthInterceptor interface {
	InterceptorHTTP(next http.Handler) http.Handler
	InterceptorGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)
}

type opts func(s *ServiceManager)

func WithServiceName(serviceName string) opts {
	return func(s *ServiceManager) {
		s.serviceName = serviceName
	}
}

func WithAuthInterceptor(auth AuthInterceptor) opts {
	return func(s *ServiceManager) {
		s.auth = auth
	}
}

func WithPorts(httpPort, grpcPort, adminPort int) opts {
	return func(s *ServiceManager) {
		s.grpcPort = grpcPort
		s.adminPort = adminPort
		s.httpPort = httpPort
	}
}

func New(service Service, opts ...opts) (*ServiceManager, error) {
	s := &ServiceManager{}

	for _, opt := range opts {
		opt(s)
	}

	if err := tracerlib.InitGlobal(s.serviceName); err != nil {
		return nil, err
	}

	s.initGRPCServer(service)
	s.initHTTPServer(service)
	s.initAdminServer()

	return s, nil
}

func (s *ServiceManager) initGRPCServer(service Service) {
	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(panic.InterceptorGRPC),
		grpc.ChainUnaryInterceptor(tracer.InterceptorGRPC),
		grpc.ChainUnaryInterceptor(logging.InterceptorGRPC),
	}
	if s.auth != nil {
		serverOptions = append(serverOptions, grpc.UnaryInterceptor(s.auth.InterceptorGRPC))
	}

	server := grpc.NewServer(serverOptions...)
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

	var interceptors []func(http.Handler) http.Handler
	if s.auth != nil {
		interceptors = append(interceptors, s.auth.InterceptorHTTP)
	}
	interceptors = append(interceptors,
		logging.InterceptorHTTP,
		panic.InterceptorHTTP,
		cors.InterceptorHTTP,
	)

	var handler http.Handler = mux
	for _, interceptor := range interceptors {
		handler = interceptor(handler)
	}
	s.httpServer = &http.Server{
		Handler: handler,
		Addr:    fmt.Sprintf(":%d", s.httpPort),
	}
}

func (s *ServiceManager) initAdminServer() {

	mux := http.NewServeMux()

	mux.Handle("/swagger.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Загружаем оригинальный swagger.json
		var swagger map[string]interface{}
		file, err := os.ReadFile("swagger.json")
		if err != nil {
			http.Error(w, "Не удалось загрузить swagger.json", http.StatusInternalServerError)
			return
		}

		// Декодируем JSON
		if err := json.Unmarshal(file, &swagger); err != nil {
			http.Error(w, "Ошибка разбора JSON", http.StatusInternalServerError)
			return
		}

		// Подменяем servers в swagger.json
		swagger["host"] = fmt.Sprintf("localhost:%d", s.httpPort)

		// Кодируем обратно в JSON и отдаем клиенту
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(swagger); err != nil {
			http.Error(w, "Ошибка кодирования JSON", http.StatusInternalServerError)
		}
	}))

	h := v5emb.NewHandler("Rules", "/swagger.json", "/docs")

	mux.Handle("/docs/", h)

	s.adminServer = &http.Server{
		Handler: cors.InterceptorHTTP(mux),
		Addr:    fmt.Sprintf(":%d", s.adminPort),
	}
}

func (s *ServiceManager) RunService() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.grpcPort))
	if err != nil {
		log.Fatalf("failed to listen tcp: %v", err)
	}

	go func() {
		if err = s.grpcServer.Serve(lis); err != nil {
			log.Fatalf("grpc server failed to serve: %v", err)
		}
	}()

	go func() {
		if err = s.httpServer.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("http server failed to serve: %v", err)
		}
	}()

	go func() {
		if err = s.adminServer.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("admin server failed to serve: %v", err)
		}
	}()
	log.Warnf("service started, grpc port: %d, http port: %d, admin port: %d", s.grpcPort, s.httpPort, s.adminPort)
}
