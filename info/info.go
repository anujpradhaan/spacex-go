package info

import (
	"spacex-go/common"
)

type CompanyInfo struct {
	Name                string         `json:"name"`
	Founder             string         `json:"founder"`
	Founded             string         `json:"founded"`
	NumberOfEmployees   int            `json:"employees"`
	NumberOfVehicles    int            `json:"vehicles"`
	NumberOfLaunchSites int            `json:"launch_sites"`
	NumberOfTestSites   int            `json:"test_sites"`
	Ceo                 string         `json:"ceo"`
	Cto                 string         `json:"cto"`
	Coo                 string         `json:"coo"`
	CtoPropulsion       string         `json:"cto_propulsion"`
	Valuation           string         `json:"valuation"`
	HeadQuarter         common.Address `json:"headquarters"`
	Summary             string         `json:"founder"`
}

type ApiInfo struct {
	ProjectName      string `json:"project_name"`
	Version          string `json:"version"`
	ProjectLink      string `json:"project_link"`
	Organization     string `json:"organization"`
	OrganizationLink string `json:"organization_link"`
	Description      string `json:"description"`
}
