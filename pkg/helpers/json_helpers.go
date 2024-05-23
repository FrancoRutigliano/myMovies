package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

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

// StoreJson serializa la estructura proporcionada y la guarda en un archivo JSON.
func StoreJson(filename string, v interface{}) error {
	// Serializar la estructura a JSON
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Errorf("error serializing data: %v", err)
	}

	// Escribir el archivo
	return os.WriteFile(filename, data, 0644)
}
