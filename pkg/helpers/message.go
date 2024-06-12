package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

type APIError struct {
	StatusCode int    `json:"statuscode"`
	Msg        string `json:"msg"`
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

func ReadIdParam(r *http.Request) (int64, error) {
	param := r.PathValue("id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id param")
	}

	return id, nil
}

func InitializeStoreWithDefaults(f string, data interface{}) error {
	file, err := os.Create(f)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(data); err != nil {
		return err
	}
	return nil
}

func IsValid(r *http.Request, p interface{}) APIError {
	err := ParseJson(r, &p)
	if err != nil {
		return NewAPIError(
			http.StatusBadRequest,
			"invalid request data",
		)
	}

	if err = Validate.Struct(p); err != nil {
		return NewAPIError(
			http.StatusUnprocessableEntity,
			err.Error(),
		)
	}
	return APIError{}
}
