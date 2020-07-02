package roadster

//Roadster .
type Roadster struct {
	Name                      string  `json"name"`
	LaunchDateInUTC           string  `json"launch_date_utc"`
	LaunchDateInUnixTimeStamp string  `json"launch_date_unix"`
	LaunchMassInKg            float32 `json"launch_mass_kg"`
	LaunchMassInLbs           float32 `json"launch_mass_lbs"`
	NoradID                   int     `json"norad_id"`
	EpochJD                   float32 `json"epoch_jd"`
	OrbitType                 string  `json"orbit_type"`
	ApoapsisAU                float32 `json"apoapsis_au"`
	PeriapsisAU               float32 `json"periapsis_au"`
	SemiMajorAxisAU           float32 `json"semi_major_axis_au"`
	Eccentricity              float32 `json"eccentricity"`
	Inclination               float32 `json"inclination"`
	Longitude                 float32 `json"longitude"`
	PeriapsisARG              float32 `json"periapsis_arg"`
	PeriodInDays              float32 `json"period_days"`
	SpeedInkph                float32 `json"speed_kph"`
	SpeedInmph                float32 `json"speed_mph"`
	EarthDistanceInKM         float32 `json"earth_distance_km"`
	EarthDistanceInMiles      float32 `json"earth_distance_mi"`
	MarsDistanceInKm          float32 `json"mars_distance_km"`
	MarsDistanceInMiles       float32 `json"mars_distance_mi"`
	Wikipedia                 string  `json"wikipedia"`
	Details                   string  `json"details"`
}
