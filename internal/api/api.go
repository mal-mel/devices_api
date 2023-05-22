// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// Device defines model for device.
type Device struct {
	BatteryLevel *float32                `json:"battery_level,omitempty"`
	Id           *string                 `json:"id,omitempty"`
	IsCharging   *bool                   `json:"is_charging,omitempty"`
	Tags         *map[string]interface{} `json:"tags,omitempty"`
	Vendor       *string                 `json:"vendor,omitempty"`
}

// HttpError defines model for httpError.
type HttpError struct {
	Description *string `json:"description,omitempty"`
	Error       *string `json:"error,omitempty"`
	StatusCode  *int    `json:"status_code,omitempty"`
}

// Vendor defines model for vendor.
type Vendor struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// SaveDeviceJSONBody defines parameters for SaveDevice.
type SaveDeviceJSONBody Device

// SaveDeviceJSONRequestBody defines body for SaveDevice for application/json ContentType.
type SaveDeviceJSONRequestBody SaveDeviceJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получение девайса по его ID
	// (GET /device/{deviceId})
	GetDevice(w http.ResponseWriter, r *http.Request, deviceId string)
	// Сохранение нового девайса
	// (POST /device/{deviceId})
	SaveDevice(w http.ResponseWriter, r *http.Request, deviceId string)
	// Получение девайса по тегу
	// (GET /devices/tag/{tag})
	GetDeviceByTag(w http.ResponseWriter, r *http.Request, tag string)
	// Получение девайса по вендору
	// (GET /devices/vendor/{vendor})
	GetDeviceByVendor(w http.ResponseWriter, r *http.Request, vendor string)
	// Сохранение нового вендора
	// (POST /vendor/{vendorName})
	SaveVendor(w http.ResponseWriter, r *http.Request, vendorName string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// GetDevice operation middleware
func (siw *ServerInterfaceWrapper) GetDevice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "deviceId" -------------
	var deviceId string

	err = runtime.BindStyledParameter("simple", false, "deviceId", chi.URLParam(r, "deviceId"), &deviceId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "deviceId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDevice(w, r, deviceId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// SaveDevice operation middleware
func (siw *ServerInterfaceWrapper) SaveDevice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "deviceId" -------------
	var deviceId string

	err = runtime.BindStyledParameter("simple", false, "deviceId", chi.URLParam(r, "deviceId"), &deviceId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "deviceId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SaveDevice(w, r, deviceId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetDeviceByTag operation middleware
func (siw *ServerInterfaceWrapper) GetDeviceByTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "tag" -------------
	var tag string

	err = runtime.BindStyledParameter("simple", false, "tag", chi.URLParam(r, "tag"), &tag)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tag", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDeviceByTag(w, r, tag)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetDeviceByVendor operation middleware
func (siw *ServerInterfaceWrapper) GetDeviceByVendor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "vendor" -------------
	var vendor string

	err = runtime.BindStyledParameter("simple", false, "vendor", chi.URLParam(r, "vendor"), &vendor)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "vendor", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDeviceByVendor(w, r, vendor)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// SaveVendor operation middleware
func (siw *ServerInterfaceWrapper) SaveVendor(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "vendorName" -------------
	var vendorName string

	err = runtime.BindStyledParameter("simple", false, "vendorName", chi.URLParam(r, "vendorName"), &vendorName)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "vendorName", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SaveVendor(w, r, vendorName)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/device/{deviceId}", wrapper.GetDevice)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/device/{deviceId}", wrapper.SaveDevice)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/devices/tag/{tag}", wrapper.GetDeviceByTag)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/devices/vendor/{vendor}", wrapper.GetDeviceByVendor)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/vendor/{vendorName}", wrapper.SaveVendor)
	})

	return r
}
