package dragon

import (
	"spacex-go/common"
)

/*
{
		"pressurized_capsule": {
		  "payload_volume": {
			"cubic_meters": 11,
			"cubic_feet": 388
		  }
		}
	  }

*/
//Dragon .
type Dragon struct {
	ID                    int             `json:"id"`
	Name                  string          `json:"name"`
	Type                  int             `json:"type"`
	Active                bool            `json:"active"`
	CrewCapacity          int             `json:"crew_capacity"`
	SideWallAngleDeg      int             `json:"sidewall_angle_deg"`
	OrbitDurationYear     int             `json:"orbit_duration_yr"`
	DryMassInKg           int             `json:"dry_mass_kg"`
	DryMassInLb           int             `json:"dry_mass_lb"`
	FirstFlight           string          `json:"first_flight"`
	HeatShield            Shield          `json:"head_shield"`
	Thrusters             []Thruster      `json:"thrusters"`
	LaunchPayloadMass     common.Mass     `json:"launch_payload_mass"`
	LaunchPayloadVolumne  common.Volume   `json:"launch_payload_vol"`
	ReturnPayloadMass     common.Mass     `json:"return_payload_mass"`
	ReturnPayloadVolumne  common.Volume   `json:"return_payload_vol"`
	PressurizedCapsule    DragonCapsule   `json:"pressurized_capsule"`
	Trunk                 DragonTrunk     `json:"trunk"`
	HeightAndWidthOfTrunk common.Distance `json:"height_w_trunk"`
	Diameter              common.Distance `json:"diameter"`
	WikipediaLink         string          `json:"wikipedia"`
	BriefDescription      string          `json:"description"`
}

//Shield .
type Shield struct {
	Material          string  `json:"material"`
	Size              float32 `json:"size_meters"`
	Temperature       float32 `json:"temp_degrees"`
	DevelopingPartner string  `json:"dev_partner"`
}

//Thruster .
type Thruster struct {
	Type    string        `json:"type"`
	Amount  int           `json:"amount"`
	Pods    int           `json:"pods"`
	FuelOne string        `json:"fuel_1"`
	FuelTwo string        `json:"fuel_2"`
	Thrust  common.Thrust `json:"thrust"`
}

//DragonCapsule .
type DragonCapsule struct {
	PayloadVolume common.Volume `json:"paylog_volume"`
}

//DragonTrunk .
type DragonTrunk struct {
	TrunkVolumne common.Volume `json:"trunk_volume"`
	TrunkCargo   Cargo         `json:"cargo"`
}

//Cargo .
type Cargo struct {
	SolarArray           int  `json:"solar_array"`
	IsUnpressurizedCargo bool `json:"unpressurized_cargo"`
}
