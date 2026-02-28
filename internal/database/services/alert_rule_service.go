package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"noc-monitoring-bases/internal/database/filters"
	"noc-monitoring-bases/internal/database/models"
	"noc-monitoring-bases/internal/database/repositories"
)

type AlertRuleServiceInterface interface {
	FindAll(string) ([]*models.AlertRuleModel, error)
	FindByActive() ([]*models.AlertRuleModel, error)
	FindByFilter(filters.AlertRuleFilter) ([]*models.AlertRuleWithContacts, error)
}

type AlertRuleService struct {
	queryRuleRepo repositories.AlertRulesRepositoryInterface
}

func NewQueryRuleService(repos repositories.AlertRulesRepositoryInterface) AlertRuleServiceInterface {
	return &AlertRuleService{
		queryRuleRepo: repos,
	}
}

func (s *AlertRuleService) FindAll(id string) ([]*models.AlertRuleModel, error) {
	return s.queryRuleRepo.FindAll(id)
}

func (s *AlertRuleService) FindByActive() ([]*models.AlertRuleModel, error) {

	rules, err := s.queryRuleRepo.FindByActive()
	if err != nil {
		return nil, err
	}

	return rules, nil
}

func (s *AlertRuleService) FindByFilter(filter filters.AlertRuleFilter) ([]*models.AlertRuleWithContacts, error) {
	rules, err := s.queryRuleRepo.FindByCustomFilter(s.buildMongoFilter(filter))
	if err != nil {
		return nil, err
	}

	return rules, nil
}

func (s *AlertRuleService) buildMongoFilter(filter filters.AlertRuleFilter) bson.M {
	mongoFilter := bson.M{}

	if filter.RuleName != nil {
		mongoFilter["rule_name"] = *filter.RuleName
	}
	if filter.IsActive != nil {
		mongoFilter["is_active"] = *filter.IsActive
	}
	if filter.Description != nil {
		mongoFilter["description"] = bson.M{"$regex": *filter.Description, "$options": "i"}
	}

	return mongoFilter
}
