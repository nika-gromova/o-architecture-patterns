package rules

import (
	"context"
	"errors"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

func (s *Service) UpdateRule(ctx context.Context, rule *models.Rule) error {
	err := s.validateRule(ctx, rule)
	if err != nil {
		return err
	}
	ruleByBaseLink, err := s.storage.GetRuleByBaseLink(ctx, rule.BaseLink)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return err
	}
	if ruleByBaseLink != nil && ruleByBaseLink.Name != rule.Name {
		return fmt.Errorf("%w: base link %s already exists", models.ErrAlreadyExists, ruleByBaseLink.BaseLink)
	}

	return s.storage.UpdateRule(ctx, rule)
}
