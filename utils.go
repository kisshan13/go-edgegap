package edgegap

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Utility function to get query paramters for pagination parameters.
func (pp *PaginationParams) GetParams() string {
	return fmt.Sprintf("?page=%d&limit=%d", pp.Page, pp.Size)
}

func makeRequest[T any](e *EdgegapClient, request func(c *resty.Request) (*resty.Response, error), response *T) (*Response[T], error) {
	var errorResponse ErrorResponse

	res, err := request(e.client.R().SetError(&errorResponse).SetResult(response))

	if err != nil {
		return &Response[T]{
			Success:  false,
			Data:     nil,
			Response: res,
			Error:    err,
		}, err
	}

	if res.StatusCode() > 300 {
		err := fmt.Errorf("API Error : %s", errorResponse.Message)

		return &Response[T]{
			Success:  false,
			Data:     nil,
			Response: res,
			Error:    err,
		}, err
	}

	return &Response[T]{
		Success:  true,
		Response: res,
		Data:     response,
		Error:    nil,
	}, nil
}
