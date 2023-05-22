package db

import (
	"context"

	e "github.com/mal-mel/devices_api/internal/entity"
)

func (db *PGCon) SaveDevice(ctx context.Context, device e.Device, vendorId int) error {
	insertQuery := `INSERT INTO device (id, is_charging, battery_level, vendor_id, tags) VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Connection.Exec(ctx, insertQuery,
		device.Id,
		device.IsCharging,
		device.BatteryLevel,
		vendorId,
		device.Tags)
	if err != nil {
		return err
	}

	return nil
}

func (db *PGCon) SaveVendor(ctx context.Context, vendor e.Vendor) error {
	insertQuery := `INSERT INTO vendor (name) VALUES ($1)`

	_, err := db.Connection.Exec(ctx, insertQuery, vendor.Name)
	if err != nil {
		return err
	}

	return nil
}
