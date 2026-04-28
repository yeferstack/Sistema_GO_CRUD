package controllers

import (
	"encoding/json"
	"net/http"
)

// respondJSON - helper compartido para todos los controllers
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
