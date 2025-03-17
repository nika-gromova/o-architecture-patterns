package rules

import (
	"context"
	"errors"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

func (s *Service) CreateRule(ctx context.Context, rule *models.Rule) error {
	err := s.validateRule(ctx, rule)
	if err != nil {
		return err
	}
	ruleByBaseLink, err := s.storage.GetRuleByBaseLink(ctx, rule.BaseLink)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return err
	}
	if ruleByBaseLink != nil {
		return fmt.Errorf("%w: base link %s already exists", models.ErrAlreadyExists, ruleByBaseLink.BaseLink)
	}
	err = s.storage.CreateRule(ctx, rule)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) validateRule(_ context.Context, rule *models.Rule) error {
	if rule.BaseLink == nil {
		return fmt.Errorf("%w: empty base link", models.ErrInvalidArgument)
	}
	if rule.DefaultRedirectTo == nil && len(rule.Redirections) == 0 {
		return fmt.Errorf("%w: empty target links list", models.ErrInvalidArgument)
	}
	if rule.Owner == nil {
		return fmt.Errorf("%w: empty owner", models.ErrInvalidArgument)
	}
	return nil
}
