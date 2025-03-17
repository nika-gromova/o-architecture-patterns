package api

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/auth"
	"github.com/nika-gromova/o-architecture-patterns/project/pkg/rules_v1"
)

func (s *Service) ListRulesV1(ctx context.Context, _ *rules_v1.ListRulesV1Request) (*rules_v1.ListRulesV1Response, error) {
	user, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	rules, err := s.rules.ListRules(ctx, user)
	if err != nil {
		return nil, err
	}
	resp := &rules_v1.ListRulesV1Response{
		Rules: make([]*rules_v1.RuleV1, 0, len(rules)),
	}
	for _, rule := range rules {
		resp.Rules = append(resp.Rules, ruleToProto(rule))
	}
	return resp, nil
}
