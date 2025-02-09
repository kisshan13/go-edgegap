// Matchmaker
// Matchmaker Control API - Please refer to this documentation (https://docs.edgegap.com/docs/matchmaker) to get started with matchmakers.
// API Reference : https://docs.edgegap.com/api/#tag/Matchmaker

package edgegap

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Matchmaker struct {
	Name      string `json:"name"` // Name of the Matchmaker.
	URL       string `json:"url"`  // URL of the Matchmaker.
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type MatchmakerListRes struct {
	Count int          `json:"count"` // Number of matchmakers owned by the user.
	Data  []Matchmaker `json:"data"`
}

type MatchmakerReleaseConfig struct {
	Name          string `json:"name,omitempty"`          // Matchmaker configuration name. Must be unique.
	Configuration string `json:"configuration,omitempty"` // Matchmaker configuration, parsed as a string.
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

type MatchmakerReleaseConfigListRes struct {
	Count int                       `json:"count,omitempty"` // Number of matchmaker component envs for the component.
	Data  []MatchmakerReleaseConfig `json:"data"`
}

type MatchmakerReleaseCreate struct {
	Version                    string `json:"version,omitempty"`                       // Name of the matchmaker release. Should be unique, and will be used to differentiate your releases.
	FrontendComponentName      string `json:"frontend_component_name,omitempty"`       // Name of the matchmaker component to use as the Open Match frontend.
	DirectorComponentName      string `json:"director_component_name,omitempty"`       // Name of the matchmaker component to use as the Open Match director.
	MatchFunctionComponentName string `json:"match_function_component_name,omitempty"` // Name of the matchmaker component to use as the Open Match match function.
}

type MatchmakerRelease struct {
	AppName                    string `json:"app_name,omitempty"`                      // Name of the app to deploy using the matchmaker.
	Version                    string `json:"version,omitempty"`                       // Name of the matchmaker release. Should be unique, and will be used to differentiate your releases.
	VersionName                string `json:"version_name,omitempty"`                  // Name of the version of the specified app to deploy using the matchmaker.
	FrontendComponentName      string `json:"frontend_component_name,omitempty"`       // Name of the matchmaker component to use as the Open Match frontend.
	DirectorComponentName      string `json:"director_component_name,omitempty"`       // Name of the matchmaker component to use as the Open Match director.
	MatchFunctionComponentName string `json:"match_function_component_name,omitempty"` // Name of the matchmaker component to use as the Open Match match function.
	CreatedAt                  string `json:"created_at"`
	UpdatedAt                  string `json:"updated_at"`
}

type MatchmakerReleaseListRes struct {
	Count int                 `json:"count"`
	Data  []MatchmakerRelease `json:"data"`
}

type MatchmakerManagedRelease struct {
	AppName           string `json:"app_name,omitempty"`            // Name of the app to deploy using the matchmaker.
	VersionName       string `json:"version_name,omitempty"`        // Name of the version of the specified app to deploy using the matchmaker.
	Version           string `json:"version,omitempty"`             // Name of the matchmaker release. Should be unique, and will be used to differentiate your releases.
	ReleaseConfigName string `json:"release_config_name,omitempty"` // Name of the matchmaker release configuration to use for this managed release.
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type MatchmakerManagedReleaseCreate struct {
	Version           string `json:"version,omitempty"`             // Name of the matchmaker release. Should be unique, and will be used to differentiate your releases.
	ReleaseConfigName string `json:"release_config_name,omitempty"` // Name of the matchmaker release configuration to use for this managed release.
}

type MatchmakerComponentCreate struct {
	Name        string      `json:"name,omitempty"`        // Matchmaker component name. Must be unique.
	Repo        string      `json:"repository,omitempty"`  // Container repository where the component's image is hosted.
	Image       string      `json:"image,omitempty"`       // Container image to use for this component.
	Tag         string      `json:"tag,omitempty"`         // Tag of the container image to use for this component.
	Credentials Credentials `json:"credentials,omitempty"` // Private repo credentials to use for pulling the image, if applicable.
}

type MatchmakerComponent struct {
	Name        string      `json:"name"`                  // Matchmaker component name. Must be unique.
	Repo        string      `json:"repository"`            // Container repository where the component's image is hosted.
	Image       string      `json:"image"`                 // Container image to use for this component.
	Tag         string      `json:"tag"`                   // Tag of the container image to use for this component.
	Credentials Credentials `json:"credentials,omitempty"` // Private repo credentials to use for pulling the image, if applicable.
	CreatedAt   string      `json:"created_at"`
	UpdatedAt   string      `json:"updated_at"`
}

type MatchmakerEnvRes struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type MatchmakerEnv struct {
	Key   string `json:"key"`   // Name of the ENV variable.
	Value string `json:"value"` // Value of the ENV variable.
}

type MatchmakerEnvListRes struct {
	Count int                `json:"count"` // Number of matchmaker component envs for the component.
	Data  []MatchmakerEnvRes `json:"data"`
}

type MatchmakerComponentListRes struct {
	Count int                   `json:"count"` // Number of matchmaker components owned by the user.
	Data  []MatchmakerComponent `json:"data"`
}

// Create a new matchmaker component.
func (e *EdgegapClient) MatchmakerCreateComponent(component MatchmakerComponentCreate) (*Response[MatchmakerComponent], error) {
	var response MatchmakerComponent

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(component).Post("/aom/component")
	}, &response)
}

// Update a matchmaker component with new specifications.
func (e *EdgegapClient) MatchmakerUpdateComponent(name string, component MatchmakerComponentCreate) (*Response[MatchmakerComponent], error) {
	var response MatchmakerComponent

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(component).Patch(fmt.Sprintf("/aom/component/%s", name))
	}, &response)
}

