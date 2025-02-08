// Metrics
// Metrics API - Please refer to this documentation (https://docs.edgegap.com/docs/deployment/metrics) to get started with your deployment metrics.
// API Reference - https://docs.edgegap.com/api/#tag/Metrics

package edgegap

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type Steps string

const (
	SSeconds = Steps("s")
	SMinutes = Steps("m")
	SHour    = Steps("h")
)

type MetricsFilter struct {
	StartTime *time.Time
	EndTime   *time.Time
	StepValue *int
	Step      *Steps
	Raw       *string
}

type MetricsModel struct {
	Labels     []string `json:"labels"`
	Datasets   []string `json:"datasets"`
	Timestamps []string `json:"timestamps"`
}

type MetricsTotalModel struct {
	RecieveTotal   MetricsModel `json:"receive_total"`
	TransmitTotal  MetricsModel `json:"transmit_total"`
	DiskReadTotal  MetricsModel `json:"disk_read_total"`
	DiskWriteTotal MetricsModel `json:"disk_write_total"`
}

type MetricsNetworkModel struct {
	Recieve  MetricsModel `json:"receive"`
	Transmit MetricsModel `json:"transmit"`
}

type Metrics struct {
	Total   MetricsTotalModel   `json:"total"`
	CPU     MetricsModel        `json:"cpu"`
	Memory  MetricsModel        `json:"mem"`
	Network MetricsNetworkModel `json:"network"`
}

// Get the metrics for a specific deployment based on the start_time, end_time and steps. raw parameter can be set to true to get the raw data.
func (e *EdgegapClient) MetricsByDeploymentID(id string, filter MetricsFilter) (*Response[Metrics], error) {
	query := "?"

	timeFmtString := "2006-01-02 15:04:05.000000"

	if filter.StartTime != nil {
		timeString := filter.StartTime.Format(timeFmtString)
		query += fmt.Sprintf("start_time=%s&", timeString)
	}

	if filter.EndTime != nil {
		timeString := filter.EndTime.Format(timeFmtString)
		query += fmt.Sprintf("end_time=%s&", timeString)
	}

	if filter.StepValue != nil {
		stepString := fmt.Sprintf("step=%d", *filter.StepValue)

		if filter.Step != nil {
			stepString += string(*filter.Step)
		} else {
			stepString += "s"
		}

		query += fmt.Sprintf("%s&", stepString)
	}

	if filter.Raw != nil {
		query += fmt.Sprintf("raw=%s&", *filter.Raw)
	}

	var response Metrics

	return makeRequest(e, func(c *resty.Request) (*resty.Response, error) {
		return c.Get(fmt.Sprintf("/metrics/deployment/%s%s", id, query))
	}, &response)
}
