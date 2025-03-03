package ioc

import (
	"context"
	"errors"
	"maps"
	"sync"

	"go.uber.org/multierr"
)

var (
	ErrInvalidArgumentType = errors.New("invalid argument type")
	ErrNoDependenciesFound = errors.New("no dependencies found")
)

type ScopeDependenciesKey struct{}

type Constructor func(args ...any) (any, error)

type Dependencies map[string]Constructor

type Resolver interface {
	Resolve(key string, args ...any) (any, error)
}

// Container контейнер не хранит зависимости, но облегчает доступ к ним
// каждая игра, запущенная в горутине, хранит свои зависимости в своем контексте,
// скоупы в данном случае ограничиваются контекстом.
// Важно перед запуском новой игры создать скоуп через NewScope для инициализации контекста.
type Container struct {
	mu                  sync.RWMutex
	defaultDependencies Dependencies
}

func New() *Container {
	defaultDependencies := make(Dependencies)

	defaultDependencies["IoC.Register"] = func(args ...any) (any, error) {
		if len(args) < 3 {
			return nil, errors.New("invalid number of args")
		}
		ctx, ok := args[0].(context.Context)
		if !ok {
			return nil, multierr.Append(ErrInvalidArgumentType, errors.New("args[0] is not context.Context"))
		}
		key, ok := args[1].(string)
		if !ok {
			return nil, multierr.Append(ErrInvalidArgumentType, errors.New("args[1] is not string"))
		}
		constructor, ok := args[2].(Constructor)
		if !ok {
			return nil, multierr.Append(ErrInvalidArgumentType, errors.New("args[2] is not constructor function"))
		}

		return &RegisterCommand{
			scope:       ctx,
			key:         key,
			constructor: constructor,
		}, nil
	}
	return &Container{
		defaultDependencies: defaultDependencies,
		mu:                  sync.RWMutex{},
	}
}

func (c *Container) NewScope(ctx context.Context) context.Context {
	c.mu.RLock()
	defer c.mu.RUnlock()

	defaultCopy := make(Dependencies, len(c.defaultDependencies))
	maps.Copy(defaultCopy, c.defaultDependencies)

	dependencies, ok := ctx.Value(ScopeDependenciesKey{}).(*Dependencies)
	if ok && dependencies != nil {
		for k, v := range *dependencies {
			defaultCopy[k] = v
		}
	}

	return context.WithValue(ctx, ScopeDependenciesKey{}, &defaultCopy)
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
