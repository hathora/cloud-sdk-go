// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/hathoracloud/models/components"
)

type LoginGoogleGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *LoginGoogleGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type LoginGoogleRequest struct {
	AppID               *string                        `pathParam:"style=simple,explode=false,name=appId"`
	GoogleIDTokenObject components.GoogleIDTokenObject `request:"mediaType=application/json"`
}

func (o *LoginGoogleRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *LoginGoogleRequest) GetGoogleIDTokenObject() components.GoogleIDTokenObject {
	if o == nil {
		return components.GoogleIDTokenObject{}
	}
	return o.GoogleIDTokenObject
}
