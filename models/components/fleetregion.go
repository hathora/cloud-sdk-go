// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/hathora/cloud-sdk-go/internal/utils"
	"time"
)

// FleetRegion - A fleet region is a region in which a fleet can be deployed.
// You can update cloudMinVcpus once every five minutes. It must be a multiple of
// scaleIncrementVcpus
type FleetRegion struct {
	CloudMinVcpusUpdatedAt time.Time `json:"cloudMinVcpusUpdatedAt"`
	CloudMinVcpus          int       `json:"cloudMinVcpus"`
	ScaleIncrementVcpus    int       `json:"scaleIncrementVcpus"`
}

func (f FleetRegion) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FleetRegion) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FleetRegion) GetCloudMinVcpusUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CloudMinVcpusUpdatedAt
}

func (o *FleetRegion) GetCloudMinVcpus() int {
	if o == nil {
		return 0
	}
	return o.CloudMinVcpus
}

func (o *FleetRegion) GetScaleIncrementVcpus() int {
	if o == nil {
		return 0
	}
	return o.ScaleIncrementVcpus
}
