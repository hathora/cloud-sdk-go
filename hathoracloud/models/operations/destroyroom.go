// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DestroyRoomGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *DestroyRoomGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type DestroyRoomRequest struct {
	AppID  *string `pathParam:"style=simple,explode=false,name=appId"`
	RoomID string  `pathParam:"style=simple,explode=false,name=roomId"`
}

func (o *DestroyRoomRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *DestroyRoomRequest) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}
