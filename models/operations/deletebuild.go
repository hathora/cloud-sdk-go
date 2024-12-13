// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DeleteBuildGlobals struct {
	OrgID *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *DeleteBuildGlobals) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

type DeleteBuildRequest struct {
	BuildID string  `pathParam:"style=simple,explode=false,name=buildId"`
	OrgID   *string `queryParam:"style=form,explode=true,name=orgId"`
}

func (o *DeleteBuildRequest) GetBuildID() string {
	if o == nil {
		return ""
	}
	return o.BuildID
}

func (o *DeleteBuildRequest) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}