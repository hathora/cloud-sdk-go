// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SetLobbyStateParams struct {
	// JSON blob to store metadata for a room. Must be smaller than 1MB.
	State any `json:"state"`
}

func (o *SetLobbyStateParams) GetState() any {
	if o == nil {
		return nil
	}
	return o.State
}
