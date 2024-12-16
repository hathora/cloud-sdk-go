// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"HathoraCloud/models/components"
)

type ListActivePublicLobbiesGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *ListActivePublicLobbiesGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type ListActivePublicLobbiesRequest struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
	// If omitted, active public lobbies in all regions will be returned.
	Region *components.Region `queryParam:"style=form,explode=true,name=region"`
}

func (o *ListActivePublicLobbiesRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *ListActivePublicLobbiesRequest) GetRegion() *components.Region {
	if o == nil {
		return nil
	}
	return o.Region
}
