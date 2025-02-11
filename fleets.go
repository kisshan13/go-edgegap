// Fleets
// Fleets Control API - Please refer to this documentation (https://docs.edgegap.com/docs/deployment/session/fleet-manager/fleet) to get started with fleets.
// API Reference - https://docs.edgegap.com/api/#tag/Fleets

package edgegap

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type FleetCreatePayload struct {
	Name    string `json:"name,omitempty"`    // Name of the Fleet
	Enabled bool   `json:"enabled,omitempty"` // If the Fleet is enabled. Defaults to false.
}

type Fleet struct {
	Name        string `json:"name,omitempty"`         // Name of the Fleet
	Enabled     bool   `json:"enabled,omitempty"`      // If the Fleet is enabled. Defaults to false.
	CreateTime  string `json:"create_time,omitempty"`  // UTC time of fleet creation
	LastUpdated string `json:"last_updated,omitempty"` // UTC time of fleet last update
}

type FleetApplication struct {
	Name        string `json:"name,omitempty"`         // Name of the Fleet
	AppName     string `json:"app,omitempty"`          // Name of the linked app of the linked version
	AppVersion  string `json:"app_version,omitempty"`  // Name of the linked app version.
	Enabled     bool   `json:"enabled,omitempty"`      // If the Fleet is enabled. Defaults to false.
	CreateTime  string `json:"create_time,omitempty"`  // UTC time of fleet creation
	LastUpdated string `json:"last_updated,omitempty"` // UTC time of fleet last update
}

type FleetList struct {
	Fleets     []Fleet    `json:"fleets,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// Create a fleet. A fleet is a top-level object; you must create child resources to work properly.
func (e *EdgegapClient) FleetCreate(payload FleetCreatePayload) (*Response[Fleet], error) {
	var response Fleet

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(payload).Post("/feet")
	}, &response)
}

// Retrieve a fleet with its details.
func (e *EdgegapClient) FleetGet(name string) (*Response[Fleet], error) {
	var response Fleet

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/fleet/%s", name))
	}, &response)
}

// Update a fleet with new specifications
func (e *EdgegapClient) FleetUpdate(name string, payload FleetCreatePayload) (*Response[Fleet], error) {
	var response Fleet

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(payload).Patch(fmt.Sprintf("/fleet/%s", name))
	}, &response)
}

// Delete a fleet, its policies and links between the application versions.
func (e *EdgegapClient) FleetDelete(name string) (*Response[interface{}], error) {
	var response interface{}

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/fleet/%s", name))
	}, &response)
}

// List all the fleets you own.
func (e *EdgegapClient) FleetList() (*Response[FleetList], error) {
	var response FleetList

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get("/fleets")
	}, &response)
}

// Link an application version to a fleet. By linking this version, the fleet will automatically create deployments of this version according to the fleet policies.
func (e *EdgegapClient) FleetLinkApplication(fleet, app, version string) (*Response[FleetApplication], error) {
	var response FleetApplication

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Put(fmt.Sprintf("/fleet/%s/app/%s/version/%s", fleet, app, version))
	}, &response)
}

// Unlink an application version from a fleet. It will not delete the application version or the fleet
func (e *EdgegapClient) FleetUnlinkApplication(fleet, app, version string) (*Response[interface{}], error) {
	var response interface{}

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/fleet/%s/app/%s/version/%s", fleet, app, version))
	}, &response)
}
