package ioc

import (
	"context"
	"errors"
)

var (
	ErrNoDependenciesFound = errors.New("no dependencies found")
)

type ScopeDependenciesKey struct{}

type Constructor func(args ...any) (any, error)

type Dependencies map[string]Constructor

type Resolver interface {
	Resolve(key string, args ...any) (any, error)
}

func NewScope(ctx context.Context) context.Context {
	dependenciesCopy := make(Dependencies)

	dependencies, ok := ctx.Value(ScopeDependenciesKey{}).(*Dependencies)
	if ok && dependencies != nil {
		for k, v := range *dependencies {
			dependenciesCopy[k] = v
		}
	}

	return context.WithValue(ctx, ScopeDependenciesKey{}, &dependenciesCopy)
}

func Resolve(ctx context.Context, key string, args ...any) (any, error) {
	dependencies, ok := ctx.Value(ScopeDependenciesKey{}).(*Dependencies)
	if !ok || dependencies == nil {
		return nil, ErrNoDependenciesFound
	}
	resolver := NewDefaultResolver(*dependencies)
	return resolver.Resolve(key, args...)
}

func Register(ctx context.Context, key string, constructor Constructor) error {
	cmd := &RegisterCommand{
		scope:       ctx,
		key:         key,
		constructor: constructor,
	}
	return cmd.Execute(ctx)
}
