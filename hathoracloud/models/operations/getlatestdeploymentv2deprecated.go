// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetLatestDeploymentV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLatestDeploymentV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetLatestDeploymentV2DeprecatedRequest struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLatestDeploymentV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}
