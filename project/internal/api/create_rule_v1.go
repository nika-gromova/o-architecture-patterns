package api

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/auth"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/pkg/rules_v1"
)

func (s *Service) CreateRuleV1(ctx context.Context, req *rules_v1.CreateRuleV1Request) (*rules_v1.CreateRuleV1Response, error) {
	user, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	redirections := make([]*models.Redirection, 0, len(req.Rule.GetRedirects()))
	for _, redirect := range req.Rule.GetRedirects() {
		redirections = append(redirections, redirectFromProto(redirect))
	}
	if err = s.rules.CreateRule(ctx, &models.Rule{
		Name:              req.Rule.GetName(),
		Owner:             user,
		BaseLink:          linkFromProto(req.Rule.GetBaseLink()),
		DefaultRedirectTo: linkFromProto(req.Rule.GetDefaultLink()),
		Redirections:      redirections,
	}); err != nil {
		return nil, err
	}

	return &rules_v1.CreateRuleV1Response{}, nil
}

func linkFromProto(link *rules_v1.LinkV1) *models.Link {
	return &models.Link{
		URL: link.GetUrl(),
	}
}

func redirectFromProto(redirect *rules_v1.RedirectV1) *models.Redirection {
	return &models.Redirection{
		Formula: formulaFromProto(redirect.GetFormula()),
		Target:  linkFromProto(redirect.TargetLink),
	}
}

func formulaFromProto(formula *rules_v1.FormulaV1) *models.Formula {
	return &models.Formula{
		Expression: formula.GetExpression(),
	}
}
