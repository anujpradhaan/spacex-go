package landing

import (
	"spacex-go/common"
)

//Landing .
type Landing struct {
	ID                 string          `json:"id"`
	FullName           string          `json:"full_name"`
	Status             string          `json:"status"`
	LandingSite        common.Location `json:"location"`
	LandingType        string          `json:"landing_type"`
	AttemptedLandings  int             `json:"attempted_landings"`
	SuccessfulLandings int             `json:"successful_landings"`
	WikipediaLink      string          `json:"wikipedia"`
	Details            string          `json:"details"`
}
