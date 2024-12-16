// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"hathoracloud/internal/utils"
	"hathoracloud/models/components"
)

type GetMetricsDeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetMetricsDeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetMetricsDeprecatedRequest struct {
	AppID     *string `pathParam:"style=simple,explode=false,name=appId"`
	ProcessID string  `pathParam:"style=simple,explode=false,name=processId"`
	// Available metrics to query over time.
	Metrics []components.DeprecatedProcessMetricName `queryParam:"style=form,explode=true,name=metrics"`
	// Unix timestamp. Default is current time.
	End *float64 `queryParam:"style=form,explode=true,name=end"`
	// Unix timestamp. Default is -1 hour from `end`.
	Start *float64 `queryParam:"style=form,explode=true,name=start"`
	Step  *int     `default:"60" queryParam:"style=form,explode=true,name=step"`
}

func (g GetMetricsDeprecatedRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetMetricsDeprecatedRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetMetricsDeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetMetricsDeprecatedRequest) GetProcessID() string {
	if o == nil {
		return ""
	}
	return o.ProcessID
}

func (o *GetMetricsDeprecatedRequest) GetMetrics() []components.DeprecatedProcessMetricName {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *GetMetricsDeprecatedRequest) GetEnd() *float64 {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *GetMetricsDeprecatedRequest) GetStart() *float64 {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *GetMetricsDeprecatedRequest) GetStep() *int {
	if o == nil {
		return nil
	}
	return o.Step
}
