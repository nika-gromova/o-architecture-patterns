package rules

import (
	"context"
	"fmt"
	"strings"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
)

var (
	ErrNotFound        = fmt.Errorf("not found")
	ErrInvalidArgument = fmt.Errorf("invalid argument")
)

type Storage interface {
	CreateRule(context.Context, *models.Rule) error
	DeleteRule(context.Context, *models.Rule) error
	UpdateRule(context.Context, *models.Rule) error
	ListRules(context.Context, *models.Owner) ([]*models.Rule, error)
	GetRule(context.Context, *models.Owner, string) (*models.Rule, error)
}

type Service struct {
	storage Storage
}

type opts func(s *Service)

func NewService(storage Storage, opts ...opts) *Service {
	s := &Service{
		storage: storage,
	}

	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Service) CreateRule(ctx context.Context, rule *models.Rule) error {
	return s.storage.CreateRule(ctx, rule)
}

func (s *Service) DeleteRule(ctx context.Context, rule *models.Rule) error {
	return s.storage.DeleteRule(ctx, rule)
}

func (s *Service) UpdateRule(ctx context.Context, rule *models.Rule) error {
	return s.storage.UpdateRule(ctx, rule)
}

func (s *Service) ListRules(ctx context.Context, owner *models.Owner) ([]*models.Rule, error) {
	rules, err := s.storage.ListRules(ctx, owner)
	if err != nil {
		if strings.Contains(err.Error(), ErrNotFound.Error()) {
			return nil, fmt.Errorf("%w: %s", ErrNotFound, err.Error())
		}
		return nil, err
	}
	return rules, nil
}

func (s *Service) GetRule(ctx context.Context, owner *models.Owner, name string) (*models.Rule, error) {
	return s.storage.GetRule(ctx, owner, name)
}
