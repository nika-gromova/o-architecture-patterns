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
	log "github.com/sirupsen/logrus"
	"github.com/swaggest/swgui/v5emb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type HTTPInterceptor func(next http.Handler) http.Handler

type ServiceManager struct {
	grpcServer  *grpc.Server
	httpServer  *http.Server
	adminServer *http.Server

	serviceName                   string
	httpPort, grpcPort, adminPort int
	grpcInterceptors              []grpc.UnaryServerInterceptor
	httpInterceptors              []HTTPInterceptor
	customErrorHandler            runtime.ErrorHandlerFunc
}

type Service interface {
	RegisterGRPC(gs *grpc.Server)
	RegisterHTTP(ctx context.Context, mux *runtime.ServeMux) error
}

type opts func(s *ServiceManager)

func WithServiceName(serviceName string) opts {
	return func(s *ServiceManager) {
		s.serviceName = serviceName
	}
}

func WithPorts(httpPort, grpcPort, adminPort int) opts {
	return func(s *ServiceManager) {
		s.grpcPort = grpcPort
		s.adminPort = adminPort
		s.httpPort = httpPort
	}
}

func WithGRPCInterceptors(interceptors ...grpc.UnaryServerInterceptor) opts {
	return func(s *ServiceManager) {
		for _, interceptor := range interceptors {
			s.grpcInterceptors = append(s.grpcInterceptors, interceptor)
		}
	}
}

func WithHTTPInterceptors(interceptors ...func(next http.Handler) http.Handler) opts {
	return func(s *ServiceManager) {
		for _, interceptor := range interceptors {
			s.httpInterceptors = append(s.httpInterceptors, interceptor)
		}
	}
}

func WithCustomErrorHandler(customErrorHandler runtime.ErrorHandlerFunc) opts {
	return func(s *ServiceManager) {
		s.customErrorHandler = customErrorHandler
	}
}

func New(service Service, opts ...opts) (*ServiceManager, error) {
	s := &ServiceManager{}

	for _, opt := range opts {
		opt(s)
	}

	s.initGRPCServer(service)
	s.initHTTPServer(service)
	s.initAdminServer()

	return s, nil
}

func (s *ServiceManager) initGRPCServer(service Service) {
	serverOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(panic.InterceptorGRPC),
		grpc.ChainUnaryInterceptor(logging.InterceptorGRPC),
	}
	for _, interceptor := range s.grpcInterceptors {
		serverOptions = append(serverOptions, grpc.ChainUnaryInterceptor(interceptor))
	}

	server := grpc.NewServer(serverOptions...)
	reflection.Register(server)
	service.RegisterGRPC(server)

	s.grpcServer = server
}

func (s *ServiceManager) initHTTPServer(service Service) {
	var options []runtime.ServeMuxOption
	if s.customErrorHandler != nil {
		options = append(options, runtime.WithErrorHandler(s.customErrorHandler))
	}
	mux := runtime.NewServeMux(options...)
	err := service.RegisterHTTP(context.Background(), mux)
	if err != nil {
		log.Fatalf("failed to register gateway: %s", err)
	}

	interceptors := make([]HTTPInterceptor, 0, len(s.httpInterceptors))
	for _, interceptor := range s.httpInterceptors {
		interceptors = append(interceptors, interceptor)
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
			http.Error(w, fmt.Sprintf("Не удалось загрузить swagger.json: %s", err), http.StatusInternalServerError)
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
