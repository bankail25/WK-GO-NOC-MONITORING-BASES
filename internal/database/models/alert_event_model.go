package models

import (
	"time"
)

type SuppressionMetadata struct {
	IdTicket       string    `bson:"id_ticket" json:"id_ticket"`
	ResponseTicket string    `bson:"response_ticket" json:"response_ticket"`
	DynamicText    string    `bson:"dynamic_text" json:"dynamic_text"`
	CreatedAt      time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at" json:"updated_at"`
}

type AlertEventModel struct {
	SuppressionMetadata `bson:",inline"`
}
