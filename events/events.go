package events

type Events struct {
	Id            int        `json:"id"`
	Title         string     `json:"title"`
	EventDateUtc  string     `json:"event_date_utc"`
	EventDateUnix string     `json:"event_date_unix"`
	FlightNumber  int        `json:"flight_number"`
	Details       string     `json:"details"`
	Links         EventLinks `json:"links"`
}

type EventLinks struct {
	Reddit    string `json:"reddit"`
	Article   string `json:"article"`
	Wikipedia string `json:"wikipedia"`
}
