package main

import (
	"fmt"
    "net/http"
	"log"

	"github.com/joho/godotenv"
	"os"
    "github.com/gorilla/mux"
	"github.com/devkishor8007/word_master/src/database"
	"github.com/devkishor8007/word_master/src/routes"
)

func main() {
	database.InitDB()

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	router := mux.NewRouter()

	routes.SetupRoutes(router)

	server := &http.Server{
		Addr:    ":"+ port,
		Handler: router,
	}

	fmt.Println("Server is running on :"+port)
	log.Fatal(server.ListenAndServe())
}