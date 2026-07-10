package handlers

import (
	"encoding/json"
	"net/http"
)

func GetConversations(w http.ResponseWriter, r *http.Request) {

	conversations := []map[string]string{
		{
			"id":     "conversation_1",
			"status": "active",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(conversations)
}
