package rules

import (
	"context"
	"fmt"
	"strings"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

func (s *Service) ListRules(ctx context.Context, owner *models.User) ([]*models.Rule, error) {
	rules, err := s.storage.ListRules(ctx, owner)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, fmt.Errorf("%w: %s", ErrNotFound, err.Error())
		}
		return nil, err
	}
	return rules, nil
}

func (s *Service) GetRule(ctx context.Context, owner *models.User, name string) (*models.Rule, error) {
	return s.storage.GetRule(ctx, owner, name)
}
