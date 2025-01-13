// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/models/components"
)

type GetLatestProcessesV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLatestProcessesV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetLatestProcessesV2DeprecatedRequest struct {
	AppID  *string                    `pathParam:"style=simple,explode=false,name=appId"`
	Status []components.ProcessStatus `queryParam:"style=form,explode=true,name=status"`
	Region []components.Region        `queryParam:"style=form,explode=true,name=region"`
}

func (o *GetLatestProcessesV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetLatestProcessesV2DeprecatedRequest) GetStatus() []components.ProcessStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetLatestProcessesV2DeprecatedRequest) GetRegion() []components.Region {
	if o == nil {
		return nil
	}
	return o.Region
}
