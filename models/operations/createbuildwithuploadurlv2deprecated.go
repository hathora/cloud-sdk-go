// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"hathoracloud/models/components"
)

type CreateBuildWithUploadURLV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateBuildWithUploadURLV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateBuildWithUploadURLV2DeprecatedRequest struct {
	AppID             *string                      `pathParam:"style=simple,explode=false,name=appId"`
	CreateBuildParams components.CreateBuildParams `request:"mediaType=application/json"`
}

func (o *CreateBuildWithUploadURLV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateBuildWithUploadURLV2DeprecatedRequest) GetCreateBuildParams() components.CreateBuildParams {
	if o == nil {
		return components.CreateBuildParams{}
	}
	return o.CreateBuildParams
}
