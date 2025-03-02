// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetRoomInfoGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetRoomInfoGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetRoomInfoRequest struct {
	AppID  *string `pathParam:"style=simple,explode=false,name=appId"`
	RoomID string  `pathParam:"style=simple,explode=false,name=roomId"`
}

func (o *GetRoomInfoRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetRoomInfoRequest) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}
