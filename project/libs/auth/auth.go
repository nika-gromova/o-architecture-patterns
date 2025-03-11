package auth

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

const claims = "auth-claims"

type Authenticator struct {
	secretKey string
}

func NewAuthenticator(secretKey string) *Authenticator {
	return &Authenticator{
		secretKey: secretKey,
	}
}

func (a *Authenticator) Authenticate(token string) (*jwt.Token, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(a.secretKey))
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA key: %w", err)
	}

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return the public key that is used to validate the token.
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !t.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return t, nil
}

func ToContext(ctx context.Context, token *jwt.Token) context.Context {
	return context.WithValue(ctx, claims, token.Claims)
}

func FromContext(ctx context.Context) jwt.Claims {
	var result jwt.Claims
	result, ok := ctx.Value(claims).(jwt.Claims)
	if !ok {
		log.Errorf("failed to get claims from context")
	}
	return result
}
