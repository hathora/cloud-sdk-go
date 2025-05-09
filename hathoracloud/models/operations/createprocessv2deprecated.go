// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/hathoracloud/models/components"
)

type CreateProcessV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateProcessV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateProcessV2DeprecatedRequest struct {
	AppID  *string           `pathParam:"style=simple,explode=false,name=appId"`
	Region components.Region `pathParam:"style=simple,explode=false,name=region"`
}

func (o *CreateProcessV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateProcessV2DeprecatedRequest) GetRegion() components.Region {
	if o == nil {
		return components.Region("")
	}
	return o.Region
}
