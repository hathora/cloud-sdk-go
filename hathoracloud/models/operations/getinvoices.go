// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetInvoicesGlobals struct {
	OrgID *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *GetInvoicesGlobals) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

type GetInvoicesRequest struct {
	OrgID *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *GetInvoicesRequest) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}
