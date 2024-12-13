// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/models/components"
)

type CreateDeploymentV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreateDeploymentV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreateDeploymentV2DeprecatedRequest struct {
	AppID              *string                       `pathParam:"style=simple,explode=false,name=appId"`
	BuildID            int                           `pathParam:"style=simple,explode=false,name=buildId"`
	DeploymentConfigV2 components.DeploymentConfigV2 `request:"mediaType=application/json"`
}

func (o *CreateDeploymentV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreateDeploymentV2DeprecatedRequest) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}

func (o *CreateDeploymentV2DeprecatedRequest) GetDeploymentConfigV2() components.DeploymentConfigV2 {
	if o == nil {
		return components.DeploymentConfigV2{}
	}
	return o.DeploymentConfigV2
}