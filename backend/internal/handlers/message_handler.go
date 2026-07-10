package handlers

import (
	"encoding/json"
	"net/http"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {

	messages := []map[string]string{
		{
			"sender":  "user_1",
			"message": "Hello DCP",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(messages)
}
