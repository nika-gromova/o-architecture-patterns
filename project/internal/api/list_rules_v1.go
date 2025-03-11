package api

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/auth"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/pkg/rules_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListRulesV1(ctx context.Context, req *rules_v1.ListRulesV1Request) (*rules_v1.ListRulesV1Response, error) {
	uuid, err := auth.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	_, err = s.rules.ListRules(ctx, &models.Owner{
		UUID: uuid,
	})
	if err != nil {
		return nil, err
	}
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
