// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"hathoracloud/models/components"
)

type CreateDeploymentGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateDeploymentGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateDeploymentRequest struct {
	AppID              *string                       `pathParam:"style=simple,explode=false,name=appId"`
	DeploymentConfigV3 components.DeploymentConfigV3 `request:"mediaType=application/json"`
}

func (o *CreateDeploymentRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateDeploymentRequest) GetDeploymentConfigV3() components.DeploymentConfigV3 {
	if o == nil {
		return components.DeploymentConfigV3{}
	}
	return o.DeploymentConfigV3
}
