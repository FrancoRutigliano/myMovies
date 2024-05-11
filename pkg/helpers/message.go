package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
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

func ParseJson(r *http.Request, Payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(Payload)
}

func WriteJson(w http.ResponseWriter, status int, data interface{}, entity string) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	jsonData := map[string]interface{}{
		entity: data,
	}

	return json.NewEncoder(w).Encode(jsonData)
}

func ReadIdParam(r *http.Request) (int64, error) {
	param := r.PathValue("id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id param")
	}

	return id, nil
}
