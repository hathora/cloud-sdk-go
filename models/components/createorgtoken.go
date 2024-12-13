// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/hathora/cloud-sdk-go/internal/utils"
)

type Scopes2 string

const (
	Scopes2Admin Scopes2 = "admin"
)

func (e Scopes2) ToPointer() *Scopes2 {
	return &e
}

type CreateOrgTokenScopesType string

const (
	CreateOrgTokenScopesTypeArrayOfScope CreateOrgTokenScopesType = "arrayOfScope"
	CreateOrgTokenScopesTypeScopes2      CreateOrgTokenScopesType = "scopes_2"
)

// CreateOrgTokenScopes - If not defined, the token has Admin access.
type CreateOrgTokenScopes struct {
	ArrayOfScope []Scope  `queryParam:"inline"`
	Scopes2      *Scopes2 `queryParam:"inline"`

	Type CreateOrgTokenScopesType
}

func CreateCreateOrgTokenScopesArrayOfScope(arrayOfScope []Scope) CreateOrgTokenScopes {
	typ := CreateOrgTokenScopesTypeArrayOfScope

	return CreateOrgTokenScopes{
		ArrayOfScope: arrayOfScope,
		Type:         typ,
	}
}

func CreateCreateOrgTokenScopesScopes2(scopes2 Scopes2) CreateOrgTokenScopes {
	typ := CreateOrgTokenScopesTypeScopes2

	return CreateOrgTokenScopes{
		Scopes2: &scopes2,
		Type:    typ,
	}
}

func (u *CreateOrgTokenScopes) UnmarshalJSON(data []byte) error {

	var arrayOfScope []Scope = []Scope{}
	if err := utils.UnmarshalJSON(data, &arrayOfScope, "", true, true); err == nil {
		u.ArrayOfScope = arrayOfScope
		u.Type = CreateOrgTokenScopesTypeArrayOfScope
		return nil
	}

	var scopes2 Scopes2 = Scopes2("")
	if err := utils.UnmarshalJSON(data, &scopes2, "", true, true); err == nil {
		u.Scopes2 = &scopes2
		u.Type = CreateOrgTokenScopesTypeScopes2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateOrgTokenScopes", string(data))
}

func (u CreateOrgTokenScopes) MarshalJSON() ([]byte, error) {
	if u.ArrayOfScope != nil {
		return utils.MarshalJSON(u.ArrayOfScope, "", true)
	}

	if u.Scopes2 != nil {
		return utils.MarshalJSON(u.Scopes2, "", true)
	}

	return nil, errors.New("could not marshal union type CreateOrgTokenScopes: all fields are null")
}

type CreateOrgToken struct {
	// If not defined, the token has Admin access.
	Scopes *CreateOrgTokenScopes `json:"scopes,omitempty"`
	// Readable name for a token. Must be unique within an organization.
	Name string `json:"name"`
}

func (o *CreateOrgToken) GetScopes() *CreateOrgTokenScopes {
	if o == nil {
		return nil
	}
	return o.Scopes
}

func (o *CreateOrgToken) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
