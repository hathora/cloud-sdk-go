// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type StaticProcessAllocationConfig struct {
	// The maximum number of running processes that can be spun up during upgrades
	// Invariant: minProcesses <= maxProcesses
	MaxProcesses int `json:"maxProcesses"`
	// The target number of running processes
	TargetProcesses int `json:"targetProcesses"`
	// The minimum running processes required during upgrades.
	// Invariant: 0 <= minProcesses < targetProcesses
	MinProcesses int    `json:"minProcesses"`
	Region       Region `json:"region"`
}

func (o *StaticProcessAllocationConfig) GetMaxProcesses() int {
	if o == nil {
		return 0
	}
	return o.MaxProcesses
}

func (o *StaticProcessAllocationConfig) GetTargetProcesses() int {
	if o == nil {
		return 0
	}
	return o.TargetProcesses
}

func (o *StaticProcessAllocationConfig) GetMinProcesses() int {
	if o == nil {
		return 0
	}
	return o.MinProcesses
}

func (o *StaticProcessAllocationConfig) GetRegion() Region {
	if o == nil {
		return Region("")
	}
	return o.Region
}
