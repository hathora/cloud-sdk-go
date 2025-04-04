// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/hathora/cloud-sdk-go/hathoracloud/internal/utils"
	"time"
)

type DeploymentV2Env struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

func (o *DeploymentV2Env) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *DeploymentV2Env) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type DeploymentV2 struct {
	// Option to shut down processes that have had no new connections or rooms
	// for five minutes.
	IdleTimeoutEnabled bool `json:"idleTimeoutEnabled"`
	// The environment variable that our process will have access to at runtime.
	Env []DeploymentV2Env `json:"env"`
	// Governs how many [rooms](https://hathora.dev/docs/concepts/hathora-entities#room) can be scheduled in a process.
	RoomsPerProcess int `json:"roomsPerProcess"`
	// Additional ports your server listens on.
	AdditionalContainerPorts []ContainerPort `json:"additionalContainerPorts"`
	// A container port object represents the transport configruations for how your server will listen.
	DefaultContainerPort ContainerPort `json:"defaultContainerPort"`
	// When the deployment was created.
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	// The amount of memory allocated to your process.
	RequestedMemoryMB float64 `json:"requestedMemoryMB"`
	// The number of cores allocated to your process.
	RequestedCPU float64 `json:"requestedCPU"`
	// System generated id for a deployment. Increments by 1.
	DeploymentID int `json:"deploymentId"`
	// System generated id for a build. Increments by 1.
	BuildID int `json:"buildId"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
}

func (d DeploymentV2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DeploymentV2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DeploymentV2) GetIdleTimeoutEnabled() bool {
	if o == nil {
		return false
	}
	return o.IdleTimeoutEnabled
}

func (o *DeploymentV2) GetEnv() []DeploymentV2Env {
	if o == nil {
		return []DeploymentV2Env{}
	}
	return o.Env
}

func (o *DeploymentV2) GetRoomsPerProcess() int {
	if o == nil {
		return 0
	}
	return o.RoomsPerProcess
}

func (o *DeploymentV2) GetAdditionalContainerPorts() []ContainerPort {
	if o == nil {
		return []ContainerPort{}
	}
	return o.AdditionalContainerPorts
}

func (o *DeploymentV2) GetDefaultContainerPort() ContainerPort {
	if o == nil {
		return ContainerPort{}
	}
	return o.DefaultContainerPort
}

func (o *DeploymentV2) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *DeploymentV2) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *DeploymentV2) GetRequestedMemoryMB() float64 {
	if o == nil {
		return 0.0
	}
	return o.RequestedMemoryMB
}

func (o *DeploymentV2) GetRequestedCPU() float64 {
	if o == nil {
		return 0.0
	}
	return o.RequestedCPU
}

func (o *DeploymentV2) GetDeploymentID() int {
	if o == nil {
		return 0
	}
	return o.DeploymentID
}

func (o *DeploymentV2) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}

func (o *DeploymentV2) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}
