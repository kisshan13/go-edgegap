// Matchmaker
// Matchmaker Control API - Please refer to this documentation (https://docs.edgegap.com/docs/matchmaker) to get started with matchmakers.
// API Reference : https://docs.edgegap.com/api/#tag/Matchmaker

package edgegap

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

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
