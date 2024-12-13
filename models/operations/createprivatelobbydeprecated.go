// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/internal/utils"
	"github.com/hathora/cloud-sdk-go/models/components"
)

type CreatePrivateLobbyDeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *CreatePrivateLobbyDeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type CreatePrivateLobbyDeprecatedSecurity struct {
	PlayerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=hathora_player_auth"`
}

func (o *CreatePrivateLobbyDeprecatedSecurity) GetPlayerAuth() string {
	if o == nil {
		return ""
	}
	return o.PlayerAuth
}

type CreatePrivateLobbyDeprecatedRequest struct {
	AppID  *string            `pathParam:"style=simple,explode=false,name=appId"`
	Region *components.Region `queryParam:"style=form,explode=true,name=region"`
	Local  *bool              `default:"false" queryParam:"style=form,explode=true,name=local"`
}

func (c CreatePrivateLobbyDeprecatedRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePrivateLobbyDeprecatedRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePrivateLobbyDeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *CreatePrivateLobbyDeprecatedRequest) GetRegion() *components.Region {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *CreatePrivateLobbyDeprecatedRequest) GetLocal() *bool {
	if o == nil {
		return nil
	}
	return o.Local
}
