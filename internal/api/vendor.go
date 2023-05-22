package api

import (
	"net/http"

	"github.com/mal-mel/devices_api/internal/entity"
)

func (env *Env) SaveVendor(w http.ResponseWriter, r *http.Request, vendorName string) {
	ctx := r.Context()

	vendor := entity.Vendor{
		Name: vendorName,
	}

	err := env.DB.SaveVendor(ctx, vendor)
	if err != nil {
		e := UnknownError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}

	err = sendResponse(w, http.StatusOK, vendor)
	if err != nil {
		e := UnknownError
		env.Log.Error(e.errorMessage(err))
		_ = e.sendError(w)
		return
	}
}
