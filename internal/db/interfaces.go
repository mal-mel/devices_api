package db

import (
	c "context"

	e "github.com/mal-mel/devices_api/internal/entity"
)

type DeviceStorage interface {
	GetDevice(ctx c.Context, deviceId string) (e.Device, error)
	SaveDevice(ctx c.Context, device e.Device, vendorId int) error
	IsDeviceExists(ctx c.Context, deviceId string) (bool, error)
	GetVendorIdByName(ctx c.Context, vendorName string) (int, error)
}

type Database interface {
	DeviceStorage

	Ping(ctx c.Context) error

	StartTransaction(ctx c.Context) (Database, error)
	CommitTransaction(ctx c.Context) error
	RollbackTransaction(ctx c.Context) error
}
