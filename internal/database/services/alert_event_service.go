package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"noc-monitoring-bases/internal/database/filters"
	"noc-monitoring-bases/internal/database/models"
	"noc-monitoring-bases/internal/database/repositories"
)

type AlertEventServiceInterface interface {
	Create(alertEvent *models.AlertEventModel) (*mongo.InsertOneResult, error)
	FindByCustomFilter(filter filters.AlertEventFilter) ([]*models.AlertEventModel, error)
}

type AlertEventService struct {
	alertEventRepo repositories.AlertEventRepositoryInterface
}

func NewAlertEventService(repos repositories.AlertEventRepositoryInterface) *AlertEventService {
	return &AlertEventService{
		alertEventRepo: repos,
	}
}

func (s *AlertEventService) Create(alertEvent *models.AlertEventModel) (*mongo.InsertOneResult, error) {
	response, err := s.alertEventRepo.Create(alertEvent)

	if err != nil {
		return nil, err
	}

	return response, nil

}

func (s *AlertEventService) FindByCustomFilter(filter filters.AlertEventFilter) ([]*models.AlertEventModel, error) {
	events, err := s.alertEventRepo.FindByCustomFilter(s.buildMongoFilter(filter))
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *AlertEventService) buildMongoFilter(filter filters.AlertEventFilter) bson.M {
	mongoFilter := bson.M{}

	if filter.IdTicket != nil {
		mongoFilter["id_ticket"] = *filter.IdTicket
	}

	if filter.ResponseTicket != nil {
		mongoFilter["response_ticket"] = *filter.ResponseTicket
	}

	return mongoFilter

}
