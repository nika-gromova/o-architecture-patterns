package api

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/auth"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/pkg/rules_v1"
)

func (s *Service) ListRulesV1(ctx context.Context, _ *rules_v1.ListRulesV1Request) (*rules_v1.ListRulesV1Response, error) {
	uuid, err := auth.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	rules, err := s.rules.ListRules(ctx, &models.Owner{
		UUID: uuid,
	})
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

func ruleToProto(rule *models.Rule) *rules_v1.RuleV1 {
	result := &rules_v1.RuleV1{
		Name:     rule.Name,
		BaseLink: linkToProto(rule.BaseLink),
	}
	if rule.DefaultRedirectTo != nil {
		result.DefaultLink = linkToProto(rule.DefaultRedirectTo)
	}
	for _, r := range rule.Redirections {
		result.Redirects = append(result.Redirects, redirectToProto(r))
	}

	return result
}

func redirectToProto(redirect *models.Redirection) *rules_v1.RedirectV1 {
	return &rules_v1.RedirectV1{
		Formula: &rules_v1.FormulaV1{
			Expression: redirect.Formula.Expression,
		},
		TargetLink: linkToProto(redirect.Target),
	}
}

func linkToProto(link *models.Link) *rules_v1.LinkV1 {
	return &rules_v1.LinkV1{
		Url: link.URL,
	}
}
