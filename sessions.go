// Sessions Control API
// Please refer to this documentation (https://docs.edgegap.com/docs/deployment/session) to get started with sessions.

package edgegap

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SessionCreate struct {
	App                 string          `json:"app"`                             // The Name of the App you want to deploy
	Version             string          `json:"version_name,omitempty"`          // The Name of the App Version you want to deploy
	IPList              []string        `json:"ip_list,omitempty"`               // The List of IP of your user, Array of String
	GeoIPList           []GeoIPList     `json:"geo_ip_list,omitempty"`           // The list of IP of your user with their location (latitude, longitude)
	DeploymentRequestID string          `json:"deployment_request_id,omitempty"` // The request id of your deployment. If specified, the session will link to the deployment
	Location            Location        `json:"location,omitempty"`
	City                string          `json:"city,omitempty"`                    // If you want your session in a specific city
	Country             string          `json:"country,omitempty"`                 // If you want your session in a specific country
	Continent           string          `json:"continent,omitempty"`               // If you want your session in a specific continent
	AdminDivision       string          `json:"administrative_division,omitempty"` // If you want your session in a specific administrative division
	Region              string          `json:"region,omitempty"`                  // If you want your session in a specific region
	Selectors           []SelectorModel `json:"selectors,omitempty"`               // List of Selectors to filter potential Deployment to link and tag the Session
	WebhookURL          string          `json:"webhook_url,omitempty"`             // When your Session is Linked, Unprocessable or in Error, we will POST the session's details on the webhook_url
	Filters             []Filter        `json:"filters,omitempty"`                 // List of location filters to apply to the session
	SkipTelemetry       bool            `json:"skip_telemetry,omitempty"`          // If system should skip the telemetry and use GeoBase decision only
}

type Session struct {
	ID         string         `json:"session_id"`    // Unique UUID
	CustomID   string         `json:"custom_id"`     // Custom ID if Available
	Status     string         `json:"status"`        // Current status of the session
	Ready      bool           `json:"ready"`         // If the session is linked to a ready deployment
	Linked     bool           `json:"linked"`        // If the session is linked to a deployment
	Kind       string         `json:"kind"`          // Type of session created
	UserCount  int            `json:"user_count"`    // Count of user this session currently have
	Version    int            `json:"app_version"`   // App version linked to the session
	CreateTime string         `json:"create_time"`   // Session created at
	Elapsed    int            `json:"elapsed"`       // Elapsed time
	Error      string         `json:"error"`         // Error Detail
	Users      []SessionUser  `json:"session_users"` // Users in the session
	IPs        []SessionUser  `json:"session_ips"`   // IPs in the session
	Deployment DeploymentInfo `json:"deployment"`
	WebhookURL string         `json:"webhook_url"` // When your Session is Linked, Unprocessable or in Error, we will POST the session's details on the webhook_url
}

type SessionCreateRes struct {
	SessionID           string          `json:"session_id,omitempty"`
	CustomID            string          `json:"custom_id,omitempty"`
	App                 string          `json:"app,omitempty"`
	Version             string          `json:"version,omitempty"`
	DeploymentRequestId string          `json:"deployment_request_id,omitempty"`
	Selectors           []SelectorModel `json:"selectors,omitempty"`
	WebhookURL          string          `json:"webhook_url,omitempty"`
}

type SessionDeleteRes struct {
	Message   string `json:"message"`    // A message depending of the request termination
	SessionID string `json:"session_id"` // The Unique Identifier of the Session
	CustomID  string `json:"custom_id"`  // Custom ID if Available
}

type SessionUser struct {
	IP        string `json:"ip"`
	Latitude  int    `json:"latitude"`
	Longitude int    `json:"longitude"`
}

// Create a session with users. Sessions are linked to a deployment.
func (e *EdgegapClient) SessionCreate(session *SessionCreate) (*Response[SessionCreateRes], error) {
	var response SessionCreateRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(session).Post("/session")
	}, &response)
}

// Delete a session. Once deleted, a session is no more accessible and does not have a history. The deployment associated will not be deleted.
func (e *EdgegapClient) SessionDelete(id string) (*Response[SessionDeleteRes], error) {
	var response SessionDeleteRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/session/%s", id))
	}, &response)
}

// Retrieve the information for a session.
func (e *EdgegapClient) SessionGet(id string) (*Response[Session], error) {
	var response Session

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/session/%s", id))
	}, &response)
}
