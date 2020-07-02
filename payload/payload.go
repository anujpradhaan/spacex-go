package payload

//PayLoad .
type PayLoad struct {
	PayLoadID       string       `json:"payload_id"`
	NoradID         []string     `json:"norad_id"`
	IsResued        bool         `json:"reused"`
	CustomerNames   []string     `json:"customers"`
	Nationality     string       `json:"nationality"`
	Manufacturer    string       `json:"manufacturer"`
	PayloadType     string       `json:"payload_type"`
	PayloadMassInKd float32      `json:"payload_mass_kg"`
	PaylodMassInLBS float32      `json:"payload_mass_lbs"`
	Orbit           string       `json:"orbit"`
	OrbitParams     OrbitDetails `json:"orbit_params"`
}

//OrbitDetails .
type OrbitDetails struct {
	ReferenceSystem   string  `json:"reference_system"`
	Regime            string  `json:"regime"`
	Longitude         float32 `json:"longitude"`
	SemiMajorAxisInKm float32 `json:"semi_major_axis_km"`
	Eccentricity      float32 `json:"eccentricity"`
	PeriapsisInKm     float32 `json:"periapsis_km"`
	ApoapsisInKm      float32 `json:"apoapsis_km"`
	InclinationDeg    float32 `json:"inclination_deg"`
	PeriodMin         float32 `json:"period_min"`
	LifespanInYears   int     `json:"lifespan_years"`
	Epoch             string  `json:"epoch"`
	MeanMotion        float32 `json:"mean_motion"`
	Raan              float32 `json:"raan"`
	ArgOfPericenter   float32 `json:"arg_of_pericenter"`
	MeanAnomaly       float32 `json:"mean_anomaly"`
}
