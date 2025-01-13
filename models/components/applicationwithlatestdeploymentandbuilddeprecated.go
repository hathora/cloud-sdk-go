// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/hathora/cloud-sdk-go/internal/utils"
	"time"
)

type ApplicationWithLatestDeploymentAndBuildDeprecatedEnv struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecatedEnv) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecatedEnv) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type Deployment struct {
	// Option to shut down processes that have had no new connections or rooms
	// for five minutes.
	IdleTimeoutEnabled bool `json:"idleTimeoutEnabled"`
	// The environment variable that our process will have access to at runtime.
	Env []ApplicationWithLatestDeploymentAndBuildDeprecatedEnv `json:"env"`
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
	// A build represents a game server artifact and its associated metadata.
	Build Build `json:"build"`
}

func (d Deployment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Deployment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Deployment) GetIdleTimeoutEnabled() bool {
	if o == nil {
		return false
	}
	return o.IdleTimeoutEnabled
}

func (o *Deployment) GetEnv() []ApplicationWithLatestDeploymentAndBuildDeprecatedEnv {
	if o == nil {
		return []ApplicationWithLatestDeploymentAndBuildDeprecatedEnv{}
	}
	return o.Env
}

func (o *Deployment) GetRoomsPerProcess() int {
	if o == nil {
		return 0
	}
	return o.RoomsPerProcess
}

func (o *Deployment) GetAdditionalContainerPorts() []ContainerPort {
	if o == nil {
		return []ContainerPort{}
	}
	return o.AdditionalContainerPorts
}

func (o *Deployment) GetDefaultContainerPort() ContainerPort {
	if o == nil {
		return ContainerPort{}
	}
	return o.DefaultContainerPort
}

func (o *Deployment) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Deployment) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *Deployment) GetRequestedMemoryMB() float64 {
	if o == nil {
		return 0.0
	}
	return o.RequestedMemoryMB
}

func (o *Deployment) GetRequestedCPU() float64 {
	if o == nil {
		return 0.0
	}
	return o.RequestedCPU
}

func (o *Deployment) GetDeploymentID() int {
	if o == nil {
		return 0
	}
	return o.DeploymentID
}

func (o *Deployment) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}

func (o *Deployment) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *Deployment) GetBuild() Build {
	if o == nil {
		return Build{}
	}
	return o.Build
}

// ApplicationWithLatestDeploymentAndBuildDeprecated - An application object is the top level namespace for the game server.
type ApplicationWithLatestDeploymentAndBuildDeprecated struct {
	// The email address or token id for the user that deleted the application.
	DeletedBy *string `json:"deletedBy"`
	// When the application was deleted.
	DeletedAt *time.Time `json:"deletedAt"`
	// When the application was created.
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	// System generated unique identifier for an organization. Not guaranteed to have a specific format.
	OrgID string `json:"orgId"`
	// Configure [player authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service) for your application. Use Hathora's built-in auth providers or use your own [custom authentication](https://hathora.dev/docs/lobbies-and-matchmaking/auth-service#custom-auth-provider).
	AuthConfiguration AuthConfiguration `json:"authConfiguration"`
	// Secret that is used for identity and access management.
	AppSecret string `json:"appSecret"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
	// Readable name for an application. Must be unique within an organization.
	AppName    string      `json:"appName"`
	Deployment *Deployment `json:"deployment,omitempty"`
}

func (a ApplicationWithLatestDeploymentAndBuildDeprecated) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicationWithLatestDeploymentAndBuildDeprecated) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecated) GetDeletedBy() *string {
	if o == nil {
		return nil
	}
	return o.DeletedBy
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecated) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecated) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecated) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecated) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecated) GetAuthConfiguration() AuthConfiguration {
	if o == nil {
		return AuthConfiguration{}
	}
	return o.AuthConfiguration
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecated) GetAppSecret() string {
	if o == nil {
		return ""
	}
	return o.AppSecret
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecated) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecated) GetAppName() string {
	if o == nil {
		return ""
	}
	return o.AppName
}

func (o *ApplicationWithLatestDeploymentAndBuildDeprecated) GetDeployment() *Deployment {
	if o == nil {
		return nil
	}
	return o.Deployment
}
