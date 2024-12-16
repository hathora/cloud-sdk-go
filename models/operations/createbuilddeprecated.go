// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"HathoraCloud/models/components"
)

type CreateBuildDeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateBuildDeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateBuildDeprecatedRequest struct {
	AppID             *string                      `pathParam:"style=simple,explode=false,name=appId"`
	CreateBuildParams components.CreateBuildParams `request:"mediaType=application/json"`
}

func (o *CreateBuildDeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateBuildDeprecatedRequest) GetCreateBuildParams() components.CreateBuildParams {
	if o == nil {
		return components.CreateBuildParams{}
	}
	return o.CreateBuildParams
}
