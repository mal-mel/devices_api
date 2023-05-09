package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type httpError struct {
	StatusCode  int    `json:"status_code"`
	Error       string `json:"error"`
	Description string `json:"description"`
}

type Err int

const (
	UnknownError Err = iota
	UUIDParsingError
	DeviceNotFoundError
	DeviceAlreadyExistsError
	DeviceBodyParsingError
	UnknownVendorNameError
)

var errorText = map[Err]string{
	UnknownError:             "Unknown error",
	UUIDParsingError:         "UUID parsing error",
	DeviceNotFoundError:      "Device not found error",
	DeviceAlreadyExistsError: "Device already exists error",
	DeviceBodyParsingError:   "Device body parsing error",
	UnknownVendorNameError:   "Unknown vendor name error",
}

var errorDescription = map[Err]string{
	UUIDParsingError:         "Не удалось прочитать идентификатор",
	DeviceNotFoundError:      "Не удалось получить объект",
	DeviceAlreadyExistsError: "Данный девайс уже существует",
	DeviceBodyParsingError:   "Ошибка при обработке данных девайса",
	UnknownVendorNameError:   "Неизвестное имя вендора",
}

var errorStatus = map[Err]int{
	UUIDParsingError:         http.StatusBadRequest,
	DeviceNotFoundError:      http.StatusNotFound,
	DeviceAlreadyExistsError: http.StatusConflict,
	DeviceBodyParsingError:   http.StatusBadRequest,
	UnknownVendorNameError:   http.StatusBadRequest,
}

func (c Err) errorMessage(e error) string {
	if e != nil {
		return fmt.Sprintf("%s: %s", c.error(), e.Error())
	}

	return c.error()
}

func (c Err) error() string {
	if s, ok := errorText[c]; ok {
		return s
	}

	return errorText[UnknownError]
}

func (c Err) status() int {
	if s, ok := errorStatus[c]; ok {
		return s
	}

	return http.StatusInternalServerError
}

func (c Err) description() string {
	if s, ok := errorDescription[c]; ok {
		return s
	}

	return errorDescription[UnknownError]
}

func (c Err) sendError(w http.ResponseWriter) error {
	resp := httpError{
		StatusCode:  c.status(),
		Error:       c.error(),
		Description: c.description(),
	}

	w.WriteHeader(resp.StatusCode)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}

	return nil
}
