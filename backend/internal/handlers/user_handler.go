package handlers

import (
	"encoding/json"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	user := map[string]string{
		"username": "dcp_user",
		"status":   "active",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
