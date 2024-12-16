// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"hathoracloud/models/components"
)

type LoginNicknameGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *LoginNicknameGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type LoginNicknameRequest struct {
	AppID          *string                   `pathParam:"style=simple,explode=false,name=appId"`
	NicknameObject components.NicknameObject `request:"mediaType=application/json"`
}

func (o *LoginNicknameRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *LoginNicknameRequest) GetNicknameObject() components.NicknameObject {
	if o == nil {
		return components.NicknameObject{}
	}
	return o.NicknameObject
}
