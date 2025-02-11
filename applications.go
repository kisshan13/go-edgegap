// Applications
// Applications Control API - Please refer to this documentation (https://docs.edgegap.com/docs/application) to get started with applications.
// API Reference - https://docs.edgegap.com/api/#tag/Applications

package edgegap

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ApplicationSessionKind string

const (
	SessionDefault = ApplicationSessionKind("Default")
	SessionSeat    = ApplicationSessionKind("Seat")
	SessionMatch   = ApplicationSessionKind("Match")
)

type ApplicationCreate struct {
	Name                   string `json:"name"`                                // The application name
	IsActive               bool   `json:"is_active"`                           // If the application can be deployed
	IsTelemetryAgentActive bool   `json:"is_telemetry_agent_active,omitempty"` // If the telemetry agent is installed on the versions of this app.
	Image                  string `json:"image"`                               // Image base64 string
}

type Application struct {
	Name                   string `json:"name"`                                // The application name
	IsActive               bool   `json:"is_active"`                           // If the application can be deployed
	IsTelemetryAgentActive bool   `json:"is_telemetry_agent_active,omitempty"` // If the telemetry agent is installed on the versions of this app.
	Image                  string `json:"image"`                               // Image base64 string
	CreateTime             string `json:"create_time"`
	LastUpdate             string `json:"last_updated"`
}

type ApplicationVersionSession struct {
	Kind               ApplicationSessionKind `json:"kind"`                           // The kind of session to create. If 'Default' if chosen, the 'session_config' will be ignored. The kind of session must be: Default, Seat, Match
	Sockets            int                    `json:"sockets"`                        // The number of game slots on each deployment of this app version.
	AutoDeploy         bool                   `json:"autodeploy,omitempty"`           // If a deployment should be made autonomously if there is not enough sockets open on a new session.
	EmptyTTL           int                    `json:"empty_ttl,omitempty"`            // The number of minutes a deployment of this app version can spend with no session connected before being terminated.
	SessionMaxDuration int                    `json:"session_max_duration,omitempty"` // The number of minutes after a session-type deployment has been terminated to remove all the session information connected to your deployment. Minimum and default value is set to 60 minutes so you can manage your session termination before it is removed.
}

type ApplicationPort struct {
	Port       int      `json:"port"`                  // The Port to Expose your service. Port 0 reserved for one-to-one port mapping. See our doc for more information.
	Protocol   Protocol `json:"protocol"`              // Available protocols: TCP, UDP, TCP/UDP, HTTP, HTTPS, WS or WSS
	ToCheck    bool     `json:"to_check,omitempty"`    // If the port must be verified by our port validations
	TLSUpgrade bool     `json:"tls_upgrade,omitempty"` // Enabling with HTTP or WS will inject a sidecar proxy that upgrades the connection with TLS
	Name       string   `json:"name,omitempty"`        // An optional name for the port for easier handling. Mandatory if using port 0
}

type ApplicationProbe struct {
	OptimalPing  int `json:"optimal_ping,omitempty"`  // Your optimal value for Latency
	RejectedPing int `json:"rejected_ping,omitempty"` // Your reject value for Latency
}

type ApplicationList struct {
	Applications []Application `json:"applications"` // List of applications
}

type ApplicationVersionCreateResponse struct {
	Success bool               `json:"success"`
	Version ApplicationVersion `json:"version"`
}

type ApplicationACLCreateResponse struct {
	Success        bool           `json:"success"`
	WhiteListEntry ApplicationACL `json:"whitelist_entry"`
}

type ApplicationACLEntries struct {
	WhitelistEntries []ApplicationACL `json:"whitelist_entries"`
}

type ApplicationVersionList struct {
	Versions   []ApplicationVersion `json:"versions"`
	TotalCount int                  `json:"total_count"`
}

type ApplicationACL struct {
	ID       string `json:"id,omitempty"`
	CIDR     string `json:"cidr"`                // CIDR to allow
	Label    string `json:"label,omitempty"`     // Label to organized your entries
	IsActive bool   `json:"is_active,omitempty"` // If the Rule will be applied on runtime
}

