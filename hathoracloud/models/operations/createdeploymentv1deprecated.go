// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/hathoracloud/models/components"
)

type CreateDeploymentV1DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateDeploymentV1DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateDeploymentV1DeprecatedRequest struct {
	AppID            *string                     `pathParam:"style=simple,explode=false,name=appId"`
	BuildID          int                         `pathParam:"style=simple,explode=false,name=buildId"`
	DeploymentConfig components.DeploymentConfig `request:"mediaType=application/json"`
}

func (o *CreateDeploymentV1DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateDeploymentV1DeprecatedRequest) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}

func (o *CreateDeploymentV1DeprecatedRequest) GetDeploymentConfig() components.DeploymentConfig {
	if o == nil {
		return components.DeploymentConfig{}
	}
	return o.DeploymentConfig
}
