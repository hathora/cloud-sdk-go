// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/hathoracloud/models/components"
)

type GetProcessesCountExperimentalV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetProcessesCountExperimentalV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetProcessesCountExperimentalV2DeprecatedRequest struct {
	AppID  *string                    `pathParam:"style=simple,explode=false,name=appId"`
	Status []components.ProcessStatus `queryParam:"style=form,explode=true,name=status"`
	Region []components.Region        `queryParam:"style=form,explode=true,name=region"`
}

func (o *GetProcessesCountExperimentalV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetProcessesCountExperimentalV2DeprecatedRequest) GetStatus() []components.ProcessStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetProcessesCountExperimentalV2DeprecatedRequest) GetRegion() []components.Region {
	if o == nil {
		return nil
	}
	return o.Region
}

// GetProcessesCountExperimentalV2DeprecatedResponseBody - Ok
type GetProcessesCountExperimentalV2DeprecatedResponseBody struct {
	Count float64 `json:"count"`
}

func (o *GetProcessesCountExperimentalV2DeprecatedResponseBody) GetCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Count
}
