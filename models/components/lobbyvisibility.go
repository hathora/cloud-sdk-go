// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// LobbyVisibility - Types of lobbies a player can create.
//
// `private`: the player who created the room must share the roomId with their friends
//
// `public`: visible in the public lobby list, anyone can join
//
// `local`: for testing with a server running locally
type LobbyVisibility string

const (
	LobbyVisibilityPrivate LobbyVisibility = "private"
	LobbyVisibilityPublic  LobbyVisibility = "public"
	LobbyVisibilityLocal   LobbyVisibility = "local"
)

func (e LobbyVisibility) ToPointer() *LobbyVisibility {
	return &e
}