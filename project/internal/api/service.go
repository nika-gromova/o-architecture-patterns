package api

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/pkg/rules_v1"
	"google.golang.org/grpc"
)

type RulesService interface {
	CreateRule(ctx context.Context, rule *models.Rule) error
	DeleteRule(ctx context.Context, rule *models.Rule) error
	UpdateRule(ctx context.Context, rule *models.Rule) error
	ListRules(ctx context.Context, owner *models.Owner) ([]*models.Rule, error)
	GetRule(ctx context.Context, owner *models.Owner, name string) (*models.Rule, error)
}

type Service struct {
	rules_v1.UnimplementedRulesServer

	rules RulesService
}

func NewService(rs RulesService) *Service {
	return &Service{
		rules: rs,
	}
}

func (s *Service) RegisterGRPC(grpcServer *grpc.Server) {
	rules_v1.RegisterRulesServer(grpcServer, s)
}

func (s *Service) RegisterHTTP(ctx context.Context, mux *runtime.ServeMux) error {
	return rules_v1.RegisterRulesHandlerServer(ctx, mux, s)
}
