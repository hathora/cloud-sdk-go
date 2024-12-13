// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetLobbyInfoGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLobbyInfoGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetLobbyInfoRequest struct {
	AppID  *string `pathParam:"style=simple,explode=false,name=appId"`
	RoomID string  `pathParam:"style=simple,explode=false,name=roomId"`
}

func (o *GetLobbyInfoRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetLobbyInfoRequest) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}