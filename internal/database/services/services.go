package services

import (
	"noc-monitoring-bases/internal/database/repositories"
)

type Services struct {
	AlertRule    AlertRuleServiceInterface
	AlertEvent   AlertEventServiceInterface
	AuthProperty AuthPropertyServiceInterface
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		AlertRule:    NewQueryRuleService(repos.AlertRules),
		AlertEvent:   NewAlertEventService(repos.AlertEvent),
		AuthProperty: NewAuthPropertyService(repos.AuthProperty),
	}
}
