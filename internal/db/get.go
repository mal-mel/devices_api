package db

import (
	"context"

	e "github.com/mal-mel/devices_api/internal/entity"
)

func (db *PGCon) GetDevice(ctx context.Context, deviceId string) (e.Device, error) {
	q := `SELECT * FROM device LEFT JOIN vendor ON device.vendor_id = vendor.id WHERE device.id = $1`

	var device e.Device

	if err := db.Connection.QueryRow(ctx, q, deviceId).Scan(&device); err != nil {
		return device, err
	}

	return device, nil
}

func (db *PGCon) IsDeviceExists(ctx context.Context, deviceId string) (bool, error) {
	q := `SELECT EXISTS(SELECT 1 FROM device WHERE device.id = $1);`

	var exists bool

	if err := db.Connection.QueryRow(ctx, q, deviceId).Scan(&exists); err != nil {
		return exists, err
	}

	return exists, nil
}

func (db *PGCon) GetVendorIdByName(ctx context.Context, vendorName string) (int, error) {
	q := `SELECT id FROM vendor WHERE name = $1`

	var id int

	if err := db.Connection.QueryRow(ctx, q, vendorName).Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (db *PGCon) GetDevicesByVendor(ctx context.Context, vendor string) ([]e.Device, error) {
	q := `SELECT
			device.id,
			device.is_charging,
			device.battery_level,
			vendor.name,
			device.tags
		  FROM device
		  LEFT JOIN vendor ON device.vendor_id = vendor.id
		  WHERE vendor.name = $1`

	var devices []e.Device

	rows, err := db.Connection.Query(ctx, q, vendor)
	if err != nil {
		return devices, err
	}

	for rows.Next() {
		var device e.Device

		err = rows.Scan(
			&device.Id,
			&device.IsCharging,
			&device.BatteryLevel,
			&device.Vendor,
			&device.Tags,
		)
		if err != nil {
			return devices, err
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func (db *PGCon) GetDevicesByTag(ctx context.Context, tag string) ([]e.Device, error) {
	q := `SELECT
			device.id,
			device.is_charging,
			device.battery_level,
			vendor.name,
			device.tags
		  FROM device
		  LEFT JOIN vendor ON device.vendor_id = vendor.id 
		  WHERE device.tags->$1 IS NOT NULL`

	var devices []e.Device

	rows, err := db.Connection.Query(ctx, q, tag)
	if err != nil {
		return devices, err
	}

	for rows.Next() {
		var device e.Device

		err = rows.Scan(
			&device.Id,
			&device.IsCharging,
			&device.BatteryLevel,
			&device.Vendor,
			&device.Tags,
		)
		if err != nil {
			return devices, err
		}

		devices = append(devices, device)
	}

	return devices, nil
}
