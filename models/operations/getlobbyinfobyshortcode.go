// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetLobbyInfoByShortCodeGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLobbyInfoByShortCodeGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetLobbyInfoByShortCodeRequest struct {
	AppID     *string `pathParam:"style=simple,explode=false,name=appId"`
	ShortCode string  `pathParam:"style=simple,explode=false,name=shortCode"`
}

func (o *GetLobbyInfoByShortCodeRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetLobbyInfoByShortCodeRequest) GetShortCode() string {
	if o == nil {
		return ""
	}
	return o.ShortCode
}
