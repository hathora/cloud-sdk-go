// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/models/components"
)

type GetFleetRegionGlobals struct {
	OrgID *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *GetFleetRegionGlobals) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

type GetFleetRegionRequest struct {
	FleetID string            `pathParam:"style=simple,explode=false,name=fleetId"`
	Region  components.Region `pathParam:"style=simple,explode=false,name=region"`
	OrgID   *string           `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *GetFleetRegionRequest) GetFleetID() string {
	if o == nil {
		return ""
	}
	return o.FleetID
}

func (o *GetFleetRegionRequest) GetRegion() components.Region {
	if o == nil {
		return components.Region("")
	}
	return o.Region
}

func (o *GetFleetRegionRequest) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}