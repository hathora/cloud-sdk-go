// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"hathoracloud/internal/utils"
	"time"
)

type LobbyV3CreatedByType string

const (
	LobbyV3CreatedByTypeStr    LobbyV3CreatedByType = "str"
	LobbyV3CreatedByTypeNumber LobbyV3CreatedByType = "number"
)

// LobbyV3CreatedBy - UserId or email address for the user that created the lobby.
type LobbyV3CreatedBy struct {
	Str    *string  `queryParam:"inline"`
	Number *float64 `queryParam:"inline"`

	Type LobbyV3CreatedByType
}

func CreateLobbyV3CreatedByStr(str string) LobbyV3CreatedBy {
	typ := LobbyV3CreatedByTypeStr

	return LobbyV3CreatedBy{
		Str:  &str,
		Type: typ,
	}
}

func CreateLobbyV3CreatedByNumber(number float64) LobbyV3CreatedBy {
	typ := LobbyV3CreatedByTypeNumber

	return LobbyV3CreatedBy{
		Number: &number,
		Type:   typ,
	}
}

func (u *LobbyV3CreatedBy) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = LobbyV3CreatedByTypeStr
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = LobbyV3CreatedByTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for LobbyV3CreatedBy", string(data))
}

func (u LobbyV3CreatedBy) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type LobbyV3CreatedBy: all fields are null")
}

// LobbyV3 - A lobby object allows you to store and manage metadata for your rooms.
type LobbyV3 struct {
	// User-defined identifier for a lobby.
	ShortCode string `json:"shortCode"`
	// When the lobby was created.
	CreatedAt time.Time `json:"createdAt"`
	// UserId or email address for the user that created the lobby.
	CreatedBy  LobbyV3CreatedBy `json:"createdBy"`
	RoomConfig *string          `json:"roomConfig,omitempty"`
	// Types of lobbies a player can create.
	//
	// `private`: the player who created the room must share the roomId with their friends
	//
	// `public`: visible in the public lobby list, anyone can join
	//
	// `local`: for testing with a server running locally
	Visibility LobbyVisibility `json:"visibility"`
	Region     Region          `json:"region"`
	// Unique identifier to a game session or match. Use the default system generated ID or overwrite it with your own.
	// Note: error will be returned if `roomId` is not globally unique.
	RoomID string `json:"roomId"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
}

func (l LobbyV3) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LobbyV3) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LobbyV3) GetShortCode() string {
	if o == nil {
		return ""
	}
	return o.ShortCode
}

func (o *LobbyV3) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *LobbyV3) GetCreatedBy() LobbyV3CreatedBy {
	if o == nil {
		return LobbyV3CreatedBy{}
	}
	return o.CreatedBy
}

func (o *LobbyV3) GetRoomConfig() *string {
	if o == nil {
		return nil
	}
	return o.RoomConfig
}

func (o *LobbyV3) GetVisibility() LobbyVisibility {
	if o == nil {
		return LobbyVisibility("")
	}
	return o.Visibility
}

func (o *LobbyV3) GetRegion() Region {
	if o == nil {
		return Region("")
	}
	return o.Region
}

func (o *LobbyV3) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}

func (o *LobbyV3) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}
