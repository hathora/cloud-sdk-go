// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RescindUserInvite struct {
	// A user's email.
	UserEmail string `json:"userEmail"`
}

func (o *RescindUserInvite) GetUserEmail() string {
	if o == nil {
		return ""
	}
	return o.UserEmail
}
