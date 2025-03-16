package in_memory

import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/project/internal/models"
	"github.com/samber/lo"
)

type Storage struct {
	rules map[string]map[string]*models.Rule
}

func NewStorage() *Storage {
	return &Storage{
		rules: make(map[string]map[string]*models.Rule),
	}
}

func (s *Storage) CreateRule(_ context.Context, rule *models.Rule) error {
	rules, found := s.rules[rule.Owner.UUID]
	if !found {
		s.rules[rule.Owner.UUID] = make(map[string]*models.Rule)
	}
	if _, exists := rules[rule.Name]; exists {
		return fmt.Errorf("rule %s already exists for user %s", rule.Name, rule.Owner.UUID)
	}

	s.rules[rule.Owner.UUID][rule.Name] = rule
	return nil
}

func (s *Storage) DeleteRule(_ context.Context, rule *models.Rule) error {
	_, err := s.getRule(rule.Owner, rule.Name)
	if err != nil {
		return err
	}
	delete(s.rules[rule.Owner.UUID], rule.Name)
	return nil
}

func (s *Storage) UpdateRule(_ context.Context, rule *models.Rule) error {
	_, err := s.getRule(rule.Owner, rule.Name)
	if err != nil {
		return err
	}
	s.rules[rule.Owner.UUID][rule.Name] = rule
	return nil
}

func (s *Storage) ListRules(_ context.Context, owner *models.User) ([]*models.Rule, error) {
	rules, err := s.getRules(owner)
	if err != nil {
		return nil, err
	}
	return lo.Values(rules), nil
}

func (s *Storage) GetRule(_ context.Context, owner *models.User, name string) (*models.Rule, error) {
	return s.getRule(owner, name)
}

func (s *Storage) getRule(owner *models.User, name string) (*models.Rule, error) {
	rules, err := s.getRules(owner)
	if err != nil {
		return nil, err
	}
	rule, exists := rules[name]
	if !exists {
		return nil, fmt.Errorf("rule %s not found for user %s", rule.Name, rule.Owner.UUID)
	}
	return rule, nil
}

func (s *Storage) getRules(owner *models.User) (map[string]*models.Rule, error) {
	rules, found := s.rules[owner.UUID]
	if !found {
		return nil, fmt.Errorf("rules for owner %s not found", owner.UUID)
	}
	return rules, nil
}

func (s *Storage) GetRuleByBaseLink(_ context.Context, base *models.Link) (*models.Rule, error) {
	for _, rules := range s.rules {
		for _, rule := range rules {
			if rule.BaseLink.Equals(base) {
				return rule, nil
			}
		}
	}
	return nil, fmt.Errorf("rule with base URL %s not found", base.URL)
}
