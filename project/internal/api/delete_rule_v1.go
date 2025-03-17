package api

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/auth"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/pkg/rules_v1"
)

func (s *Service) DeleteRuleV1(ctx context.Context, req *rules_v1.DeleteRuleV1Request) (*rules_v1.DeleteRuleV1Response, error) {
	user, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if err = s.rules.DeleteRule(ctx, &models.Rule{
		Name:  req.GetName(),
		Owner: user,
	}); err != nil {
		return nil, err
	}

	return &rules_v1.DeleteRuleV1Response{}, nil
}
