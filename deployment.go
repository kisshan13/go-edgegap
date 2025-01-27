package edgegap

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const DEPLOYMENT_ENDPOINT = "/deploy"

// Create a new deployment. Deployment is a server instance of your application version.
func (e *EdgegapClient) DeploymentCreate(data *DeployementCreatePayload) (*Response[DeploymentCreateResponse], error) {
	var successResponse DeploymentCreateResponse

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(data).Post(DEPLOYMENT_ENDPOINT)
	}, &successResponse)
}

func (e *EdgegapClient) DeploymentContainerLogs(requestId string) (*Response[DeploymentContainerLogs], error) {
	endpoint := fmt.Sprintf("%s/%s/container-logs", DEPLOYMENT_ENDPOINT, requestId)
	var containerLogs DeploymentContainerLogs

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(endpoint)
	}, &containerLogs)
}
