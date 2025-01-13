// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/cloud-sdk-go/models/components"
)

type InviteUserRequest struct {
	OrgID            string                      `pathParam:"style=simple,explode=false,name=orgId"`
	CreateUserInvite components.CreateUserInvite `request:"mediaType=application/json"`
}

func (o *InviteUserRequest) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *InviteUserRequest) GetCreateUserInvite() components.CreateUserInvite {
	if o == nil {
		return components.CreateUserInvite{}
	}
	return o.CreateUserInvite
}
