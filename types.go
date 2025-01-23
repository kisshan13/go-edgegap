package edgegap

import "github.com/go-resty/resty/v2"

type Version string

type EField string
type EFilterType string
type ESortStrategy string

const (
	VersionOne = Version("v1")
)

const (
	ECity          = EField("city")
	ECountry       = EField("country")
	EContinent     = EField("continent")
	ERegion        = EField("region")
	EAdminDivision = EField("administrative_division")
	ELocationTags  = EField("location_tags")

	EAny = EFilterType("any")
	EAll = EFilterType("all")
	ENot = EFilterType("not")

	EBasic    = ESortStrategy("basic")
	EWeighted = ESortStrategy("weighted")
)

const EDGEGAP_BASE_URL = "https://api.edgegap.com"

type Response[T any] struct {
	Success  bool            // Represent if a request successfully completed or not.
	Response *resty.Response // Raw response for the request
	Data     *T              // Data from the server on request complete
	Error    error           // error
}

type Pagination struct {
	Number             int       `json:"number,omitempty"`               // The Current page, default=1
	NextPageNumber     int       `json:"next_page_number,omitempty"`     // The Next page number or null
	PreviousPageNumber int       `json:"previous_page_number,omitempty"` // The Previous page number or null
	Paginator          Paginator `json:"paginator,omitempty"`            // Pagination info (like total pages)
	HasNext            bool      `json:"has_next,omitempty"`
	HasPrevious        bool      `json:"has_previous,omitempty"`
}

type Paginator struct {
	NumPages int `json:"num_pages,omitempty"` // Total number of pages
}

type ResponseBody[T any] struct {
	Count      int        `json:"count,omitempty"`
	Data       []T        `json:"data"`
	Success    bool       `json:"success"`
	Pagination Pagination `json:"pagination"`
}

type PaginationParams struct {
	Page int
	Size int
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

type EnvVariabls struct {
	Key      string `json:"key,omitempty"`       // The Key to retrieve the value in your instance
	Value    string `json:"value,omitempty"`     // The value to set in your instance
	IsHidden bool   `json:"is_hidden,omitempty"` // If set to true, the value will be encrypted during the process of deployment
}

type GeoIPList struct {
	IP        string  `json:"ip,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type Location struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type Filter struct {
	Field      EField      `json:"field,omitempty"`       // Auto Generated Field for field
	Values     []string    `json:"values,omitempty"`      // Auto Generated Field for values
	FilterType EFilterType `json:"filter_type,omitempty"` // Auto Generated Field for filter_type
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

type ContainerLogStorage struct {
	Enabled         bool   `json:"enabled,omitempty"` // Will override the app version container log storage for this deployment
	EndpointStorage string `json:"enpoint_storage"`   // The name of your endpoint storage. If container log storage is enabled without this parameter, we will try to take the app version endpoint storage. If there is no endpoint storage in your app version, the container logs will not be stored. If we don't find any endpoint storage associated with this name, the container logs will not be stored.
}

type ErrorResponse struct {
	Message string `json:"message"`
}
