package main

import (
	"fmt"
	"net/http"

	"github.com/dcp-project/backend/internal/config"
	"github.com/dcp-project/backend/internal/handlers"
)

func main() {

	appConfig := config.LoadConfig()

	http.HandleFunc("/health", handlers.HealthCheck)
	http.HandleFunc("/user", handlers.GetUser)
	http.HandleFunc("/messages", handlers.GetMessages)
	http.HandleFunc("/conversations", handlers.GetConversations)
	http.HandleFunc("/media", handlers.GetMedia)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	fmt.Println(appConfig.AppName + " server started on port " + appConfig.ServerPort)

	err := http.ListenAndServe(":"+appConfig.ServerPort, nil)

	if err != nil {
		panic(err)
	}
}
