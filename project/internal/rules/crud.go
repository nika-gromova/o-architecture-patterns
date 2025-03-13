package rules

import (
	"context"
	"fmt"
	"strings"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

func (s *Service) CreateRule(ctx context.Context, rule *models.Rule) error {
	return s.storage.CreateRule(ctx, rule)
}

func (s *Service) DeleteRule(ctx context.Context, rule *models.Rule) error {
	return s.storage.DeleteRule(ctx, rule)
}

func (s *Service) UpdateRule(ctx context.Context, rule *models.Rule) error {
	return s.storage.UpdateRule(ctx, rule)
}

func (s *Service) ListRules(ctx context.Context, owner *models.Owner) ([]*models.Rule, error) {
	rules, err := s.storage.ListRules(ctx, owner)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, fmt.Errorf("%w: %s", ErrNotFound, err.Error())
		}
		return nil, err
	}
	return rules, nil
}

func (s *Service) GetRule(ctx context.Context, owner *models.Owner, name string) (*models.Rule, error) {
	return s.storage.GetRule(ctx, owner, name)
}
