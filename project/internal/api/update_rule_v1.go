package api

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/auth"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/pkg/rules_v1"
)

func (s *Service) UpdateRuleV1(ctx context.Context, req *rules_v1.UpdateRuleV1Request) (*rules_v1.UpdateRuleV1Response, error) {
	user, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	rule, err := s.rules.GetRule(ctx, user, req.GetName())
	if err != nil {
		return nil, err
	}
	redirections := make([]*models.Redirection, 0, len(req.GetRedirects()))
	for _, redirect := range req.GetRedirects() {
		redirections = append(redirections, redirectFromProto(redirect))
	}
	rule.DefaultRedirectTo = linkFromProto(req.GetDefaultLink())
	rule.Redirections = redirections
	if err = s.rules.UpdateRule(ctx, rule); err != nil {
		return nil, err
	}

	return &rules_v1.UpdateRuleV1Response{
		Rule: ruleToProto(rule),
	}, nil
}
