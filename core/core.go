package core

import (
	"spacex-go/common"
)

//CoreAPIURL .
const CoreAPIURL = "https://api.spacexdata.com/v3/cores"

//Core .
type Core struct {
	CoreSerial         string           `json:"core_serial"`
	Block              string           `json:"block"`
	Status             string           `json:"status"`
	OriginalLaunch     string           `json:"original_launch"`
	OriginalLaunchUnix string           `json:"original_launch_unix"`
	Missions           []common.Mission `json:"missions"`
	ReuseCount         int              `json:"reuse_count"`
	RTLSAttempts       int              `json:"rtls_attempts"`
	RTLSLandings       int              `json:"rtls_landings"`
	ASDSAttempts       int              `json:"asds_attempts"`
	ASDSLandings       int              `json:"asds_landings"`
	WaterLanding       bool             `json:"water_landing"`
	Details            string           `json:"details"`
}
