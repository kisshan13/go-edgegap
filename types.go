package edgegap

import "github.com/go-resty/resty/v2"

type Version string
type Protocol string
type BuildType string

type EField string
type EFilterType string
type ESortStrategy string

const (
	VersionOne = Version("v1")
)

const (
	ProtocolTCP       = Protocol("TCP")
	ProtocolUDP       = Protocol("UDP")
	ProtocolTCPAndUDP = Protocol("TPC/UDP")
	ProtocolHTTP      = Protocol("HTTP")
	ProtocolHTTPS     = Protocol("HTTPS")
	ProtocolWS        = Protocol("WS")
	ProtocolWSS       = Protocol("WSS")
)

const (
	DevelopmentBuild      = BuildType("Development")
	DevelopmentProduction = BuildType("Production")
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
	Message    string     `json:"message,omitempty"`
}

type PaginationParams struct {
	Page int
	Size int
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

type ContainerLogStorage struct {
	Enabled         bool   `json:"enabled,omitempty"` // Will override the app version container log storage for this deployment
	EndpointStorage string `json:"enpoint_storage"`   // The name of your endpoint storage. If container log storage is enabled without this parameter, we will try to take the app version endpoint storage. If there is no endpoint storage in your app version, the container logs will not be stored. If we don't find any endpoint storage associated with this name, the container logs will not be stored.
}

type ContainerCrashData struct {
	ExitCode     int    `json:"exit_code,omitempty"`
	Message      string `json:"message,omitempty"`
	RestartCount int    `json:"restart_count"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type PortDetails struct {
	External   int    `json:"external"`
	Internal   int    `json:"internal"`
	Protocol   string `json:"protocol"`
	Name       string `json:"name,omitempty"`
	TLSUpgrade bool   `json:"tls_upgrade,omitempty"`
	Link       string `json:"link,omitempty"`
	Proxy      int    `json:"proxy,omitempty"`
}

type DeploymentSession struct {
	SessionID string `json:"session_id"` // Unique UUID
	Status    string `json:"status"`     // Current status of the session
	Ready     bool   `json:"ready"`      // If the session is linked to a Ready deployment
	Linked    bool   `json:"linked"`     // If the session is linked to a deployment
	Kind      string `json:"kind"`       // Type of session created
	UserCount int    `json:"user_count"` // Count of user this session currently have
}

type SelectorModel struct {
	Tag     string      `json:"tag"`      // The Tag to filter potential Deployment with this Selector
	TagOnly bool        `json:"tag_only"` // If True, will not try to filter Deployment and only tag the Session
	Env     EnvVariabls `json:"evn"`      // Environment Variable to inject in new Deployment created by App Version with auto-deploy
}
