package helpers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

type APIError struct {
	StatusCode int `json:"statuscode"`
	Msg        any `json:"msg"`
}

func NewAPIError(statusCode int, message string) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        message,
	}
}

func SendCustom(w http.ResponseWriter, statusCode int, message string) {
	apiErr := NewAPIError(statusCode, message)
	WriteJson(w, statusCode, apiErr, "error")
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
	}
}

func ReadIdParam(r *http.Request) (int64, error) {
	param := r.PathValue("id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id param")
	}

	return id, nil
}
