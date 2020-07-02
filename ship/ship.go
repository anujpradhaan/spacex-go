package ship

import (
	"spacex-go/common"
)

//Ship .
type Ship struct {
	ShipID             string          `json:"ship_id"`
	ShipName           string          `json:"ship_name"`
	ShipModel          string          `json:"ship_model"`
	ShipType           string          `json:"ship_type"`
	Roles              []string        `json:"roles"`
	Active             bool            `json:"active"`
	IMO                int             `json:"imo"`
	MMSI               int             `json:"mmsi"`
	ABS                int             `json:"abs"`
	Class              int             `json:"class"`
	WeightInLbs        float32         `json:"weight_lbs"`
	WeightInKg         float32         `json:"weight_kg"`
	YearBuilt          int             `json:"year_built"`
	HomePort           string          `json:"home_port"`
	Status             string          `json:"status"`
	SpeedKN            int             `json:"speed_kn"`
	CourseDeg          string          `json:"course_deg"`
	Position           common.Location `json:"position"`
	SuccessfulLandings int             `json:"successful_landings"`
	AttemptedLandings  int             `json:"attempted_landings"`
	Missions           []ShipMission   `json:"missions"`
	URL                string          `json:"url"`
	ImageURL           string          `json:"image"`
}

//ShipMission .
type ShipMission struct {
	Name   string `json:"name"`
	Flight int    `json:"flight"`
}
