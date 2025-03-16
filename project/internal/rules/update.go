package rules

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

func (s *Service) UpdateRule(ctx context.Context, rule *models.Rule) error {
	return s.storage.UpdateRule(ctx, rule)
}
