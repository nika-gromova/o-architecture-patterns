package rules

import (
	"context"
	"fmt"

	data "github.com/nika-gromova/o-architecture-patterns/project/internal/data/request"
	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/nika-gromova/o-architecture-patterns/project/libs/ioc"
	log "github.com/sirupsen/logrus"
)

var (
	ErrNotFound        = fmt.Errorf("not found")
	ErrInvalidArgument = fmt.Errorf("invalid argument")
)

type Storage interface {
	CreateRule(context.Context, *models.Rule) error
	DeleteRule(context.Context, *models.Rule) error
	UpdateRule(context.Context, *models.Rule) error
	ListRules(context.Context, *models.User) ([]*models.Rule, error)
	GetRule(context.Context, *models.User, string) (*models.Rule, error)

	GetRuleByBaseLink(context.Context, *models.Link) (*models.Rule, error)
}

type FormulaProcessor interface {
	Evaluate(ctx context.Context, input *models.Formula, data models.Data[any]) (bool, error)
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

func NewService(storage Storage, processor FormulaProcessor, opts ...opts) (*Service, error) {
	s := &Service{
		storage:   storage,
		processor: processor,
		baseCtx:   context.Background(),
	}
	s.redirectStrategy = s

	for _, opt := range opts {
		opt(s)
	}

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
	ctx = ioc.NewFromParent(s.baseCtx, ctx)

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
		isTrue, err := s.processor.Evaluate(ctx, redirect.Formula, data)
		if err != nil {
			log.Errorf("evaluate formula: %s, %s", redirect.Formula.Expression, err.Error())
			continue
		}
		if isTrue {
			target = redirect.Target
			break
		}
	}

	return target
}
