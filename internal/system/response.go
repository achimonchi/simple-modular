package system

import (
	"errors"
	"fmt"
	"net/http"
)

type APIResponse struct {
	StatusCode int         `json:"-"`
	Code       string      `json:"code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload,omitempty"`
	Err        *string     `json:"error,omitempty"`
}

func ErrorInternal() APIResponse {
	return APIResponse{
		StatusCode: http.StatusInternalServerError,
		Code:       "5000000",
		Success:    false,
		Message:    "INTERNAL SERVER ERROR",
	}
}
func ErrorBadRequest() APIResponse {
	return APIResponse{
		StatusCode: http.StatusBadRequest,
		Code:       "4000000",
		Success:    false,
		Message:    "ERROR BAD REQUEST",
	}
}

func Success() APIResponse {
	return APIResponse{
		StatusCode: http.StatusOK,
		Code:       "2000000",
		Success:    true,
		Message:    "SUCCESS",
	}
}
func SuccessCreated() APIResponse {
	return APIResponse{
		StatusCode: http.StatusCreated,
		Code:       "2010000",
		Success:    true,
		Message:    "SUCCESS",
	}
}

func (a APIResponse) Error() error {
	if a.Err == nil {
		return errors.New(a.Message)
	}

	return fmt.Errorf("%v", a.Err)
}

// krna butuh overwrite, maka perlu ditambahkan pointer
func (a *APIResponse) SetPayload(payload interface{}) {
	a.Payload = payload
}

// krna butuh overwrite, maka perlu ditambahkan pointer
func (a *APIResponse) SetError(err error) {
	msg := err.Error()
	a.Err = &msg
}

// untuk lebih advance, bisa dibikin custom function seperti
// SethMessage(), SethStatusCode, SethCode, SethPayload, and others...
