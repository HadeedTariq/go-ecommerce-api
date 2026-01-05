package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ~ so here the parser exist for doing the marshalling related stuff
func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

// ~ similarly over there the unmarshaller also exist
func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// ~ ok so over there for the consistency we can also handle the error
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{
		"error": err.Error(),
	})
}
