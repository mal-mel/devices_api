package entity

import "github.com/google/uuid"

type Device struct {
	Id           uuid.UUID         `json:"id"`
	IsCharging   CustomBool        `json:"is_charging"`
	BatteryLevel CustomFloat32     `json:"battery_level"`
	Vendor       string            `json:"vendor"`
	Tags         map[string]string `json:"tags"`
}
