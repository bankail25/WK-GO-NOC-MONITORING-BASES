package repositories

type Repositories struct {
	AuthProperty AuthPropertyRepositoryInterface
	AlertEvent   AlertEventRepositoryInterface
	AlertRules   AlertRulesRepositoryInterface
}

const (
	DbName               = "core"
	AuthProperty         = "auth_properties"
	AlertEvents          = "alert_events"
	AlertRule            = "alert_rules"
	NotificationContacts = "notification_contacts"
)

func NewRepository() *Repositories {
	return &Repositories{
		AlertRules:   NewQueryRulesRepository(),
		AuthProperty: NewAuthPropertyRepository(),
		AlertEvent:   NewAlertEventRepository(),
	}
}
