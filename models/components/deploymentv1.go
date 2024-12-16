// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"HathoraCloud/internal/utils"
	"time"
)

type Env struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

func (o *Env) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *Env) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// DeploymentV1TransportType
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type DeploymentV1TransportType string

const (
	DeploymentV1TransportTypeTCP DeploymentV1TransportType = "tcp"
	DeploymentV1TransportTypeUDP DeploymentV1TransportType = "udp"
	DeploymentV1TransportTypeTLS DeploymentV1TransportType = "tls"
)

func (e DeploymentV1TransportType) ToPointer() *DeploymentV1TransportType {
	return &e
}

// DeploymentV1 - Deployment is a versioned configuration for a build that describes runtime behavior.
type DeploymentV1 struct {
	// Option to shut down processes that have had no new connections or rooms
	// for five minutes.
	IdleTimeoutEnabled *bool `default:"true" json:"idleTimeoutEnabled"`
	// The environment variable that our process will have access to at runtime.
	Env []Env `json:"env"`
	// Governs how many [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) can be scheduled in a process.
	RoomsPerProcess int `json:"roomsPerProcess"`
	// A plan defines how much CPU and memory is required to run an instance of your game server.
	//
	// `tiny`: shared core, 1gb memory
	//
	// `small`: 1 core, 2gb memory
	//
	// `medium`: 2 core, 4gb memory
	//
	// `large`: 4 core, 8gb memory
	PlanName PlanName `json:"planName"`
	// Additional ports your server listens on.
	AdditionalContainerPorts []ContainerPort `json:"additionalContainerPorts"`
	// A container port object represents the transport configruations for how your server will listen.
	DefaultContainerPort ContainerPort `json:"defaultContainerPort"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TransportType DeploymentV1TransportType `json:"transportType"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ContainerPort float64 `json:"containerPort"`
	// When the deployment was created.
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	// The amount of memory allocated to your process.
	RequestedMemoryMB int `json:"requestedMemoryMB"`
	// The number of cores allocated to your process.
	RequestedCPU float64 `json:"requestedCPU"`
	// System generated id for a deployment. Increments by 1.
	DeploymentID int `json:"deploymentId"`
	// System generated id for a build. Increments by 1.
	BuildID int `json:"buildId"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
}

func (d DeploymentV1) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeploymentV1) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DeploymentV1) GetIdleTimeoutEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.IdleTimeoutEnabled
}

func (o *DeploymentV1) GetEnv() []Env {
	if o == nil {
		return []Env{}
	}
	return o.Env
}

func (o *DeploymentV1) GetRoomsPerProcess() int {
	if o == nil {
		return 0
	}
	return o.RoomsPerProcess
}

func (o *DeploymentV1) GetPlanName() PlanName {
	if o == nil {
		return PlanName("")
	}
	return o.PlanName
}

func (o *DeploymentV1) GetAdditionalContainerPorts() []ContainerPort {
	if o == nil {
		return []ContainerPort{}
	}
	return o.AdditionalContainerPorts
}

func (o *DeploymentV1) GetDefaultContainerPort() ContainerPort {
	if o == nil {
		return ContainerPort{}
	}
	return o.DefaultContainerPort
}

func (o *DeploymentV1) GetTransportType() DeploymentV1TransportType {
	if o == nil {
		return DeploymentV1TransportType("")
	}
	return o.TransportType
}

func (o *DeploymentV1) GetContainerPort() float64 {
	if o == nil {
		return 0.0
	}
	return o.ContainerPort
}

func (o *DeploymentV1) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *DeploymentV1) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *DeploymentV1) GetRequestedMemoryMB() int {
	if o == nil {
		return 0
	}
	return o.RequestedMemoryMB
}

func (o *DeploymentV1) GetRequestedCPU() float64 {
	if o == nil {
		return 0.0
	}
	return o.RequestedCPU
}

func (o *DeploymentV1) GetDeploymentID() int {
	if o == nil {
		return 0
	}
	return o.DeploymentID
}

func (o *DeploymentV1) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}

func (o *DeploymentV1) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}
