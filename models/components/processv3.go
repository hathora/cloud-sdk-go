// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/hathora/cloud-sdk-go/internal/utils"
	"time"
)

// ProcessV3ExposedPort - Connection details for an active process.
type ProcessV3ExposedPort struct {
	// Transport type specifies the underlying communication protocol to the exposed port.
	TransportType TransportType `json:"transportType"`
	Port          int           `json:"port"`
	Host          string        `json:"host"`
	Name          string        `json:"name"`
}

func (o *ProcessV3ExposedPort) GetTransportType() TransportType {
	if o == nil {
		return TransportType("")
	}
	return o.TransportType
}

func (o *ProcessV3ExposedPort) GetPort() int {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ProcessV3ExposedPort) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *ProcessV3ExposedPort) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type ProcessV3 struct {
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
	ExposedPort            *ProcessV3ExposedPort `json:"exposedPort"`
	Region                 Region                `json:"region"`
	// System generated unique identifier to a runtime instance of your game server.
	ProcessID string `json:"processId"`
	// System generated id for a deployment.
	DeploymentID string `json:"deploymentId"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
}

func (p ProcessV3) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProcessV3) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProcessV3) GetStatus() ProcessStatus {
	if o == nil {
		return ProcessStatus("")
	}
	return o.Status
}

func (o *ProcessV3) GetRoomsAllocated() int {
	if o == nil {
		return 0
	}
	return o.RoomsAllocated
}

func (o *ProcessV3) GetTerminatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.TerminatedAt
}

func (o *ProcessV3) GetStoppingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StoppingAt
}

func (o *ProcessV3) GetStartedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *ProcessV3) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ProcessV3) GetRoomsPerProcess() int {
	if o == nil {
		return 0
	}
	return o.RoomsPerProcess
}

func (o *ProcessV3) GetAdditionalExposedPorts() []ExposedPort {
	if o == nil {
		return []ExposedPort{}
	}
	return o.AdditionalExposedPorts
}

func (o *ProcessV3) GetExposedPort() *ProcessV3ExposedPort {
	if o == nil {
		return nil
	}
	return o.ExposedPort
}

func (o *ProcessV3) GetRegion() Region {
	if o == nil {
		return Region("")
	}
	return o.Region
}

func (o *ProcessV3) GetProcessID() string {
	if o == nil {
		return ""
	}
	return o.ProcessID
}

func (o *ProcessV3) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

func (o *ProcessV3) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}
