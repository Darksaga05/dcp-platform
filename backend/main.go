package main

import (
	"fmt"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "DCP Backend is running")
}

func main() {

	http.HandleFunc("/health", healthCheck)

	fmt.Println("DCP Backend server started on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
