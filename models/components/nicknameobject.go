// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type NicknameObject struct {
	// An alias to represent a player.
	Nickname string `json:"nickname"`
}

func (o *NicknameObject) GetNickname() string {
	if o == nil {
		return ""
	}
	return o.Nickname
}
