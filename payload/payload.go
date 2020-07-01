package payload

//PayLoad .
type PayLoad struct {
	ID              string
	NoradId         []string
	IsResued        bool
	CustomerNames   []string
	Nationality     string
	Manufacturer    string
	PayloadType     string
	PayloadMassInKd float32
	PaylodMassInLBS float32
	Orbit           string
	OrbitParams     OrbitDetails
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
}
