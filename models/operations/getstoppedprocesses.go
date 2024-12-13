// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/models/components"
)

type GetStoppedProcessesGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetStoppedProcessesGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetStoppedProcessesRequest struct {
	AppID  *string            `pathParam:"style=simple,explode=false,name=appId"`
	Region *components.Region `queryParam:"style=form,explode=true,name=region"`
}

func (o *GetStoppedProcessesRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetStoppedProcessesRequest) GetRegion() *components.Region {
	if o == nil {
		return nil
	}
	return o.Region
}