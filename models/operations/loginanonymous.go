// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type LoginAnonymousGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *LoginAnonymousGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type LoginAnonymousRequest struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *LoginAnonymousRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}