// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/models/components"
)

type CreateBuildV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateBuildV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateBuildV2DeprecatedRequest struct {
	AppID             *string                      `pathParam:"style=simple,explode=false,name=appId"`
	CreateBuildParams components.CreateBuildParams `request:"mediaType=application/json"`
}

func (o *CreateBuildV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateBuildV2DeprecatedRequest) GetCreateBuildParams() components.CreateBuildParams {
	if o == nil {
		return components.CreateBuildParams{}
	}
	return o.CreateBuildParams
}