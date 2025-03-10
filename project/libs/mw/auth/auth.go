package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Authenticator interface {
	Authenticate(token string) (*jwt.Token, error)
}

type Interceptor struct {
	Authenticator Authenticator
}

func (i *Interceptor) InterceptorHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")

		token := strings.Split(header, "Bearer")
		if len(token) != 2 {
			http.Error(w, "invalid authorization header", http.StatusUnauthorized)
			return
		}
		ctx, err := i.auth(r.Context(), strings.TrimSpace(token[1]))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (i *Interceptor) InterceptorGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "no metadata found")
	}
	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Error(codes.Unauthenticated, "no authorization found")
	}
	token := values[0]

	ctx, err = i.auth(ctx, token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return handler(ctx, req)
}

func (i *Interceptor) auth(ctx context.Context, token string) (context.Context, error) {
	claims, err := i.Authenticator.Authenticate(token)
	if err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}

	return auth.ToContext(ctx, claims), nil
}
