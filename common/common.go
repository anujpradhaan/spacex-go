package common

//Mission .
type Mission struct {
	Name   string `json:"name"`
	Flight string `json:"flight"`
}

//Mass .
type Mass struct {
	Kg float32 `json:"kg"`
	LB float32 `json:"lb"`
}

//Volume .
type Volume struct {
	CubicMeter int `json:"cubic_meters"`
	CubicFeet  int `json:"cubic_feet"`
}

//Distance .
type Distance struct {
	Meters int `json:"meterts"`
	Feet   int `json:"feet"`
}

//Thrust .
type Thrust struct {
	KN  float32 `json:"kN"`
	LBF float32 `json:"lbf"`
}

//Address .
type Address struct {
	CompanyAddress string `json:"address"`
	City           string `json:"city"`
	State          string `json:"state"`
}

//Location .
type Location struct {
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
