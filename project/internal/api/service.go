package api

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nika-gromova/o-architecture-patterns/project/pkg/rules_v1"
	"google.golang.org/grpc"
)

type Service struct {
	rules_v1.UnimplementedRulesServer
}

func (s *Service) RegisterGRPC(grpcServer *grpc.Server) {
	rules_v1.RegisterRulesServer(grpcServer, s)
}

func (s *Service) RegisterHTTP(ctx context.Context, mux *runtime.ServeMux) error {
	return rules_v1.RegisterRulesHandlerServer(ctx, mux, s)
}
