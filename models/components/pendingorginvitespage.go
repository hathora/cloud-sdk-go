// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type PendingOrgInvitesPage struct {
	Invites []PendingOrgInvite `json:"invites"`
}

func (o *PendingOrgInvitesPage) GetInvites() []PendingOrgInvite {
	if o == nil {
		return []PendingOrgInvite{}
	}
	return o.Invites
}
