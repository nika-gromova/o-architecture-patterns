package ioc

import (
	"fmt"
)

type DefaultResolver struct {
	dependencies Dependencies
}

func (r *DefaultResolver) Resolve(key string, args ...any) (any, error) {
	constructor, found := r.dependencies[key]
	if !found {
		return nil, fmt.Errorf("no such dependency with key %s", key)
	}
	return constructor(args...)
}

func NewDefaultResolver(dependencies Dependencies) *DefaultResolver {
	return &DefaultResolver{dependencies}
}
