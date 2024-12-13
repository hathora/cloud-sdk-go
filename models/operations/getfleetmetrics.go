// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/internal/utils"
	"github.com/hathora/cloud-sdk-go/models/components"
)

type GetFleetMetricsGlobals struct {
	OrgID *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *GetFleetMetricsGlobals) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

type GetFleetMetricsRequest struct {
	FleetID string            `pathParam:"style=simple,explode=false,name=fleetId"`
	Region  components.Region `pathParam:"style=simple,explode=false,name=region"`
	// Available metrics to query over time.
	Metrics []components.FleetMetricName `queryParam:"style=form,explode=true,name=metrics"`
	End     *float64                     `queryParam:"style=form,explode=true,name=end"`
	// Unix timestamp. Default is -1 hour from `end`.
	Start *float64 `queryParam:"style=form,explode=true,name=start"`
	Step  *int     `default:"60" queryParam:"style=form,explode=true,name=step"`
	OrgID *string  `queryParam:"style=form,explode=true,name=orgId"`
}

func (g GetFleetMetricsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetFleetMetricsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetFleetMetricsRequest) GetFleetID() string {
	if o == nil {
		return ""
	}
	return o.FleetID
}

func (o *GetFleetMetricsRequest) GetRegion() components.Region {
	if o == nil {
		return components.Region("")
	}
	return o.Region
}

func (o *GetFleetMetricsRequest) GetMetrics() []components.FleetMetricName {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *GetFleetMetricsRequest) GetEnd() *float64 {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *GetFleetMetricsRequest) GetStart() *float64 {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *GetFleetMetricsRequest) GetStep() *int {
	if o == nil {
		return nil
	}
	return o.Step
}

func (o *GetFleetMetricsRequest) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}