// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/models/components"
)

type CreatePrivateLobbyGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreatePrivateLobbyGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreatePrivateLobbySecurity struct {
	PlayerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=hathora_player_auth"`
}

func (o *CreatePrivateLobbySecurity) GetPlayerAuth() string {
	if o == nil {
		return ""
	}
	return o.PlayerAuth
}

type CreatePrivateLobbyRequestBody struct {
	// User input to initialize the game state. Object must be smaller than 64KB.
	InitialConfig any               `json:"initialConfig"`
	Region        components.Region `json:"region"`
}

func (o *CreatePrivateLobbyRequestBody) GetInitialConfig() any {
	if o == nil {
		return nil
	}
	return o.InitialConfig
}

func (o *CreatePrivateLobbyRequestBody) GetRegion() components.Region {
	if o == nil {
		return components.Region("")
	}
	return o.Region
}

type CreatePrivateLobbyRequest struct {
	AppID       *string                       `pathParam:"style=simple,explode=false,name=appId"`
	RoomID      *string                       `queryParam:"style=form,explode=true,name=roomId"`
	RequestBody CreatePrivateLobbyRequestBody `request:"mediaType=application/json"`
}

func (o *CreatePrivateLobbyRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreatePrivateLobbyRequest) GetRoomID() *string {
	if o == nil {
		return nil
	}
	return o.RoomID
}

func (o *CreatePrivateLobbyRequest) GetRequestBody() CreatePrivateLobbyRequestBody {
	if o == nil {
		return CreatePrivateLobbyRequestBody{}
	}
	return o.RequestBody
}
