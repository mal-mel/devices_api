package entity

import "github.com/google/uuid"

type Device struct {
	Id           uuid.UUID         `json:"id"`
	IsCharging   bool              `json:"is_charging"`
	BatteryLevel float32           `json:"battery_level"`
	Vendor       string            `json:"vendor"`
	Tags         map[string]string `json:"tags"`
}
