// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// Fleet - A fleet is a collection of vCPUs accross your regions that can scale up and down based on demand.
type Fleet struct {
	// System generated unique identifier for an organization. Not guaranteed to have a specific format.
	OrgID string `json:"orgId"`
	// the id of the fleet
	FleetID string `json:"fleetId"`
}

func (o *Fleet) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *Fleet) GetFleetID() string {
	if o == nil {
		return ""
	}
	return o.FleetID
}