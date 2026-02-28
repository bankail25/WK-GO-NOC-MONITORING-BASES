package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"noc-monitoring-bases/internal/database/filters"
	"noc-monitoring-bases/internal/database/models"
	"noc-monitoring-bases/internal/database/repositories"
)

type AuthPropertyServiceInterface interface {
	FindByCustomFilter(filter filters.AuthPropertyFilter) ([]*models.AuthPropertyModel, error)
}

type AuthPropertyService struct {
	Repo repositories.AuthPropertyRepositoryInterface
}

func NewAuthPropertyService(repos repositories.AuthPropertyRepositoryInterface) *AuthPropertyService {
	return &AuthPropertyService{
		Repo: repos,
	}
}

func (s *AuthPropertyService) FindByCustomFilter(filter filters.AuthPropertyFilter) ([]*models.AuthPropertyModel, error) {
	events, err := s.Repo.FindByCustomFilter(s.buildMongoFilter(filter))
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *AuthPropertyService) buildMongoFilter(filter filters.AuthPropertyFilter) bson.M {
	mongoFilter := bson.M{}

	if filter.User != nil {
		mongoFilter["user"] = *filter.User
	}

	return mongoFilter

}
