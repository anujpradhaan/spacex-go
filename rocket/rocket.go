package rocket

import (
	"spacex-go/common"
)

//Rocket .
type Rocket struct {
	ID                    int                      `json:"id"`
	Active                bool                     `json:"active"`
	Stages                int                      `json:"stages"`
	Boosters              int                      `json:"boosters"`
	CostPerLaunch         float32                  `json:"cost_per_launch"`
	SuccessRatePercentage float32                  `json:"success_rate_pct"`
	FirstFlightDate       string                   `json:"first_flight"`
	Country               string                   `json:"country"`
	Company               string                   `json:"company"`
	Height                common.Distance          `json:"height"`
	Diameter              common.Distance          `json:"diameter"`
	Mass                  common.Mass              `json:"mass"`
	PayloadWeights        []common.MassWithDetails `json:"payload_weights"`
	FirstStage            RocketFirstStage         `json:"first_stage"`
	SecondStage           RocketSecondStage        `json:"second_stage"`
	Engines               RocketEngine             `json:"engines"`
	LandingLegs           RocketLandingLegs        `json:"landing_legs"`
	Wikipedia             string                   `json:"wikipedia"`
	Description           string                   `json:"description"`
	RocketID              string                   `json:"rocket_id"`
	RocketName            string                   `json:"rocket_name"`
	RocketType            string                   `json:"rocket_type"`
}

//RocketFirstStage .
type RocketFirstStage struct {
	Reusable         bool          `json:"reusable"`
	NumberOfEngines  int           `json:"engines"`
	FuelAmountInTons int           `json:"fuel_amount_tons"`
	BurnTimeInSec    int           `json:"burn_time_sec"`
	ThrustSeaLevel   common.Thrust `json:"thrust_sea_level"`
	ThrustVaccum     common.Thrust `json:"thrust_vacuum"`
}

//RocketSecondStage .
type RocketSecondStage struct {
	NumberOfEngines  int                      `json:"engines"`
	FuelAmountInTons int                      `json:"fuel_amount_tons"`
	BurnTimeInSec    int                      `json:"burn_time_sec"`
	Thrust           common.Thrust            `json:"thrust"`
	Payloads         RocketSecondStagePayload `json:"payloads"`
}

//RocketSecondStagePayload .
type RocketSecondStagePayload struct {
	Option1          string                     `json:"option_1"`
	Option2          string                     `json:"option_2"`
	CompositeFairing map[string]common.Distance `json:"composite_fairing"`
}

//RocketEngine .
type RocketEngine struct {
	Number         int           `json:"number"`
	Type           string        `json:"type"`
	Version        string        `json:"version"`
	Layout         string        `json:"layout"`
	EngineLossMax  int           `json:"engine_loss_max"`
	Propellant1    string        `json:"propellant_1"`
	Propellant2    string        `json:"propellant_2"`
	ThrustSeaLevel common.Thrust `json:"thrust_sea_level"`
	ThrustVaccum   common.Thrust `json:"thrust_vacuum"`
	ThrustToWeight float32       `json:"thrust_to_weight:"`
}

//RocketLandingLegs .
type RocketLandingLegs struct {
	Number   int    `json:"number"`
	Material string `json:"material"`
}
