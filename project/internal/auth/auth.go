package auth

import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/auth"
)

func UserFromContext(ctx context.Context) (*models.User, error) {
	claims := auth.FromContext(ctx)
	if claims == nil {
		return nil, fmt.Errorf("no claims found in context")
	}

	subj, err := claims.GetSubject()
	if err != nil {
		return nil, fmt.Errorf("failed to get user uuid from claims: %w", err)
	}
	return &models.User{
		UUID: subj,
	}, nil
}
