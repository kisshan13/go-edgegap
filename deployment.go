package edgegap

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const DEPLOYMENT_ENDPOINT = "/deploy"

type DeploymentUpdateResponse struct {
	IsJoinableBySession bool `json:"is_joinable_by_session,omitempty"`
}

type DeploymentAvailableSocketPayload struct {
	AppName        string   `json:"app_name"`                  // The name of the application
	AppVersion     string   `json:"app_version"`               // The name of the application version
	MinimumSockets int      `json:"minimum_sockets,omitempty"` // The minimum number of sockets required
	IPLists        []string `json:"ip_list,omitempty"`         // The list of IPs
	Location       string   `json:"location,omitempty"`        // The location of the deployment
	Latitude       string   `json:"latitude,omitempty"`        // The latitude of the deployment
	Longitude      string   `json:"longitude,omitempty"`       // The longitude of the deployment
}

type DeploymentInfo struct {
	RequestID          string                 `json:"request_id"`          // The Unique ID of the Deployment's request
	FDQN               string                 `json:"fdqn"`                // The FQDN that allow to connect to your Deployment
	AppName            string                 `json:"app_name"`            // The name of the deployed App
	AppVersion         string                 `json:"app_version"`         // The version of the deployed App
	CurrentStatus      string                 `json:"current_status"`      // The current status of the Deployment
	Running            bool                   `json:"running"`             // True if the current Deployment is ready to be connected and running
	WhitelistingActive bool                   `json:"whitelisting_active"` // True if the current Deployment is ACL protected
	StartTime          string                 `json:"start_time"`          // Timestamp of the Deployment when it is up and running
	RemovalTime        string                 `json:"removal_time"`        // Timestamp of the end of the Deployment
	ElapsedTime        int                    `json:"elapsed_time"`        // Time since the Deployment is up and running in seconds
	LastStatus         string                 `json:"last_status"`         // The last status of the Deployment
	Error              bool                   `json:"error"`               // True if there is an error with the Deployment
	ErrorDetails       bool                   `json:"error_details"`       // The error details of the Deployment
	Ports              map[string]PortDetails `json:"ports"`
	PublicIP           string                 `json:"public_ip"` // The public IP
	Sessions           []Session              `json:"sessions"`  // List of Active Sessions if Deployment App is Session Based
	Location           Location               `json:"location"`
	Tags               []string               `json:"tags"`          // List of tags associated with the deployment
	Sockets            int                    `json:"sockets"`       // The Capacity of the Deployment
	SocketsUsage       int                    `json:"sockets_usage"` // The Capacity Usage of the Deployment
	Command            string                 `json:"command"`       // The command to use in the container, null mean it will take the default of the container
	Arguments          string                 `json:"arguments"`     // The arguments to use in the container, null mean it will take the default of the container
	MaxDuration        int                    `json:"max_duration"`  // The deployment's maximum duration is the time, in minutes, that the deployment will remain active before automatically closing.
}

type DeployementCreatePayload struct {
	AppName                  string        `json:"app_name,omitempty"`      // The name of the App you want to deploy
	VersionName              string        `json:"version_name,omitempty"`  // The name of the App Version you want to deploy, if not present, the last version created is picked
	IsPublicApp              bool          `json:"is_public_app,omitempty"` // If the Application is public or private. If not specified, we will look for a private Application
	IpList                   []string      `json:"ip_list,omitempty"`       // The List of IP of your user
	GeoIPList                []GeoIPList   `json:"geo_ip_list,omitempty"`   // The list of IP of your user with their location (latitude, longitude)
	TelemetryProfileUUIDList []string      `json:"telemetry_profile_uuid_list,omitempty"`
	EnvVariables             []EnvVariabls `json:"env_vars,omitempty"`       // A list of deployment variables
	SkipTelemetry            bool          `json:"skip_telemetry,omitempty"` // If you want to skip the Telemetry and use a geolocations decision only
	Location                 Location      `json:"location,omitempty"`
	WebhookURL               string        `json:"webhook_url,omitempty"`      // A web URL. This url will be called with method POST. The deployment status will be send in JSON format
	Tags                     []string      `json:"tags,omitempty"`             // The list of tags for your deployment
	Filters                  []Filter      `json:"filters,omitempty"`          // Filters to use while choosing the deployment locatresty.
	ApSortStrategy           ESortStrategy `json:"ap_sort_strategy,omitempty"` // Algorithm used to select the edge location
	Command                  string        `json:"command,omitempty"`          // Allows to override the Container command for this deployment.
	Arguments                string        `json:"arguments,omitempty"`        // Allows to override the Container arguments for this deployment.
}

