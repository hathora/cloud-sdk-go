// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DeleteBuildDeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *DeleteBuildDeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type DeleteBuildDeprecatedRequest struct {
	AppID   *string `pathParam:"style=simple,explode=false,name=appId"`
	BuildID int     `pathParam:"style=simple,explode=false,name=buildId"`
}

func (o *DeleteBuildDeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *DeleteBuildDeprecatedRequest) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}