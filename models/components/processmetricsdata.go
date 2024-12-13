// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ProcessMetricsData struct {
	ActiveConnections []MetricValue `json:"activeConnections,omitempty"`
	TotalEgress       []MetricValue `json:"totalEgress,omitempty"`
	RateEgress        []MetricValue `json:"rateEgress,omitempty"`
	Memory            []MetricValue `json:"memory,omitempty"`
	CPU               []MetricValue `json:"cpu,omitempty"`
}

func (o *ProcessMetricsData) GetActiveConnections() []MetricValue {
	if o == nil {
		return nil
	}
	return o.ActiveConnections
}

func (o *ProcessMetricsData) GetTotalEgress() []MetricValue {
	if o == nil {
		return nil
	}
	return o.TotalEgress
}

func (o *ProcessMetricsData) GetRateEgress() []MetricValue {
	if o == nil {
		return nil
	}
	return o.RateEgress
}

func (o *ProcessMetricsData) GetMemory() []MetricValue {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *ProcessMetricsData) GetCPU() []MetricValue {
	if o == nil {
		return nil
	}
	return o.CPU
}
