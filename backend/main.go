package main

import (
	"fmt"
	"net/http"

	"github.com/dcp-project/backend/internal/handlers"
)

func main() {

	http.HandleFunc("/health", handlers.HealthCheck)

	fmt.Println("DCP Backend server started on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