// Delete a matchmaker component. It will not delete the matchmaker.
func (e *EdgegapClient) MatchmakerDeleteComponent(name string) (*Response[map[string]string], error) {
	var response map[string]string

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/aom/component/%s", name))
	}, &response)
}

// Retrieve a matchmaker component.
func (e *EdgegapClient) MatchmakerGetComponent(name string) (*Response[MatchmakerComponent], error) {
	var response MatchmakerComponent

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/aom/component/%s", name))
	}, &response)
}

// Create a new matchmaker component ENV.
func (e *EdgegapClient) MatchmakerComponentAddEnv(name string, env MatchmakerEnv) (*Response[MatchmakerEnvRes], error) {
	var response MatchmakerEnvRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(env).Post(fmt.Sprintf("/aom/component/%s/env", name))
	}, &response)
}

// Update a matchmaker component ENV.
func (e *EdgegapClient) MatchmakerComponentUpdateEnv(name string, env MatchmakerEnv) (*Response[MatchmakerEnvRes], error) {
	var response MatchmakerEnvRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(env).Patch(fmt.Sprintf("/aom/component/%s/env/%s", name, env.Key))
	}, &response)
}

// Delete a matchmaker component ENV. It will not delete the component or the matchmaker.
func (e *EdgegapClient) MatchmakerComponentDeleteEnv(name string, env string) (*Response[map[string]string], error) {
	var response map[string]string

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/aom/component/%s/env/%s", name, env))
	}, &response)
}

// Retrieve a matchmaker component ENV.
func (e *EdgegapClient) MatchmakeComponentGetEnv(name string, env string) (*Response[MatchmakerEnvRes], error) {
	var response MatchmakerEnvRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/aom/component/%s/env/%s", name, env))
	}, &response)
}

// List all ENVs for a specific matchmaker component.
func (e *EdgegapClient) MatchmakerComponentListEnv(name string) (*Response[MatchmakerEnvListRes], error) {
	var response MatchmakerEnvListRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/aom/component/%s/envs", name))
	}, &response)
}

/*
List all components for a specific matchmaker.
API Reference : https://docs.edgegap.com/api/#tag/Matchmaker/operation/get-component-list
*/
func (e *EdgegapClient) MatchmakerComponentList() (*Response[MatchmakerComponentListRes], error) {
	var response MatchmakerComponentListRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get("/aom/components")
	}, &response)
}

// Create a new matchmaker. A matchmaker is a top-level object; you must create child resources to work properly.
func (e *EdgegapClient) MatchmakerCreate(name string) (*Response[Matchmaker], error) {
	var response Matchmaker

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(map[string]string{
			"name": name,
		}).Post("/aom/matchmaker")
	}, &response)
}

// Update a matchmaker with new specifications.
func (e *EdgegapClient) MatchmakerUpdate(name string, newName string) (*Response[Matchmaker], error) {
	var response Matchmaker

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(map[string]string{
			"name": newName,
		}).Patch(fmt.Sprintf("/aom/matchmaker/%s", name))
	}, &response)
}

// Delete a matchmaker.
func (e *EdgegapClient) MatchmakerDelete(name string) (*Response[interface{}], error) {
	var response interface{}

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/aom/matchmaker/%s", name))
	}, &response)
}

// Retrieve a matchmaker.
func (e *EdgegapClient) MatchmakerGet(name string) (*Response[Matchmaker], error) {
	var response Matchmaker

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/aom/matchmaker/%s", name))
	}, &response)
}

