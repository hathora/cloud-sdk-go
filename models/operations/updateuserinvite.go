// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"HathoraCloud/models/components"
)

type UpdateUserInviteRequest struct {
	OrgID            string                      `pathParam:"style=simple,explode=false,name=orgId"`
	UpdateUserInvite components.UpdateUserInvite `request:"mediaType=application/json"`
}

func (o *UpdateUserInviteRequest) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *UpdateUserInviteRequest) GetUpdateUserInvite() components.UpdateUserInvite {
	if o == nil {
		return components.UpdateUserInvite{}
	}
	return o.UpdateUserInvite
}
