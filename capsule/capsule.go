package capsule

import (
	"spacex-go/common"
)

//CapsuleAPIURL .
const CapsuleAPIURL = "https://api.spacexdata.com/v3/capsules"

//Capsule .
type Capsule struct {
	CapsuleSerial      string           `json:"capsule_serial"`
	CapsuleID          string           `json:"capsule_id"`
	Status             string           `json:"status"`
	OriginalLaunch     string           `json:"original_launch"`
	OriginalLaunchUnix string           `json:"original_launch_unix"`
	Missions           []common.Mission `json:"missions"`
	Landings           int              `json:"landings"`
	Type               string           `json:"type"`
	Details            string           `json:"details"`
	ReuseCount         int              `json:"reuse_count"`
}