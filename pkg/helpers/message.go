package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
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

func NewAPIError(statusCode int, message string) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        message,
	}
}

func SendCustom(w http.ResponseWriter, statusCode int, message string) {
	apiErr := NewAPIError(statusCode, message)
	WriteJson(w, statusCode, apiErr)
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
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
