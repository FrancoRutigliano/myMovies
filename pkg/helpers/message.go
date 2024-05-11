package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
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

func InvalidRequestData(errors map[string]string) APIError {
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
				WriteJson(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"statuscode": http.StatusInternalServerError,
					"msg":        "internal server error",
				}
				WriteJson(w, http.StatusInternalServerError, errResp)
			}
			slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func ParseJson(r *http.Request, Payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(Payload)
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "applcation/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func ReadIdParam(r *http.Request) (int64, error) {
	param := r.PathValue("id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id param")
	}

	return id, nil
}
