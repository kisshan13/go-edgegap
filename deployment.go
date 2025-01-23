package edgegap

import "fmt"

const DEPLOYMENT_ENDPOINT = "/deploy"

// Create a new deployment. Deployment is a server instance of your application version.
func (e *EdgegapClient) DeploymentCreate(data *DeployementCreatePayload) (*Response[DeploymentCreateResponse], error) {
	var errorResponse ErrorResponse
	var successResponse DeploymentCreateResponse

	res, err := e.client.R().SetBody(data).SetError(&errorResponse).SetResult(&successResponse).Post(DEPLOYMENT_ENDPOINT)

	if err != nil {
		return &Response[DeploymentCreateResponse]{
			Success:  false,
			Data:     nil,
			Response: res,
			Error:    err,
		}, err
	}

	if res.StatusCode() > 300 {
		err := fmt.Errorf("%s", errorResponse.Message)

		return &Response[DeploymentCreateResponse]{
			Success:  false,
			Data:     nil,
			Response: res,
			Error:    err,
		}, err
	}

	return &Response[DeploymentCreateResponse]{
		Success:  true,
		Response: res,
		Data:     &successResponse,
		Error:    nil,
	}, nil
}
