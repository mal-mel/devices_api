package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/mal-mel/devices_api/internal/entity"
)

func (env *Env) GetDevice(w http.ResponseWriter, r *http.Request, deviceId string) {
	ctx := r.Context()

	_, err := uuid.Parse(deviceId)
	if err != nil {
		e := UUIDParsingError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	device, err := env.DB.GetDevice(ctx, deviceId)
	if err != nil {
		e := DeviceNotFoundError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	err = sendResponse(w, http.StatusOK, device)
	if err != nil {
		e := UnknownError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}
}

func (env *Env) GetDeviceByVendor(w http.ResponseWriter, r *http.Request, vendor string) {
	ctx := r.Context()

	devices, err := env.DB.GetDevicesByVendor(ctx, vendor)
	if err != nil {
		e := UnknownError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	if devices == nil {
		e := DevicesByTagNotFoundError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	err = sendResponse(w, http.StatusOK, devices)
	if err != nil {
		e := UnknownError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}
}

func (env *Env) GetDeviceByTag(w http.ResponseWriter, r *http.Request, tag string) {
	ctx := r.Context()

	devices, err := env.DB.GetDevicesByTag(ctx, tag)
	if err != nil {
		e := UnknownError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	if devices == nil {
		e := DevicesByTagNotFoundError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	err = sendResponse(w, http.StatusOK, devices)
	if err != nil {
		e := UnknownError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}
}

func (env *Env) SaveDevice(w http.ResponseWriter, r *http.Request, deviceId string) {
	var err error

	ctx := r.Context()

	txStorage, transactionErr := env.DB.StartTransaction(ctx)
	if transactionErr != nil {
		return
	}

	defer func() {
		if err != nil {
			if txErr := txStorage.RollbackTransaction(ctx); txErr != nil {
				env.Log.WithField("funcName", "SaveDevice").
					WithError(errors.New("couldn't rollback tx")).
					Error(txErr)
			}
			env.Log.WithError(err).Error("couldn't save device")
		} else {
			_ = txStorage.CommitTransaction(ctx)
		}
	}()

	deviceUUID, err := uuid.Parse(deviceId)
	if err != nil {
		e := UUIDParsingError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	deviceExists, err := txStorage.IsDeviceExists(ctx, deviceId)
	if err != nil {
		e := UnknownError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	if deviceExists {
		e := DeviceAlreadyExistsError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	defer func() {
		_ = r.Body.Close()
	}()

	deviceData := entity.Device{
		Id: deviceUUID,
	}

	err = json.Unmarshal(bytes, &deviceData)
	if err != nil {
		e := DeviceBodyParsingError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	vendorId, err := txStorage.GetVendorIdByName(ctx, deviceData.Vendor)
	if err != nil {
		e := UnknownVendorNameError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	err = txStorage.SaveDevice(ctx, deviceData, vendorId)
	if err != nil {
		e := UnknownError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	err = sendResponse(w, http.StatusOK, deviceData)
	if err != nil {
		e := UnknownError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}
}
