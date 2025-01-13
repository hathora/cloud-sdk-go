// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/hathora/cloud-sdk-go/internal/utils"
	"time"
)

type PendingOrgInviteScopesType string

const (
	PendingOrgInviteScopesTypeUserRole     PendingOrgInviteScopesType = "UserRole"
	PendingOrgInviteScopesTypeArrayOfScope PendingOrgInviteScopesType = "arrayOfScope"
)

type PendingOrgInviteScopes struct {
	UserRole     *UserRole `queryParam:"inline"`
	ArrayOfScope []Scope   `queryParam:"inline"`

	Type PendingOrgInviteScopesType
}

func CreatePendingOrgInviteScopesUserRole(userRole UserRole) PendingOrgInviteScopes {
	typ := PendingOrgInviteScopesTypeUserRole

	return PendingOrgInviteScopes{
		UserRole: &userRole,
		Type:     typ,
	}
}

func CreatePendingOrgInviteScopesArrayOfScope(arrayOfScope []Scope) PendingOrgInviteScopes {
	typ := PendingOrgInviteScopesTypeArrayOfScope

	return PendingOrgInviteScopes{
		ArrayOfScope: arrayOfScope,
		Type:         typ,
	}
}

func (u *PendingOrgInviteScopes) UnmarshalJSON(data []byte) error {

	var userRole UserRole = UserRole("")
	if err := utils.UnmarshalJSON(data, &userRole, "", true, true); err == nil {
		u.UserRole = &userRole
		u.Type = PendingOrgInviteScopesTypeUserRole
		return nil
	}

	var arrayOfScope []Scope = []Scope{}
	if err := utils.UnmarshalJSON(data, &arrayOfScope, "", true, true); err == nil {
		u.ArrayOfScope = arrayOfScope
		u.Type = PendingOrgInviteScopesTypeArrayOfScope
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for PendingOrgInviteScopes", string(data))
}

func (u PendingOrgInviteScopes) MarshalJSON() ([]byte, error) {
	if u.UserRole != nil {
		return utils.MarshalJSON(u.UserRole, "", true)
	}

	if u.ArrayOfScope != nil {
		return utils.MarshalJSON(u.ArrayOfScope, "", true)
	}

	return nil, errors.New("could not marshal union type PendingOrgInviteScopes: all fields are null")
}

type PendingOrgInvite struct {
	Scopes    PendingOrgInviteScopes `json:"scopes"`
	CreatedAt time.Time              `json:"createdAt"`
	InvitedBy string                 `json:"invitedBy"`
	// A user's email.
	UserEmail string `json:"userEmail"`
	// System generated unique identifier for an organization. Not guaranteed to have a specific format.
	OrgID string `json:"orgId"`
}

func (p PendingOrgInvite) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PendingOrgInvite) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PendingOrgInvite) GetScopes() PendingOrgInviteScopes {
	if o == nil {
		return PendingOrgInviteScopes{}
	}
	return o.Scopes
}

func (o *PendingOrgInvite) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *PendingOrgInvite) GetInvitedBy() string {
	if o == nil {
		return ""
	}
	return o.InvitedBy
}

func (o *PendingOrgInvite) GetUserEmail() string {
	if o == nil {
		return ""
	}
	return o.UserEmail
}

func (o *PendingOrgInvite) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}
