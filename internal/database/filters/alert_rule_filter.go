package filters

type AlertRuleFilter struct {
	RuleName    *string `json:"rule_name,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
	Description *string `json:"description,omitempty"`
}
