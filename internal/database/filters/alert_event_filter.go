package filters

type AlertEventFilter struct {
	IdTicket       *string `json:"id_ticket"`
	ResponseTicket *string `json:"response_ticket"`
}

type AlertEventFilterList struct {
	AlertEventFilter
}
