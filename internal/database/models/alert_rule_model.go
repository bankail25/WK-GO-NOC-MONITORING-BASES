package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AlertRuleModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RuleName    string             `bson:"rule_name" json:"rule_name"`
	QueryFilter string             `bson:"query_filter" json:"query_filter"`
	IsActive    bool               `bson:"is_active" json:"is_active"`
	Description string             `bson:"description" json:"description"`
}

type AlertRuleWithContacts struct {
	AlertRuleModel `bson:",inline"`
}
