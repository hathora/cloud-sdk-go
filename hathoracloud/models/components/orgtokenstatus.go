// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type OrgTokenStatus string

const (
	OrgTokenStatusActive  OrgTokenStatus = "active"
	OrgTokenStatusRevoked OrgTokenStatus = "revoked"
)

func (e OrgTokenStatus) ToPointer() *OrgTokenStatus {
	return &e
}
