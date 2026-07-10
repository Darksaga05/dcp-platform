package handlers

import (
	"encoding/json"
	"net/http"
)

func GetMedia(w http.ResponseWriter, r *http.Request) {

	media := []map[string]string{
		{
			"type": "image",
			"url":  "example.jpg",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(media)
}
