package api

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/pkg/rules_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListRulesV1(context.Context, *rules_v1.ListRulesV1Request) (*rules_v1.ListRulesV1Response, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
