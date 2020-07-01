package landing

import (
	"spacex-go/common"
)

/*
{
  "id": "LZ-4",
  "full_name": "Landing Zone 4",
  "status": "active",
  "location": {
    "name": "Vandenberg Air Force Base",
    "region": "California",
    "latitude": 34.632989,
    "longitude": -120.615167
  },
  "landing_type": "RTLS",
  "attempted_landings": 1,
  "successful_landings": 1,
  "wikipedia": "https://en.wikipedia.org/wiki/Vandenberg_AFB_Space_Launch_Complex_4#LZ-4_landing_history",
  "details": "SpaceX's west coast landing pad. The pad is adjacent to SLC-4E, SpaceX's west coast launch site. The pad was under construction for about a year starting in 2016. After concerns with seal mating season, this pad was first used for the SAOCOM 1A mission in October 2018. Officially referred to as LZ-4 in FCC filings."
}
*/
type Landing struct {
	Id                 string          `json:"id"`
	FullName           string          `json:"full_name"`
	Status             string          `json:"status"`
	LandingSite        common.Location `json:"location"`
	LandingType        string          `json:"landing_type"`
	AttemptedLandings  int             `json:"attempted_landings"`
	SuccessfulLandings int             `json:"successful_landings"`
	WikipediaLink      string          `json:"wikipedia"`
	Details            string          `json:"details"`
}
