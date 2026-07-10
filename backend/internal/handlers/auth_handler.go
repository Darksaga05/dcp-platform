package handlers

import (
	"encoding/json"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {

	var request RegisterRequest

	json.NewDecoder(r.Body).Decode(&request)

	response := map[string]string{
		"message": "User registration endpoint ready",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
func Login(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"message": "Login endpoint ready",
		"token":   "jwt_token_here",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
