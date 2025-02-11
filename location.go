// Locations
// Locations API - Please refer to this documentation (https://docs.edgegap.com/docs/deployment/locations/beacons) to get started with locations beacons.
// API Reference - https://docs.edgegap.com/api/#tag/Locations

package edgegap

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type LocationInfo struct {
	City           string   `json:"city"`                    // City Name
	Continent      string   `json:"continent"`               // Continent Name
	Country        string   `json:"country"`                 // Country name
	Timezone       string   `json:"timezone"`                // Timezone name
	AdminiDivision string   `json:"administrative_division"` // Administrative Division
	Latitude       float64  `json:"latitude"`                // The Latitude in decimal
	Longitude      float64  `json:"longitude"`               // The Longitude in decimal
	Type           string   `json:"type"`                    // The type of location
	Tags           []string `json:"tags"`
}

type LocationListRes struct {
	List    []LocationInfo `json:"locations"`
	Message []string       `json:"message"` // Extra Messages for the query
}

type LocationBeaconRes struct {
	List  []LocationInfo `json:"locations"` // Total number of active location beacons
	Count int            `json:"count"`     // List of active location beacons
}

type LocationFilters struct {
	App     string // The App Name you want to filter with capacity
	Version string // The Version Name you want to filter with capacity
	Type    string // The type of the location
	Tags    string // Gets locations with tags. Set to: "true" to have the tags
}

// List all the locations available to deploy on. You can specify an application and a version to filter out the locations that don’t have enough resources to deploy this application version.
func (e *EdgegapClient) LocationListAll(filters LocationFilters) (*Response[LocationListRes], error) {
	query := "?"

	if filters.App != "" {
		query += fmt.Sprintf("app=%s&", filters.App)
	}
	if filters.Version != "" {
		query += fmt.Sprintf("version=%s&", filters.Version)
	}
	if filters.Type != "" {
		query += fmt.Sprintf("type=%s&", filters.Type)
	}
	if filters.Tags != "" {
		query += fmt.Sprintf("tags=%s&", filters.Tags)
	}

	var response LocationListRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/locations/%s", query))
	}, &response)
}

// List all the active location beacons. They can be used to ping them for your matchmaking system. You cannot deploy on beacons.
func (e *EdgegapClient) LocationListAllBeacons() (*Response[LocationBeaconRes], error) {
	var response LocationBeaconRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get("/locations/beacons")
	}, &response)
}