type ApplicationVersion struct {
	Name               string                    `json:"name"`                          // The Version Name
	IsActive           bool                      `json:"is_active,omitempty"`           // If the Version is active currently in the system
	DockerRepo         string                    `json:"docker_repository"`             // The Repository where the image is (i.e. 'harbor.edgegap.com' or 'docker.io')
	DockerImage        string                    `json:"docker_image"`                  // The name of your image (i.e. 'edgegap/demo')
	DockerTag          string                    `json:"docker_tag"`                    // The tag of your image (i.e. '0.1.2')
	PrivateUsername    string                    `json:"private_username,omitempty"`    // The username to access the docker repository
	PrivateToken       string                    `json:"private_token,omitempty"`       // The Private Password or Token of the username (We recommend to use a token)
	ReqCPU             int                       `json:"req_cpu"`                       // Units of vCPU needed (1024 = 1vcpu)
	ReqMemory          int                       `json:"req_memory"`                    // Units of memory in MB needed (1024 = 1GB)
	ReqVideo           int                       `json:"req_video,omitempty"`           // Units of GPU needed (1024 = 1 GPU)
	MaxDuration        int                       `json:"max_duration,omitempty"`        // The Max duration of the game in minute. 0 means forever.
	UseTelemetry       bool                      `json:"use_telemetry,omitempty"`       // Allow to inject ASA Variables
	WhitelistingActive bool                      `json:"whitelisting_active,omitempty"` // ACL Protection is active
	ForceCache         bool                      `json:"force_cache,omitempty"`         // Allow faster deployment by caching your container image in every Edge site
	CacheMinHour       int                       `json:"cache_min_hour,omitempty"`      // Start of the preferred interval for caching your container
	CacheMaxHour       int                       `json:"cache_max_hour,omitempty"`      // End of the preferred interval for caching your container
	TimeToDeploy       int                       `json:"time_to_deploy,omitempty"`      // Estimated maximum time in seconds to deploy, after this time we will consider it not working and retry.
	SessionConfig      ApplicationVersionSession `json:"session_config,omitempty"`
	Ports              []ApplicationPort         `json:"ports,omitempty"`
	Probe              ApplicationProbe          `json:"probe,omitempty"`
	Envs               []EnvVariabls             `json:"envs,omitempty"`
	VerifyImage        bool                      `json:"verify_image,omitempty"`                     // By enabling the verify_image option, your image infos (docker_repository, docker_image, docker_tag) will be tested.
	TerminationPeriod  int                       `json:"termination_grace_period_seconds,omitempty"` // Termination grace period in seconds after the SIGTERM signal has been sent
	EndpointStorage    string                    `json:"endpoint_storage,omitempty"`                 // The name of the endpoint storage to link
	Command            string                    `json:"command,omitempty"`                          // Entrypoint/Command override of your Container
	Arguments          string                    `json:"arguments,omitempty"`                        // The Arguments to pass to the command
	BuildType          BuildType                 `json:"build_type,omitempty"`                       // Available Build Types: Production or Development
}

// Create an application that will regroup application versions.
func (e *EdgegapClient) ApplicationCreate(application ApplicationCreate) (*Response[Application], error) {
	var response Application

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(application).Post("/app")
	}, &response)
}

// Update an application with new information.
func (e *EdgegapClient) ApplicationUpdate(name string, application ApplicationCreate) (*Response[Application], error) {
	var response Application

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(application).Patch(fmt.Sprintf("/app/%s", name))
	}, &response)
}

// Delete an application and all its current versions.
func (e *EdgegapClient) ApplicationDelete(name string) (*Response[map[string]interface{}], error) {
	var response map[string]interface{}

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/app/%s", name))
	}, &response)
}

// Retrieve an application and its information.
func (e *EdgegapClient) Application(name string) (*Response[Application], error) {
	var response Application

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/app/%s", name))
	}, &response)
}

// Create an application version associated with an application. The version contains all the specifications to create a deployment.
func (e *EdgegapClient) ApplicationCreateVersion(appName string, version ApplicationVersion) (*Response[ApplicationVersionCreateResponse], error) {
	var response ApplicationVersionCreateResponse

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(version).Post(fmt.Sprintf("/app/%s/version", appName))
	}, &response)
}

// Delete a specific version of an application.
func (e *EdgegapClient) ApplicationDeleteVersion(appName string, version string) (*Response[map[string]interface{}], error) {
	var response map[string]interface{}

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/app/%s/version/%s", appName, version))
	}, &response)
}

// Retrieve the specifications of an application version.
func (e *EdgegapClient) ApplicationGetVersion(appName string, version string) (*Response[ApplicationVersion], error) {
	var response ApplicationVersion

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/app/%s/version/%s", appName, version))
	}, &response)
}

// Update an application version with new specifications.
func (e *EdgegapClient) ApplicationUpdateVersion(appName string, version string, data ApplicationVersion) (*Response[ApplicationVersionCreateResponse], error) {
	var response ApplicationVersionCreateResponse

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Patch(fmt.Sprintf("/app/%s/version/%s", appName, version))
	}, &response)
}

// Create an access control list entry for an app version. This will allow the specified CIDR to connect to the deployment. The option whitelisting_active must be activated in the application version.
func (e *EdgegapClient) ApplicationCreateACLEntry(appName string, version string, data ApplicationACL) (*Response[ApplicationACLCreateResponse], error) {
	var response ApplicationACLCreateResponse

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(data).Post(fmt.Sprintf("/app/%s/version/%s/whitelist", appName, version))
	}, &response)
}

// List all the access control list entries for a specific application version.
func (e *EdgegapClient) ApplicationACLEntries(appName string, version string) (*Response[ApplicationACLEntries], error) {
	var response ApplicationACLEntries

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/app/%s/version/%s/whitelist", appName, version))
	}, &response)
}

// Delete an access control list entry for a specific application version
func (e *EdgegapClient) ApplicationDeleteACL(appName string, version string, entryId string) (*Response[ApplicationACLCreateResponse], error) {
	var response ApplicationACLCreateResponse

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/app/%s/version/%s/whitelist/%s", appName, version, entryId))
	}, &response)
}

// Retrieve a specific access control list entry for an application version.
func (e *EdgegapClient) ApplicationGetACLById(appName string, version string, entryId string) (*Response[ApplicationACL], error) {
	var response ApplicationACL

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/app/%s/version/%s/whitelist/%s", appName, version, entryId))
	}, &response)
}

// List all versions of a specific application.
func (e *EdgegapClient) ApplicationListVersion(appName string) (*Response[ApplicationVersionList], error) {
	var response ApplicationVersionList

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/app/%s/versions", appName))
	}, &response)
}

// List all the applications that you own.
func (e *EdgegapClient) ApplicationGetList() (*Response[ApplicationList], error) {
	var response ApplicationList

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get("/apps")
	}, &response)
}
