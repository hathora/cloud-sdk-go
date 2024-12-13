// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/hathora/cloud-sdk-go/internal/utils"
	"time"
)

type RegionalContainerTags struct {
	ContainerTag string `json:"containerTag"`
	Region       Region `json:"region"`
}

func (o *RegionalContainerTags) GetContainerTag() string {
	if o == nil {
		return ""
	}
	return o.ContainerTag
}

func (o *RegionalContainerTags) GetRegion() Region {
	if o == nil {
		return Region("")
	}
	return o.Region
}

// Build - A build represents a game server artifact and its associated metadata.
type Build struct {
	BuildTag *string `json:"buildTag,omitempty"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RegionalContainerTags []RegionalContainerTags `json:"regionalContainerTags"`
	// The size (in bytes) of the Docker image built by Hathora.
	ImageSize int64       `json:"imageSize"`
	Status    BuildStatus `json:"status"`
	// When the build was deleted.
	DeletedAt *time.Time `json:"deletedAt"`
	// When [`RunBuild()`](https://hathora.dev/api#tag/BuildV2/operation/RunBuild) finished executing.
	FinishedAt *time.Time `json:"finishedAt"`
	// When [`RunBuild()`](https://hathora.dev/api#tag/BuildV2/operation/RunBuild) is called.
	StartedAt *time.Time `json:"startedAt"`
	// When [`CreateBuild()`](https://hathora.dev/api#tag/BuildV2/operation/CreateBuild) is called.
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
	// System generated id for a build. Increments by 1.
	BuildID int `json:"buildId"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
}

func (b Build) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *Build) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Build) GetBuildTag() *string {
	if o == nil {
		return nil
	}
	return o.BuildTag
}

func (o *Build) GetRegionalContainerTags() []RegionalContainerTags {
	if o == nil {
		return []RegionalContainerTags{}
	}
	return o.RegionalContainerTags
}

func (o *Build) GetImageSize() int64 {
	if o == nil {
		return 0
	}
	return o.ImageSize
}

func (o *Build) GetStatus() BuildStatus {
	if o == nil {
		return BuildStatus("")
	}
	return o.Status
}

func (o *Build) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Build) GetFinishedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.FinishedAt
}

func (o *Build) GetStartedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *Build) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Build) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *Build) GetBuildID() int {
	if o == nil {
		return 0
	}
	return o.BuildID
}

func (o *Build) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}