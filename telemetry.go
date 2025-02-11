// Telemetry
// Active Deployment Telemetry API - Please refer to this documentation (https://docs.edgegap.com/docs/deployment/active-deployment-telemetry) to get started with active deployment telemetry.
// API Reference : https://docs.edgegap.com/api/#tag/Telemetry

package edgegap

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type TelemetryCreate struct {
	Deployments []string `json:"deployments,omitempty"` // List of Deployment request ID to get telemetry.
	IPs         []string `json:"ips,omitempty"`
	WebhookURL  string   `json:"webhook_url,omitempty"` // Webhook URL that we should call to send the telemetry response back.
}

type TelemetryCreateRes struct {
	RetrievalKey string `json:"retrieval_key,omitempty"` // Unique retrieval key to get the telemetry response.
	Expire       string `json:"expire,omitempty"`        // Expiration date of the retrieval key.
}

type Telemetry struct {
	RetrievalKey  string   `json:"retrieval_key,omitempty"`  // Unique retrieval key to get the telemetry response.
	Scores        []string `json:"scores,omitempty"`         // Result sorted by best score. Index 0 is the best one.
	PartialResult bool     `json:"partial_result,omitempty"` // If the score list is incomplete and missing request IDs. Can occur if you request the results before we receive telemetry from every deployment.
}

// Create a telemetry request to get the best deployment(s) for given IP(s). You can use this to add players on a running deployment. If you set a webhook URL, the result will be sent to it.
func (e *EdgegapClient) TelemetryCreate(payload TelemetryCreate) (*Response[TelemetryCreateRes], error) {
	var response TelemetryCreateRes

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.SetBody(payload).Post("/telemetry/active-deployments")
	}, &response)
}

// Retrieve the results of a telemetry request on active deployment(s) for given IP(s). The score array is sorted from the best to the worse deployment. You can use this to add players on a running deployment.
func (e *EdgegapClient) TelemetryList(id string) (*Response[Telemetry], error) {
	var response Telemetry

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/telemetry/active-deployments/%s", id))
	}, &response)
}
