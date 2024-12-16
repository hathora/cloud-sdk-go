// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"HathoraCloud/internal/utils"
	"time"
)

// ProcessV2ExposedPort - Connection details for an active process.
type ProcessV2ExposedPort struct {
	// Transport type specifies the underlying communication protocol to the exposed port.
	TransportType TransportType `json:"transportType"`
	Port          int           `json:"port"`
	Host          string        `json:"host"`
	Name          string        `json:"name"`
}

func (o *ProcessV2ExposedPort) GetTransportType() TransportType {
	if o == nil {
		return TransportType("")
	}
	return o.TransportType
}

func (o *ProcessV2ExposedPort) GetPort() int {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ProcessV2ExposedPort) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *ProcessV2ExposedPort) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type ProcessV2 struct {
	Status ProcessStatus `json:"status"`
	// Tracks the number of rooms that have been allocated to the process.
	RoomsAllocated int `json:"roomsAllocated"`
	// When the process has been terminated.
	TerminatedAt *time.Time `json:"terminatedAt"`
	// When the process is issued to stop. We use this to determine when we should stop billing.
	StoppingAt *time.Time `json:"stoppingAt"`
	// When the process bound to the specified port. We use this to determine when we should start billing.
	StartedAt *time.Time `json:"startedAt"`
	// When the process started being provisioned.
	CreatedAt time.Time `json:"createdAt"`
	// Governs how many [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) can be scheduled in a process.
	RoomsPerProcess        int                   `json:"roomsPerProcess"`
	AdditionalExposedPorts []ExposedPort         `json:"additionalExposedPorts"`
	ExposedPort            *ProcessV2ExposedPort `json:"exposedPort"`
	Region                 Region                `json:"region"`
	// System generated unique identifier to a runtime instance of your game server.
	ProcessID string `json:"processId"`
	// System generated id for a deployment. Increments by 1.
	DeploymentID int `json:"deploymentId"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
}

func (p ProcessV2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProcessV2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProcessV2) GetStatus() ProcessStatus {
	if o == nil {
		return ProcessStatus("")
	}
	return o.Status
}

func (o *ProcessV2) GetRoomsAllocated() int {
	if o == nil {
		return 0
	}
	return o.RoomsAllocated
}

func (o *ProcessV2) GetTerminatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.TerminatedAt
}

func (o *ProcessV2) GetStoppingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StoppingAt
}

func (o *ProcessV2) GetStartedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *ProcessV2) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ProcessV2) GetRoomsPerProcess() int {
	if o == nil {
		return 0
	}
	return o.RoomsPerProcess
}

func (o *ProcessV2) GetAdditionalExposedPorts() []ExposedPort {
	if o == nil {
		return []ExposedPort{}
	}
	return o.AdditionalExposedPorts
}

func (o *ProcessV2) GetExposedPort() *ProcessV2ExposedPort {
	if o == nil {
		return nil
	}
	return o.ExposedPort
}

func (o *ProcessV2) GetRegion() Region {
	if o == nil {
		return Region("")
	}
	return o.Region
}

func (o *ProcessV2) GetProcessID() string {
	if o == nil {
		return ""
	}
	return o.ProcessID
}

func (o *ProcessV2) GetDeploymentID() int {
	if o == nil {
		return 0
	}
	return o.DeploymentID
}

func (o *ProcessV2) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}