type DeploymentCreateResponse struct {
	RequestID           string              `json:"request_id,omitempty"`              // The Unique Identifier of the request
	RequestDNS          string              `json:"request_dns,omitempty"`             // The URL to connect to the instance
	RequestApp          string              `json:"request_app,omitempty"`             // The Name of the App you requested
	RequestVersion      string              `json:"request_version,omitempty"`         // The name of the App Version you requested
	RequestUserCount    int                 `json:"request_user_count,omitempty"`      // How Many Users your request contain
	City                string              `json:"city,omitempty"`                    // The city where the deployment is located
	Country             string              `json:"country,omitempty"`                 // The country where the deployment is located
	Continent           string              `json:"continent,omitempty"`               // The continent where the deployment is located
	AdminDivision       string              `json:"administrative_division,omitempty"` // The administrative division where the deployment is located
	Tags                []string            `json:"tags,omitempty"`                    // List of tags associated with the deployment
	ContainerLogStorage ContainerLogStorage `json:"container_log_storage,omitempty"`
}

type DeploymentContainerLogs struct {
	Logs      string               `json:"logs,omitempty"`
	Encoding  string               `json:"encoding,omitempty"`
	CrashLogs string               `json:"crash_logs,omitempty"`
	CrashData []ContainerCrashData `json:"crash_data,omitempty"`
	LogsLink  string               `json:"logs_link,omitempty"`
}

// Info about an active deployment
type Deployment struct {
	RequestID           string                 `json:"request_id"` // Unique UUID
	FQDN                string                 `json:"fqdn"`       // The FQDN that allow to connect to your deployment
	StartTime           string                 `json:"start_time"` // Timestamp of the deployment when it is up and running
	Ready               bool                   `json:"ready"`      // If the deployment is ready
	PublicIP            string                 `json:"public_ip"`  // The public IP
	Ports               map[string]PortDetails `json:"ports"`
	Tags                []string               `json:"tags,omitempty"`                   // List of tags associated with the deployment
	Sockets             string                 `json:"sockets,omitempty"`                // The capacity of the deployment
	SocketsUsage        string                 `json:"sockets_usage,omitempty"`          // The capacity usage of the deployment
	IsJoinableBySession bool                   `json:"is_joinable_by_session,omitempty"` // If the deployment is joinable by sessions
}

type DeploymentBulkDelete struct {
	RequestIDs []string `json:"processable,omitempty"`
}

// Create a new deployment. Deployment is a server instance of your application version.
func (e *EdgegapClient) DeploymentCreate(data *DeployementCreatePayload) (*Response[DeploymentCreateResponse], error) {
	var successResponse DeploymentCreateResponse

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(data).Post(DEPLOYMENT_ENDPOINT)
	}, &successResponse)
}

// Retrieve the logs of your container. Logs are not available when your deployment is terminated
func (e *EdgegapClient) DeploymentContainerLogs(requestId string) (*Response[DeploymentContainerLogs], error) {
	endpoint := fmt.Sprintf("%s/%s/container-logs", DEPLOYMENT_ENDPOINT, requestId)
	var containerLogs DeploymentContainerLogs

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(endpoint)
	}, &containerLogs)
}

// List all deployments.
func (e *EdgegapClient) DeploymentListAll() (*Response[ResponseBody[Deployment]], error) {
	var deploymentList ResponseBody[Deployment]

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get("/deployments")
	}, &deploymentList)
}

// Make a bulk delete of deployments using filters. All the deployments matching the given filters will be permanently deleted.
func (e *EdgegapClient) DeploymentBulkDelete(filters []Filter) (*Response[DeploymentBulkDelete], error) {
	var bulkDeleteResponse DeploymentBulkDelete

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(map[string]interface{}{
			"filters": filters,
		}).Post("/deployments/bulk-stiop")
	}, &bulkDeleteResponse)
}

// Updates properties of a deployment. Currently only the is_joinable_by_session property can be updated.
func (e *EdgegapClient) DeploymentPropertyUpdate(requestId string, isJoinableSession bool) (*Response[DeploymentUpdateResponse], error) {
	var response DeploymentUpdateResponse

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(map[string]interface{}{
			"is_joinable_by_session": isJoinableSession,
		}).Post(fmt.Sprintf("/deployments/%s", requestId))
	}, &response)
}

// Get the list of deployments that have available sockets sorted by proximity to the geographical data.
func (e *EdgegapClient) DeploymentWithAvailableSockets(data DeploymentAvailableSocketPayload) (*Response[ResponseBody[Deployment]], error) {
	var deploymentList ResponseBody[Deployment]

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(data).Post("/deployments:available")
	}, &deploymentList)
}

// Retrieve the information for a deployment.
func (e *EdgegapClient) DeploymentGetStatus(request_id string) (*Response[DeploymentInfo], error) {
	var deploymentInfo DeploymentInfo

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/status/%s", request_id))
	}, &deploymentInfo)
}
