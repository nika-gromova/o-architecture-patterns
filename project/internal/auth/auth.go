package auth

import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/libs/auth"
)

func UserIDFromContext(ctx context.Context) (string, error) {
	claims := auth.FromContext(ctx)
	if claims == nil {
		return "", fmt.Errorf("no claims found in context")
	}

	return claims.GetSubject()
}