func (e *EdgegapClient) MatchmakerList() (*Response[MatchmakerListRes], error) {
	var response MatchmakerListRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get("/aom/matchmakers")
	}, &response)
}

// Create a matchmaker release.
func (e *EdgegapClient) MatchmakerCreateRelease(name string, payload MatchmakerReleaseCreate) (*Response[MatchmakerRelease], error) {
	var response MatchmakerRelease

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(payload).Post(fmt.Sprintf("/aom/matchmaker/%s/release", name))
	}, &response)
}

// Update a matchmaker release.
func (e *EdgegapClient) MatchmakerUpdateRelease(name string, payload MatchmakerReleaseCreate) (*Response[MatchmakerRelease], error) {
	var response MatchmakerRelease

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(payload).Patch(fmt.Sprintf("/aom/matchmaker/%s/release", name))
	}, &response)
}

// Delete a matchmaker release.
func (e *EdgegapClient) MatchmakerDeleteRelease(name string, version string) (*Response[interface{}], error) {
	var response interface{}

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/aom/matchmaker/%s/release/%s", name, version))
	}, &response)
}

// Retrieve a matchmaker release.
func (e *EdgegapClient) MatchmakerGetRelease(name string, version string) (*Response[MatchmakerRelease], error) {
	var response MatchmakerRelease

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/aom/matchmaker/%s/release/%s", name, version))
	}, &response)
}

// List all releases of a specific matchmaker.
func (e *EdgegapClient) MatchmakerListRelease(name string) (*Response[MatchmakerReleaseListRes], error) {
	var response MatchmakerReleaseListRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/aom/matchmaker/%s/release", name))
	}, &response)
}

// Update a matchmaker managed release.
func (e *EdgegapClient) MatchmakerCreateManagedRelease(name string, payload MatchmakerManagedReleaseCreate) (*Response[MatchmakerManagedRelease], error) {
	var response MatchmakerManagedRelease

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(payload).Post(fmt.Sprintf("/aom/matchmaker/%s/release/managed", name))
	}, &response)
}

// Update a matchmaker managed release.
func (e *EdgegapClient) MatchmakerUpdateManagedRelease(name string, releaseVersion string, payload MatchmakerManagedReleaseCreate) (*Response[MatchmakerManagedRelease], error) {
	var response MatchmakerManagedRelease

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(payload).Patch(fmt.Sprintf("/aom/matchmaker/%s/release/managed/%s", name, releaseVersion))
	}, &response)
}

// Delete a matchmaker managed release. It will not delete the matchmaker.
func (e *EdgegapClient) MatchmakerDeleteManagedRelease(name string, releaseVersion string) (*Response[interface{}], error) {
	var response interface{}

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/aom/matchmaker/%s/release/managed/%s", name, releaseVersion))
	}, &response)
}

// Retrieve a matchmaker managed release.
func (e *EdgegapClient) MatchmakerGetManagedRelease(name string, releaseVersion string) (*Response[MatchmakerManagedRelease], error) {
	var response MatchmakerManagedRelease

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/aom/matchmaker/%s/release/managed/%s", name, releaseVersion))
	}, &response)
}

// Create a matchmaker release config.
func (e *EdgegapClient) MatchmakerCreateReleaseConfig(payload MatchmakerReleaseConfig) (*Response[MatchmakerReleaseConfig], error) {
	var response MatchmakerReleaseConfig

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get("/aom/release/config")
	}, &response)
}

// Update a matchmaker release config.
func (e *EdgegapClient) MatchmakerUpdateReleaseConfig(name string, payload MatchmakerReleaseConfig) (*Response[MatchmakerReleaseConfig], error) {
	var response MatchmakerReleaseConfig

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/aom/release/config/%s", name))
	}, &response)
}

// Delete a matchmaker release config.
func (e *EdgegapClient) MatchmakerDeleteReleaseConfig(name string) (*Response[interface{}], error) {
	var response interface{}

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Delete(fmt.Sprintf("/aom/release/config/%s", name))
	}, &response)
}

// Get a matchmaker release config.
func (e *EdgegapClient) MatchmakerGetReleaseConfig(name string) (*Response[MatchmakerReleaseConfig], error) {
	var response MatchmakerReleaseConfig

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/aom/release/config/%s", name))
	}, &response)
}

// List all configs for a specific matchmaker release.
func (e *EdgegapClient) MatchmakerListReleaseConfig() (*Response[MatchmakerReleaseConfigListRes], error) {
	var respoonse MatchmakerReleaseConfigListRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get("/aom/release/config")
	}, &respoonse)
}
