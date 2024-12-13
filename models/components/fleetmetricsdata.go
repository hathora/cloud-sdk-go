// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type FleetMetricsData struct {
	Utilized             []MetricValue `json:"utilized,omitempty"`
	ProvisionedCloud     []MetricValue `json:"provisionedCloud,omitempty"`
	ProvisionedBareMetal []MetricValue `json:"provisionedBareMetal,omitempty"`
	ProvisionedTotal     []MetricValue `json:"provisionedTotal,omitempty"`
}

func (o *FleetMetricsData) GetUtilized() []MetricValue {
	if o == nil {
		return nil
	}
	return o.Utilized
}

func (o *FleetMetricsData) GetProvisionedCloud() []MetricValue {
	if o == nil {
		return nil
	}
	return o.ProvisionedCloud
}

func (o *FleetMetricsData) GetProvisionedBareMetal() []MetricValue {
	if o == nil {
		return nil
	}
	return o.ProvisionedBareMetal
}

func (o *FleetMetricsData) GetProvisionedTotal() []MetricValue {
	if o == nil {
		return nil
	}
	return o.ProvisionedTotal
}