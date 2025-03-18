package rules

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

func (s *Service) ListRules(ctx context.Context, owner *models.User) ([]*models.Rule, error) {
	rules, err := s.storage.ListRules(ctx, owner)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func (s *Service) GetRule(ctx context.Context, owner *models.User, name string) (*models.Rule, error) {
	rule, err := s.storage.GetRule(ctx, owner, name)
	if err != nil {
		return nil, err
	}
	return rule, nil
}
