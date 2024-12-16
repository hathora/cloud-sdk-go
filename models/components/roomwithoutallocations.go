// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"HathoraCloud/internal/utils"
	"time"
)

// RoomWithoutAllocationsCurrentAllocation - Metadata on an allocated instance of a room.
type RoomWithoutAllocationsCurrentAllocation struct {
	UnscheduledAt *time.Time `json:"unscheduledAt"`
	ScheduledAt   time.Time  `json:"scheduledAt"`
	// System generated unique identifier to a runtime instance of your game server.
	ProcessID string `json:"processId"`
	// System generated unique identifier to an allocated instance of a room.
	RoomAllocationID string `json:"roomAllocationId"`
}

func (r RoomWithoutAllocationsCurrentAllocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RoomWithoutAllocationsCurrentAllocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RoomWithoutAllocationsCurrentAllocation) GetUnscheduledAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UnscheduledAt
}

func (o *RoomWithoutAllocationsCurrentAllocation) GetScheduledAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ScheduledAt
}

func (o *RoomWithoutAllocationsCurrentAllocation) GetProcessID() string {
	if o == nil {
		return ""
	}
	return o.ProcessID
}

func (o *RoomWithoutAllocationsCurrentAllocation) GetRoomAllocationID() string {
	if o == nil {
		return ""
	}
	return o.RoomAllocationID
}

type RoomWithoutAllocations struct {
	CurrentAllocation *RoomWithoutAllocationsCurrentAllocation `json:"currentAllocation"`
	// The allocation status of a room.
	//
	// `scheduling`: a process is not allocated yet and the room is waiting to be scheduled
	//
	// `active`: ready to accept connections
	//
	// `destroyed`: all associated metadata is deleted
	Status     RoomStatus `json:"status"`
	RoomConfig *string    `json:"roomConfig,omitempty"`
	// Unique identifier to a game session or match. Use the default system generated ID or overwrite it with your own.
	// Note: error will be returned if `roomId` is not globally unique.
	RoomID string `json:"roomId"`
	// System generated unique identifier for an application.
	AppID string `json:"appId"`
}

func (o *RoomWithoutAllocations) GetCurrentAllocation() *RoomWithoutAllocationsCurrentAllocation {
	if o == nil {
		return nil
	}
	return o.CurrentAllocation
}

func (o *RoomWithoutAllocations) GetStatus() RoomStatus {
	if o == nil {
		return RoomStatus("")
	}
	return o.Status
}

func (o *RoomWithoutAllocations) GetRoomConfig() *string {
	if o == nil {
		return nil
	}
	return o.RoomConfig
}

func (o *RoomWithoutAllocations) GetRoomID() string {
	if o == nil {
		return ""
	}
	return o.RoomID
}

func (o *RoomWithoutAllocations) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}
