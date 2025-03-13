package rules

import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/formula/data"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	log "github.com/sirupsen/logrus"
)

var (
	ErrNotFound        = fmt.Errorf("not found")
	ErrInvalidArgument = fmt.Errorf("invalid argument")
)

type Storage interface {
	GetRuleByBaseLink(context.Context, *models.Link) (*models.Rule, error)
}

type FormulaProcessor interface {
	Evaluate(ctx context.Context, input string, data models.Data[any]) (bool, error)
}

type RedirectStrategy interface {
	Redirect(ctx context.Context, rule *models.Rule, data models.Data[any]) *models.Link
}

type Service struct {
	baseCtx context.Context

	storage          Storage
	processor        FormulaProcessor
	redirectStrategy RedirectStrategy
}

type opts func(s *Service)

func NewService(registrar models.Registrar, storage Storage, opts ...opts) (*Service, error) {
	s := &Service{
		storage: storage,
		baseCtx: context.Background(),
	}
	s.redirectStrategy = s

	for _, opt := range opts {
		opt(s)
	}

	ctx, err := registrar.Register(s.baseCtx)
	if err != nil {
		return nil, err
	}

	s.baseCtx = ctx
	return s, nil
}

func WithRedirectStrategy(strategy RedirectStrategy) opts {
	return func(s *Service) {
		s.redirectStrategy = strategy
	}
}

func WithBaseCtx(ctx context.Context) opts {
	return func(s *Service) {
		s.baseCtx = ctx
	}
}

func (s *Service) FindRedirect(ctx context.Context, base *models.Link, request *models.Request) (*models.Link, error) {
	rule, err := s.storage.GetRuleByBaseLink(ctx, base)
	if err != nil {
		return nil, err
	}

	if rule == nil {
		return nil, fmt.Errorf("rule: %w", ErrNotFound)
	}

	requestData, err := data.NewFromRequest(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("convert request: %w, %w", ErrInvalidArgument, err)
	}

	target := s.redirectStrategy.Redirect(ctx, rule, requestData)
	if target == nil {
		return nil, fmt.Errorf("redirect for %s, %w", base.URL, ErrNotFound)
	}

	return target, nil
}

func (s *Service) Redirect(ctx context.Context, rule *models.Rule, data models.Data[any]) *models.Link {
	target := rule.DefaultRedirectTo
	for _, redirect := range rule.Redirections {
		isTrue, err := s.processor.Evaluate(ctx, redirect.Target.URL, data)
		if err != nil {
			log.Errorf("evaluate redirect url: %s, %s", redirect.Target.URL, err.Error())
			continue
		}
		if isTrue {
			target = redirect.Target
			break
		}
	}

	return target
}
