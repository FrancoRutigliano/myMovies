package helpers

import (
	"fmt"
	"log/slog"
	"net/http"
)

type APIError struct {
	StatusCode int `json:"statuscode"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        err.Error(),
	}
}

func InvalidRequestError(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
	}
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid json request data"))
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func make(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(APIError); ok {
				//write json
			} else {
				errResp := map[string]any{
					"statuscode": http.StatusInternalServerError,
					"msg":        "internal server error",
				}

				//write json
			}
			slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}
