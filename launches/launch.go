package launch

import (
	"spacex-go/common"
	"spacex-go/payload"
)

//Launch .
type Launch struct {
	FlightNumber          int             `json:"flight_number"`
	MissionName           string          `json:"mission_name"`
	MissionID             []string        `json:"mission_id"`
	LaunchYear            string          `json:"launch_year"`
	LaunchDateUnix        string          `json:"launch_date_unix"`
	LaunchDateUTC         string          `json:"launch_date_utc"`
	LaunchDateLocal       string          `json:"launch_date_local"`
	IsTentative           bool            `json:"is_tentative"`
	TentativeMaxPrecision string          `json:"tentative_max_precision"`
	TBD                   bool            `json:"tbd"`
	LaunchWindow          int             `json:"launch_window"`
	Rocket                LaunchRocket    `json:"rocket"`
	Fairings              LaunchFairing   `json:"fairings"`
	Ships                 []string        `json:"ships"`
	Telemetry             LaunchTelemetry `json:"telemetry"`
	LaunchingSite         LaunchSite      `json:"launch_site"`
	LaunchSuccess         bool            `json:"launch_success"`
	Details               string          `json:"details"`
	Upcoming              bool            `json:"upcoming"`
	StaticFireDateUtc     string          `json:"static_fire_date_utc"`
	StaticFireDateUnix    string          `json:"static_fire_date_unix"`
	Links                 LaunchLinks     `json:"links"`
	Timeline              LaunchTimeline  `json:"timeline"`
}

//LaunchRocket .
type LaunchRocket struct {
	RocketID    string            `json:"rocket_id"`
	RocketName  string            `json:"rocket_name"`
	RocketType  string            `json:"rocket_type"`
	FirstStage  RocketFirstStage  `json:"first_stage"`
	SecondStage RocketSecondStage `json:"second_stage"`
}

//RocketFirstStage .
type RocketFirstStage struct {
	Cores []RocketFirstStageCore `json:"cores"`
}

// RocketFirstStageCore .
type RocketFirstStageCore struct {
	CoreSerial     string `json:"core_serial"`
	Flight         int    `json:"flight"`
	Block          int    `json:"block"`
	Gridfins       bool   `json:"gridfins"`
	Legs           bool   `json:"legs"`
	Reused         bool   `json:"reused"`
	LandSuccess    bool   `json:"land_success"`
	LandingIntent  bool   `json:"landing_intent"`
	LandingType    string `json:"landing_type"`
	LandingVehicle string `json:"landing_vehicle"`
}

//RocketSecondStage .
type RocketSecondStage struct {
	Block    int               `json:"block"`
	Payloads []payload.Payload `json:payloads`
}

//LaunchFairing .
type LaunchFairing struct {
	Reused          bool   `json:"reused"`
	RecoveryAttempt bool   `json:"recovery_attempt"`
	Recovered       bool   `json:"recovered"`
	Ship            string `json:"ship"`
}

//LaunchTelemetry .
type LaunchTelemetry struct {
	FlightClub string `json:"flight_club"`
}

//LaunchSite .
type LaunchSite struct {
	SiteID       string `json:"site_id"`
	SiteName     string `json:"site_name"`
	SiteNameLong string `json:"site_name_long"`
}

//LaunchLinks .
type LaunchLinks struct {
	MissionPatch      string   `json:"mission_patch"`
	MissionPatchSmall string   `json:"mission_patch_small"`
	RedditCampaign    string   `json:"reddit_campaign"`
	RedditLaunch      string   `json:"reddit_launch"`
	RedditRecovery    string   `json:"reddit_recovery"`
	RedditMedia       string   `json:"reddit_media"`
	Presskit          string   `json:"presskit"`
	ArticleLink       string   `json:"article_link"`
	Wikipedia         string   `json:"wikipedia"`
	VideoLink         string   `json:"video_link"`
	YoutubeID         string   `json:"youtube_id"`
	FlickrImages      []string `json:"flickr_images"`
}

//LaunchTimeline .
type LaunchTimeline struct {
	WebcastLiftoff           int `json:"webcast_liftoff"`
	GoForPropLoading         int `json:"go_for_prop_loading"`
	Rp1Loading               int `json:"rp1_loading"`
	Stage1LoxLoading         int `json:"stage1_lox_loading"`
	Stage2LoxLoading         int `json:"stage2_lox_loading"`
	EngineChill              int `json:"engine_chill"`
	PrelaunchChecks          int `json:"prelaunch_checks"`
	PropellantPressurization int `json:"propellant_pressurization"`
	GoForLaunch              int `json:"go_for_launch"`
	Ignition                 int `json:"ignition"`
	Liftoff                  int `json:"liftoff"`
	Maxq                     int `json:"maxq"`
	Meco                     int `json:"meco"`
	StageSep                 int `json:"stage_sep"`
	SecondStageIgnition      int `json:"second_stage_ignition"`
	FairingDeploy            int `json:"fairing_deploy"`
	FirstStageEntryBurn      int `json:"first_stage_entry_burn"`
	Seco1                    int `json:"seco-1"`
	FirstStageLanding        int `json:"first_stage_landing"`
	SecondStageRestart       int `json:"second_stage_restart"`
	Seco2                    int `json:"seco-2"`
	PayloadDeploy            int `json:"payload_deploy"`
}

//LaunchPad .
type LaunchPad struct {
	ID                 int             `json:"id"`
	Status             string          `json:"status"`
	Location           common.Location `json:"location"`
	VehiclesLaunched   []string        `json:"vehicles_launched"`
	AttemptedLaunches  int             `json:"attempted_launches"`
	SuccessfulLaunches int             `json:"successful_launches"`
	Wikipedia          string          `json:"wikipedia"`
	Details            string          `json:"details"`
	SiteID             string          `json:"site_id"`
	SiteNameLong       string          `json:"site_name_long"`
}
