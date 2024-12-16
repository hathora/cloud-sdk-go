// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"HathoraCloud/models/components"
)

type InitStripeCustomerPortalURLGlobals struct {
	OrgID *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *InitStripeCustomerPortalURLGlobals) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

type InitStripeCustomerPortalURLRequest struct {
	OrgID             *string                      `queryParam:"style=form,explode=true,name=orgId"`
	CustomerPortalURL components.CustomerPortalURL `request:"mediaType=application/json"`
}

func (o *InitStripeCustomerPortalURLRequest) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

func (o *InitStripeCustomerPortalURLRequest) GetCustomerPortalURL() components.CustomerPortalURL {
	if o == nil {
		return components.CustomerPortalURL{}
	}
	return o.CustomerPortalURL
}
