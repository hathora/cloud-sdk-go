// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetBuildInfoDeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetBuildInfoDeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetBuildInfoDeprecatedRequest struct {
	AppID   *string `pathParam:"style=simple,explode=false,name=appId"`
	BuildID int     `pathParam:"style=simple,explode=false,name=buildId"`
}

func (o *GetBuildInfoDeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetBuildInfoDeprecatedRequest) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}
