// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetDeploymentInfoV1DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetDeploymentInfoV1DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetDeploymentInfoV1DeprecatedRequest struct {
	AppID        *string `pathParam:"style=simple,explode=false,name=appId"`
	DeploymentID int     `pathParam:"style=simple,explode=false,name=deploymentId"`
}

func (o *GetDeploymentInfoV1DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetDeploymentInfoV1DeprecatedRequest) GetDeploymentID() int {
	if o == nil {
		return 0
	}
	return o.DeploymentID
}
