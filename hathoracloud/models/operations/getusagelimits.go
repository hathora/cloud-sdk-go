// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetUsageLimitsGlobals struct {
	OrgID *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *GetUsageLimitsGlobals) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

type GetUsageLimitsRequest struct {
	OrgID *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *GetUsageLimitsRequest) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}
