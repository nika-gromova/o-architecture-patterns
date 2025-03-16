package rules

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

func (s *Service) DeleteRule(ctx context.Context, rule *models.Rule) error {
	return s.storage.DeleteRule(ctx, rule)
}
